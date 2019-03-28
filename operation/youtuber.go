package operation

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang/glog"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"pure-media/config"
	"pure-media/database"
	"pure-media/models"
	"pure-media/storage"
	"strings"
)

//var (
//	filename    = flag.String("filename", "./operation/sss.mp4", "Name of video file to upload")
//	title       = flag.String("title", "Test Title", "Video title")
//	description = flag.String("description", "Test Description", "Video description")
//	category    = flag.String("category", "22", "Video category")
//	keywords    = flag.String("keywords", "", "Comma separated list of video keywords")
//	privacy     = flag.String("privacy", "unlisted", "Video privacy status")
//)

func DownloadVideo(videoUrl, format string) (fid string, err error) {

	//根据video_fid检测数据库中是否已经有该视频
	id, err := GetYouTubeId(videoUrl)
	if err != nil {
		err = errors.New(config.INVALID_URL)
		return "", err
	}

	db := database.NewDB()
	defer db.Close()
	videoMsg, err := db.Rechecking(id)
	if err == nil && videoMsg.TubeFid != "" {
		glog.V(5).Infoln(ISEXIST)
		return videoMsg.TubeFid, err
	}

	cmd := exec.Command("youtube-dl", "-f", format, videoUrl, "--proxy", "http://172.16.0.18:10081", "--verbose")

	//显示运行的命令
	//fmt.Println(cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	cmd.Start()
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	var videoname string
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		//fmt.Println(line)
		if strings.Contains(line, "Merging formats into") {
			//名称不对记得改坐标，最好改成正则匹配
			videoname = line[31 : len(line)-2]
		}
		if strings.Contains(line, "has already been downloaded and merged") {
			//名称不对记得改坐标，最好改成正则匹配
			videoname = line[11 : len(line)-40]
		}

	}

	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	cmd.Wait()
	//检测文件是否存在
	if isexist := Exists(videoname); !isexist {
		return "", err
	}

	//下载完成上传s3
	key := fmt.Sprintf("/%s", videoname)
	//base64
	key = base64.RawURLEncoding.EncodeToString([]byte(key))
	s3 := storage.NewS3Storage(&config.MediaConfig.S3ConfYoutube)
	bys, err := ioutil.ReadFile(videoname)
	if err != nil {
		glog.Errorf("ioutil.ReadFile %s \n", err)
		return "", err
	}
	fid, err = s3.CopyFileByReader(bytes.NewReader(bys), int64(len(bys)), videoname, key)
	if err != nil {
		glog.Errorf("s3.CopyFileByReader %s \n", err)
		return "", err
	}

	info := &models.VideoMsg{
		TubeFid: fid,
		VideoId: id,
	}
	//更新数据库里fid
	err = db.UpdateInfo(info)
	if err != nil {
		glog.Errorf("UpdateInfo crawler_tube %s \n", err)
		//return "", err
	}
	err = os.Remove(videoname)
	if err != nil {

	}

	return fid, nil
}

//func GetDownloadLink(videoUrl string) (err error) {
//	//watchnet := "https://zh.savefrom.net/#feature=youtu.be&utm_source=youtube.com&utm_medium=short_domains&utm_campaign=www.ssyoutube.com"
//	//url=http://youtube.com/watch?v=a-6VVwObr3g
//	//videoUrl = strings.Replace(videoUrl, "youtube", "ssyoutube", 1)
//	//videoUrl = watchnet + videoUrl
//
//	u, err := url.Parse("https://zh.savefrom.net/#url=http://youtube.com/watch?v=a-6VVwObr3g&feature=youtu.be&utm_source=youtube.com&utm_medium=short_domains&utm_campaign=www.ssyoutube.com")
//	if err != nil {
//		return
//	}
//	q := u.Query()
//	//q.Set("url", videoUrl)
//	u.RawQuery = q.Encode()
//	req, err := http.NewRequest(
//		"GET",
//		u.String(),
//		nil,
//	)
//	resp, err := config.NoProxyHttpClient.Do(req)
//	if err != nil {
//		glog.Errorf("http.Get err %s", err)
//		return err
//	}
//	defer resp.Body.Close()
//
//	//glog.V(5).Infoln(resp.Body)
//	//result, err := ioutil.ReadAll(resp.Body)
//	//if err != nil {
//	//	glog.Errorf("ioutil.ReadAll err %s", err)
//	//
//	//	return err
//	//}
//	//glog.V(5).Infoln(string(result))
//
//	doc, err := goquery.NewDocumentFromReader(resp.Body)
//	if err != nil {
//		glog.Errorf("goquery NewDocument err %s \n", err)
//		return err
//	}
//	//sf_url
//	txt := doc.Find(".info-box").Find(".def-btn-box").Text()
//	//
//	//text := doc.Find("#sf_result > div > div.result-box.video > div.info-box > div.link-box.showall > div.def-btn-box > a").Text()
//	glog.Infoln("sssss", txt)
//
//	//u, err := url.Parse("https://r3---sn-a5mlrnez.googlevideo.com/videoplayback?gir=yes&ipbits=0&lmt=1551300333048502&ip=117.255.223.59&sparams=aitags,clen,dur,ei,expire,gir,id,ip,ipbits,itag,keepalive,lmt,mime,mip,mm,mn,ms,mv,pl,requiressl,source&mime=video%2Fmp4&id=o-AGs5tpCpu4Rs9bIK7YyflTxI8YSHZURjFEA7Iqj078xm&dur=723.999&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278%2C298%2C299%2C302%2C303&pl=20&itag=137&requiressl=yes&expire=1551426712&txp=2316222&key=cms1&keepalive=yes&signature=31B6389D40A861674EE815E26BBA5499E7E3FBED.308D28B9F27C8AED1A8B039C52DB266970EA965D&ei=OJB4XLaOEJiaz7sPpNKksAU&source=youtube&c=WEB&clen=315020163&fvip=3&video_id=zhQ1DbLeBko&title=Warframe+-+Update+24.3.0+-+Nightwave+Arrives+On+All+Platforms%21%21&rm=sn-cnoa-qxal7e&req_id=e397ba67b6a5a3ee&redirect_counter=2&cm2rm=sn-qxak7s&cms_redirect=yes&mip=45.79.111.80&mm=34&mn=sn-a5mlrnez&ms=ltu&mt=1551405103&mv=m")
//	//if err != nil {
//	//	return
//	//}
//	//
//	//q := u.Query()
//	//q.Set("Retry-After", "20")
//	//u.RawQuery = q.Encode()
//	//req, err := http.NewRequest(
//	//	"GET",
//	//	u.String(),
//	//	nil,
//	//)
//	//num := rand.Intn(len(ippoll.IPPoll))
//	//hc := ippoll.IPPoll[num]
//	//proxy := hc.ProxyAddr
//	//fmt.Println("proxy ", proxy)
//	//resp, err := hc.Do(req)
//	//if err != nil {
//	//	glog.Errorf("client.Get err %s", err)
//	//	return
//	//}
//	//defer resp.Body.Close()
//	//
//	//file, _ := os.Create("sss" + ".mp4")
//	//defer file.Close()
//	//io.Copy(file, resp.Body)
//	return nil
//}
//
//func UploadAccount() {
//
//	flag.Parse()
//
//	if *filename == "" {
//		glog.Error("You must provide a filename of a video file to upload")
//	}
//
//	service, err := uptube.MakeClient()
//	if err != nil {
//		glog.Errorf("Error creating YouTube client: %v", err)
//	}
//
//	upload := &youtube.Video{
//		Snippet: &youtube.VideoSnippet{
//			Title:       *title,
//			Description: *description,
//			CategoryId:  *category,
//		},
//		Status: &youtube.VideoStatus{PrivacyStatus: *privacy},
//	}
//
//	// The API returns a 400 Bad Request response if tags is an empty string.
//	if strings.Trim(*keywords, "") != "" {
//		upload.Snippet.Tags = strings.Split(*keywords, ",")
//	}
//
//	call := service.Videos.Insert("snippet,status", upload)
//
//	file, err := os.Open(*filename)
//	defer file.Close()
//	if err != nil {
//		glog.Errorf("Error opening %v: %v", *filename, err)
//	}
//
//	response, err := call.Media(file).Do()
//	if err != nil {
//		glog.Errorf("call Media err %s", err)
//		return
//	}
//	//handleError(err, "")
//	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
//
//}
// Sample go code for videos.insert
//func printVideosInsertResults(response *youtube.Video) {
//	// Handle response here
//	fmt.Println(response.Id)
//}
//func videosInsert(service *youtube.Service, part string, res string, filename string) {
//	resource := &youtube.Video{}
//	if err := json.NewDecoder(strings.NewReader(res)).Decode(&resource); err != nil {
//		log.Fatal(err)
//	}
//	call := service.Videos.Insert(part, resource)
//
//	file, err := os.Open(filename)
//	defer file.Close()
//	if err != nil {
//		log.Fatalf("Error opening %v: %v", filename, err)
//	}
//
//	response, err := call.Media(file).Do()
//	if err != nil {
//		glog.Errorf("call.Media err %s", err)
//		return
//	}
//	//handleError(err, "")
//	printVideosInsertResults(response)
//}

//func Add() {
//	properties := (map[string]string{
//		"snippet.categoryId":         "22",
//		"snippet.defaultLanguage":    "",
//		"snippet.description":        "Description of uploaded video.",
//		"snippet.tags[]":             "",
//		"snippet.title":              "Test video upload",
//		"status.embeddable":          "",
//		"status.license":             "",
//		"status.privacyStatus":       "private",
//		"status.publicStatsViewable": "",
//	})
//	res := createResource(properties)
//
//	filename := "./operation/sss.mp4"
//	// Note: service variable must already be defined.
//	service, err := uptube.MakeClient()
//	if err != nil {
//		glog.Errorf("Error creating YouTube client: %v", err)
//	}
//
//	videosInsert(service, "snippet,status", res, filename)
//
//}
//func createResource(properties map[string]string) string {
//	resource := make(map[string]interface{})
//	for key, value := range properties {
//		keys := strings.Split(key, ".")
//		ref := addPropertyToResource(resource, keys, value, 0)
//		resource = ref
//	}
//	propJson, err := json.Marshal(resource)
//	if err != nil {
//		log.Fatal("cannot encode to JSON ", err)
//	}
//	return string(propJson)
//}
//func addPropertyToResource(ref map[string]interface{}, keys []string, value string, count int) map[string]interface{} {
//	for k := count; k < (len(keys) - 1); k++ {
//		switch val := ref[keys[k]].(type) {
//		case map[string]interface{}:
//			ref[keys[k]] = addPropertyToResource(val, keys, value, (k + 1))
//		case nil:
//			next := make(map[string]interface{})
//			ref[keys[k]] = addPropertyToResource(next, keys, value, (k + 1))
//		}
//	}
//	// Only include properties that have values.
//	if count == len(keys)-1 && value != "" {
//		valueKey := keys[len(keys)-1]
//		if valueKey[len(valueKey)-2:] == "[]" {
//			ref[valueKey[0:len(valueKey)-2]] = strings.Split(value, ",")
//		} else if len(valueKey) > 4 && valueKey[len(valueKey)-4:] == "|int" {
//			ref[valueKey[0:len(valueKey)-4]], _ = strconv.Atoi(value)
//		} else if value == "true" {
//			ref[valueKey] = true
//		} else if value == "false" {
//			ref[valueKey] = false
//		} else {
//			ref[valueKey] = value
//		}
//	}
//	return ref
//}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//func Ytubeupload() (err error) {
//
//	cmd := exec.Command("yt-upload", "-u")
//	var input, output, stderr bytes.Buffer
//	//var pipIn,pipOut bytes.Buffer
//	cmd.Stderr = &stderr
//	cmd.Stdout = &output
//	cmd.Stdin = &input
//
//	err = cmd.Run()
//	if err != nil {
//		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
//		return err
//	}
//	fmt.Println("Result: " + output.String())
//	if !strings.Contains(output.String(), "Enter video Path") {
//		err = errors.New("cmd is wrong")
//		return err
//	}
//	input.Write([]byte("ssssssssss"))
//
//	//stdin, _ := cmd.StdinPipe()
//	//stdin.Write([]byte("ssssssssss"))
//
//	//video := "test.mp4"
//	//fmt.Scan(&video)
//	//fmt.Println(video)
//	//c1 := exec.Command("echo", video)
//	//c2 := exec.Command("yt-upload", "-u")
//	//c3 := exec.Command("echo", "this is a test video")
//	//c2.Stdin, _ = c1.StdoutPipe()
//	//
//	//stdin, err := c2.StdinPipe()
//	//if err != nil {
//	//	glog.Errorf("c2.StdinPipe err %v \n", stdin)
//	//	return err
//	//}
//	//go func() {
//	//	defer stdin.Close()
//	//	io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
//	//}()
//	//
//	//c2.Stdout = os.Stdout
//	//_ = c2.Start()
//	//_ = c1.Run()
//	////c2.Stdin, _ = c3.StdoutPipe()
//	//_ = c3.Run()
//	//_ = c2.Wait()
//
//	select {}
//}
