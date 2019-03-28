package operation

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/glog"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"pure-media/config"
	"pure-media/database"
	"pure-media/ippoll"
	"pure-media/models"
	"pure-media/storage"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	DEFAULTYOUTUBEAPIKEY    = "AIzaSyCe1fzSMGP2z16W2_FVp3JA-cJ7_gbbG6k"
	DEFAULTYOUTUBEAPIPART   = "snippet,statistics,contentDetails"
	DEFAULTYOUTUBEAPIFIELDS = "*"
	DEFAULTYOUTUBEAPI       = "https://www.googleapis.com/youtube/v3/videos?"
	DEFAULTLOCATION         = "Asia/Shanghai"
	ISEXIST                 = "record is exist"
)

var (
	grep         = regexp.MustCompile("/([\\w-]{11}$)")
	durationregx = regexp.MustCompile("P(|\\d+D)T(|\\d+H)(|\\d+M)(|\\d+S)$")
)

func GetYoutubeWatchReply(videoUrl string) (videoMsg *models.VideoMsg, err error) {

	//检查url合法性
	if !IsValidUrl(videoUrl) {
		err = errors.New(config.INVALID_URL)
		return nil, err
	}
	id, err := GetYouTubeId(videoUrl)
	if err != nil {
		return nil, err
	}
	//数据库查找结果为空时，重新爬虫
	db := database.NewDB()
	defer db.Close()
	videoMsg, err = db.Rechecking(id)
	if err == nil && videoMsg.Title != "" {
		glog.V(5).Infoln(ISEXIST)
		return videoMsg, err
	}

	videoMsg, err = GetVideoMsg(id)
	if err != nil {
		return videoMsg, err
	}
	//存储videoMsg
	//TODO 添加tube_fid, new_tube_id, type（default，copy）, copy_state（下载中，上传中，搬运中，完成）

	err = db.SaveTubeInfo(videoMsg)
	if err != nil {
		return videoMsg, err
	}

	return videoMsg, err
}

//根据URL获取YoutubeId
/*
1.'youtube.com', 'www.youtube.com', 'm.youtube.com', 'gaming.youtube.com')
2.'youtu.be', 'www.youtu.be'
*/
func GetYouTubeId(videoUrl string) (id string, err error) {
	u, err := url.Parse(videoUrl)
	if err != nil {
		glog.Errorf("url.Parse err %s \n", err)
		return id, err

	}
	q := u.Query()
	id = q.Get("v")
	if id != "" {
		return id, err
	}
	id = grep.FindString(videoUrl)
	if id != "" && strings.HasPrefix(id, "/") {
		id = strings.Replace(id, "/", "", 1)
		return id, err

	}
	return id, err

}

//根据Id获取info
func GetRespYouTubeById(id string) (data []byte, err error) {
	if id == "" {
		return nil, err
	}
	u, err := url.Parse(DEFAULTYOUTUBEAPI)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("id", id)
	q.Set("key", config.MediaConfig.YtApiKey)
	q.Set("part", config.MediaConfig.YtApiPart)
	q.Set("fields", config.MediaConfig.YtApiFields)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(
		"GET",
		u.String(),
		nil,
	)
	glog.V(5).Infof("get_video_info_api:%v", u.String())
	//req.Header.Set("Accept", config.HTTP_ACCEPT)
	//req.Header.Set("User-Agent", config.HTTP_USER_AGENT)
	num := rand.Intn(len(ippoll.IPPoll))
	hc := ippoll.IPPoll[num]
	resp, err := hc.Do(req)
	if err != nil {
		glog.Errorf("id %s,err %v \n", id, err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("getVideoInfo hc.Do url:%v resp.code:%v", u.String(), resp.StatusCode)
		return nil, err
	}

	//Handle HTTP compression for gzip
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
		break
	default:
		reader = resp.Body
	}

	data, err = ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetVideoMsg(id string) (videoMsg *models.VideoMsg, err error) {
	data, err := GetRespYouTubeById(id)
	if err != nil {
		return nil, err
	}

	answer := models.YTApiV3Video{}
	err = json.Unmarshal(data, &answer)
	if err != nil {
		glog.Errorf("videoId: %s, json.Unmarshal err  %s \n", id, err)
		return nil, err
	}
	if len(answer.Items) <= 0 {
		err = errors.New(fmt.Sprintf("videoId:%s,%s", id, "response.answer.Items is nil"))
		return nil, err
	}
	snippet := answer.Items[0].Snippet
	if snippet == nil {
		err = fmt.Errorf("err:%v snippet:%v", err, snippet)
		return nil, err
	}
	info := map[string]interface{}{}
	info["statistics"] = answer.Items[0].Statistics
	info["contentDetails"] = answer.Items[0].ContentDetails
	info["snippet"] = answer.Items[0].Snippet
	info["title"] = snippet.ChannelTitle
	infoByte, err := json.Marshal(info)
	if err != nil {
		glog.Errorf("json.Marshal: %s \n", err)
		return nil, err
	}

	//获取视频信息
	thumbnails := answer.Items[0].Snippet.Thumbnails
	sceenUrl := getThumbnailUrl(thumbnails)
	lengthSeconds := strconv.Itoa(Parseduration(answer.Items[0].ContentDetails.Duration))

	//上传截图
	screenFid, err := UploadSceenToS3(sceenUrl, id)
	if err != nil {
		glog.Infof("UploadSceenToS3 url: %s, err:%s\n", sceenUrl, err)
	}

	videoMsg = &models.VideoMsg{
		VideoId:          id,
		Title:            snippet.Title,
		LengthSeconds:    lengthSeconds,
		ScreenFid:        screenFid,
		ScreenYoutubeUrl: sceenUrl,
		KeyWords:         snippet.Tags,
		Info:             string(infoByte),
		Who:              snippet.ChannelTitle,
		//HL:,
		//Author:     snippet.ChannelTitle,
		//Rating:     answer.Items[0].Statistics.LikeCount,
		//ViewCount:  answer.Items[0].Statistics.ViewCount,
		//Definition: answer.Items[0].ContentDetails.Definition,
		Created: time.Now(),
	}

	return videoMsg, nil
}

func getThumbnailUrl(thumbnail *models.ThumbnailInfo) (link string) {
	if thumbnail == nil {
		return ""
	}
	//优先级max
	if thumbnail.Maxres != nil {
		return thumbnail.Maxres.Url
	}

	if thumbnail.Standard != nil {
		return thumbnail.Standard.Url
	}

	if thumbnail.High != nil {
		return thumbnail.High.Url
	}

	if thumbnail.Medium != nil {
		return thumbnail.Medium.Url
	}

	if thumbnail.Default != nil {
		return thumbnail.Default.Url
	}
	return ""
}

func UploadSceenToS3(sceenUrl, videoId string) (fid string, err error) {
	key := fmt.Sprintf("/%s", videoId)
	s3 := storage.NewS3Storage(&config.MediaConfig.S3ConfYoutube)
	fid, err = s3.CopyFileByUrl(sceenUrl, key)
	if err != nil {
		return "", err
	}
	/*

		if sceenUrl == "" {
			err = errors.New(config.INVALID)
			return fid, err
		}
		purl, _ := url.Parse(config.MediaConfig.ProxyHttp)
		proxy := http.ProxyURL(purl)
		client := &http.Client{
			Transport: &http.Transport{
				Proxy: proxy,
			},
		}

		req, err := MakeCheckUrl(sceenUrl)
		if err != nil {
			return fid, err

		}
		resp, err := client.Do(req)
		if err != nil {
			glog.V(0).Infof("client.Get %s\n", err)
			return fid, err
		}
		//glog.V(5).Infof("open %s ret: %v", sceenUrl, resp)
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			err = errors.New(fmt.Sprintf("%d", resp.StatusCode))
			return fid, err
		}

		fid, err = s3.CopyFileByReader(resp.Body, resp.ContentLength, config.IMAGE, key)
	*/
	return fid, err
}

//PT29M23S
func Parseduration(str string) (lengthSeconds int) {
	if str == "" {
		return lengthSeconds
	}

	lenth := durationregx.FindStringSubmatch(str)
	if len(lenth) <= 1 {
		return lengthSeconds
	}
	lenth = lenth[1:]
	//fmt.Println("lenth", lenth)

	for _, v := range lenth {
		if v != "" {
			//glog.V(2).Infoln("v = ", v[len(v)-1:])
			switch v[len(v)-1:] {
			case "D":
				day, _ := strconv.Atoi(v[:len(v)-1])
				lengthSeconds += day * 60 * 60 * 24

			case "H":
				hour, _ := strconv.Atoi(v[:len(v)-1])
				lengthSeconds += hour * 60 * 60

			case "M":
				minute, _ := strconv.Atoi(v[:len(v)-1])
				lengthSeconds += minute * 60

			case "S":
				second, _ := strconv.Atoi(v[:len(v)-1])
				lengthSeconds += second * 1

			}
		}

	}

	return lengthSeconds
}

func HalfSearch(nums []int, low, high, key int) int {

	if key == nums[high] {
		return high
	}
	if key == nums[low] {
		return low
	}
	if low <= high && high > 0 {
		for {
			mid := low + (high-low)/2
			if nums[mid] == key {
				return mid
			} else if key < nums[mid] {
				high = mid
			} else if key > nums[mid] {
				low = mid
			}
		}
	}

	return -1
}

func FirstUniqChar(s string) int {
	length := len(s)
	k := 0
	if length <= 0 {
		return -1
	}

	repeated := false
	for i := 0; i < length; i++ {
		repeated = false
		for j := 0; j < length; j++ {
			if j != i && s[j] == s[i] {
				repeated = true
				break
			}
			k = i
		}
		if !repeated {
			return k
		}

	}

	return -1

}

func IsValidUrl(link string) bool {
	if link == "" {
		return false
	}
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		return false
	}

	//strRegex := "(http://|ftp://|https://|www){0,1}[^\u4e00-\u9fa5\\s]*?\\.(com|net|cn|me|tw|fr)[^\u4e00-\u9fa5\\s]*"
	//match, _ := regexp.MatchString(strRegex, link)
	return true
}
