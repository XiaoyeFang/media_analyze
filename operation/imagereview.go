package operation

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/glog"
	_ "github.com/qiniu/api.v7"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/webp"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"math/rand"
	"net/http"
	"pure-media/config"
	"pure-media/green-go-sdk/src/greensdksample"
	"pure-media/green-go-sdk/src/uuid"
	"pure-media/ippoll"
)

var (
	accessKey = "cbZkHxsv04xB-AsBajbFAUW4ZFk0GCAV0WtTPlP2"
	secretKey = "Ohwl8s8zsXEkF3do610pDuOjyTzm7Ty1DTVwqTkn"
	key       = "icon.png"
	localFile = "icon.png"
)

const (
	METHOD      = "POST"
	URL         = "http://ai.qiniuapi.com/v3/image/censor"
	HOST        = "ai.qiniuapi.com"
	CONTENTTYPE = "application/json"
)

//func ImgScanPorn(imgUrl string) (reply *protos.ScanPornReply, err error) {
//	//response := []byte(`{"code":200,
//	//		"data":[{
//	//			"code":200,
//	//			"dataId":"59ee6fa7-83b1-418d-98e7-6ed99131eb3d",
//	//			"extras":{},
//	//			"msg":"OK",
//	//			"results":[{
//	//				"label":"porn",
//	//				"rate":100.0,
//	//				"scene":"porn",
//	//				"suggestion":"block"}],
//	//			"taskId":"img49S2s7tVXfs4nY0cL8cQnO-1pQYuH",
//	//			"url":"https://img5.njqyjlyh.com/html5/xin/vip1/3.jpg"}],
//	//		"msg":"OK",
//	//		"requestId":"C10BD447-C8A9-43B6-902B-4EF7ACF1BA7B"}`)
//	response, err := scanPron(imgUrl)
//	if err != nil {
//		return &protos.ScanPornReply{}, err
//	}
//	var Response = &models.Response{}
//	err = json.Unmarshal(response, Response)
//	if err != nil {
//
//		return &protos.ScanPornReply{}, err
//	}
//	glog.V(3).Infof("Response = %v \n", Response)
//	if Response.Code != 200 {
//		err = errors.New(fmt.Sprintf("code:%d,msg: %s", Response.Code, Response.Msg))
//		return reply, err
//	}
//	if len(Response.Data) == 0 {
//		err = errors.New("images cannot be analyzed")
//		return reply, err
//	}
//	if len(Response.Data[0].Results) == 0 {
//		err = errors.New("images cannot be analyzed")
//		return reply, err
//	}
//	reply = &protos.ScanPornReply{
//		Link:  imgUrl,
//		Level: Response.Data[0].Results[0].Label,
//	}
//
//	return reply, err
//}

func scanPron(imgUrl string) ([]byte, error) {
	profile := greensdksample.Profile{AccessKeyId: config.MediaConfig.AccessKeyId, AccessKeySecret: config.MediaConfig.AccessHeySecret}

	path := "/green/image/scan"

	clientInfo := greensdksample.ClinetInfo{Ip: "127.0.0.1"}

	// 构造请求数据
	bizType := "Green"
	scenes := []string{"porn"}

	task := greensdksample.Task{DataId: uuid.Rand().Hex(), Url: imgUrl}
	tasks := []greensdksample.Task{task}

	bizData := greensdksample.BizData{bizType, scenes, tasks}

	var client greensdksample.IAliYunClient = greensdksample.DefaultClient{Profile: profile}

	// your biz code
	resp, err := client.GetResponse(path, clientInfo, bizData)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func AuditPorn(imgUrl, imgName string) (string, error) {
	//	AccessKey  cbZkHxsv04xB-AsBajbFAUW4ZFk0GCAV0WtTPlP2
	//  SecretKey  Ohwl8s8zsXEkF3do610pDuOjyTzm7Ty1DTVwqTkn

	//简单上传凭证
	putPolicy := storage.PutPolicy{
		Scope: config.MediaConfig.QiniuApi.Bucket,
	}
	putPolicy.Expires = config.MediaConfig.QiniuApi.Expires
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	//fmt.Println("upToken = ", upToken)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	//err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	//if err != nil {
	//	fmt.Println(err)
	//	return "", nil
	//}
	//fmt.Println("====", ret.Key, ret.Hash)

	req, err := http.NewRequest(
		"GET",
		imgUrl,
		nil,
	)

	num := rand.Intn(len(ippoll.IPPoll))
	hc := ippoll.IPPoll[num]

	resp, err := hc.Do(req)
	if err != nil {
		glog.Errorf("url %s http err %s \n", imgUrl, err)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		glog.V(4).Infof("url: %s,resp.StatusCode: %d\n", imgUrl, resp.StatusCode)
		err = errors.New(config.HTTPERRORCODE)
		return "", err
	}
	//byslice, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//
	//	glog.V(4).Infof("url: %s ,ioutil.ReadAll: %s \n", imgUrl, err)
	//	err = errors.New(config.IMAGEERRORCODE)
	//	return "", err
	//}
	//
	////glog.V(5).Infoln(string(byslice))

	//_, format, err := image.Decode(resp.Body)
	//if err != nil {
	//	glog.Errorf("url: %s image.DecodeConfig err %s", imgUrl, err)
	//	return "", err
	//}
	//glog.V(4).Infof("url: %s，format: %s，resp.ContentLength %d", imgUrl, format,resp.ContentLength)

	err = formUploader.Put(context.Background(), &ret, upToken, imgName, resp.Body, resp.ContentLength, &putExtra)
	if err != nil {
		glog.Errorf("url %s,formUploader.Put %s", imgUrl, err)
	}
	fmt.Println("====", ret.Hash, ret.PersistentID)

	return "Success", nil
}

type OcrIdcardData struct {
	Uri string
}
type Params struct {
	Scenes []string
}

type OcrIdcard struct {
	Data   *OcrIdcardData
	Params *Params
}

func CreateToken(imgUrl string) {
	bodyUri := OcrIdcardData{Uri: imgUrl}
	params := Params{Scenes: []string{"pulp"}}
	body := OcrIdcard{Data: &bodyUri, Params: &params}

	reqData, _ := json.Marshal(body)

	req, reqErr := http.NewRequest(METHOD, URL, bytes.NewReader(reqData))
	if reqErr != nil {
		return
	}

	req.Header.Add("Content-Type", CONTENTTYPE)
	req.Header.Add("Host", config.MediaConfig.QiniuApi.QiniuHost)
	mac := qbox.NewMac(accessKey, secretKey)
	qiniuToken, signErr := mac.SignRequestV2(req)
	if signErr != nil {
		fmt.Printf(signErr.Error())
	}

	req.Header.Add("Authorization", "Qiniu "+qiniuToken)

	fmt.Println("url = ", URL)
	fmt.Println("reqData = ", string(reqData))
	fmt.Println("token = ", string("Qiniu "+qiniuToken))

	resp, respErr := http.DefaultClient.Do(req)
	if respErr != nil {
		fmt.Printf(respErr.Error())
	}
	defer resp.Body.Close()

	resData, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(resData))
}

//func SignRequestV2(mac *qbox.Mac, req *http.Request) (token string, err error) {
//	h := hmac.New(sha1.New, mac.SecretKey)
//
//	u := req.URL
//	//write method path?query
//	//fmt.Println("token = ", fmt.Sprintf("%s %s?%v\nHost: %v\nContent-Type: %s %s %v", req.Method, u.Path, u.RawQuery,req.Host,req.Header.Get("Content-Type"),"\n\n",req.Body))
//	io.WriteString(h, fmt.Sprintf("%s %s", req.Method, u.Path))
//	if u.RawQuery != "" {
//		io.WriteString(h, "?")
//		io.WriteString(h, u.RawQuery)
//	}
//
//	//write host and post
//	io.WriteString(h, "\nHost: ")
//	io.WriteString(h, req.Host)
//	fmt.Printf("h=========== %v \n", string(h.Sum(nil)))
//
//	//write content type
//	contentType := req.Header.Get("Content-Type")
//	if contentType != "" {
//		io.WriteString(h, "\n")
//		io.WriteString(h, fmt.Sprintf("Content-Type: %s", contentType))
//	}
//
//	io.WriteString(h, "\n\n")
//
//	//write body
//	if incBodyV2(req) {
//		s2, err2 := seekable.New(req)
//		if err2 != nil {
//			return "", err2
//		}
//		h.Write(s2.Bytes())
//	}
//	fmt.Printf("h=========== %v \n", string(h.Sum(nil)))
//	sign := base64.URLEncoding.EncodeToString(h.Sum(nil))
//	token = fmt.Sprintf("%s:%s", mac.AccessKey, sign)
//	fmt.Println("token  = ", token)
//	return
//}
//
//func incBodyV2(req *http.Request) bool {
//	contentType := req.Header.Get("Content-Type")
//	return req.Body != nil && (contentType == conf.CONTENT_TYPE_FORM || contentType == conf.CONTENT_TYPE_JSON)
//}

func UrltoPorn(imgUrl string) {

	bodyUri := OcrIdcardData{Uri: imgUrl}
	params := Params{Scenes: []string{"pulp"}}
	body := OcrIdcard{Data: &bodyUri, Params: &params}

	reqData, _ := json.Marshal(body)

	req, reqErr := http.NewRequest(METHOD, URL, bytes.NewReader(reqData))
	if reqErr != nil {
		return
	}

	req.Header.Add("Content-Type", CONTENTTYPE)
	req.Header.Add("Host", HOST)
	mac := qbox.NewMac(accessKey, secretKey)
	qiniuToken, signErr := mac.SignRequestV2(req)
	if signErr != nil {
		fmt.Printf(signErr.Error())
	}

	req.Header.Add("Authorization", "Qiniu "+qiniuToken)

	fmt.Println("url = ", string(URL))
	fmt.Println("reqData = ", string(reqData))
	fmt.Println("token = ", string("Qiniu "+qiniuToken))

	resp, respErr := http.DefaultClient.Do(req)
	if respErr != nil {
		fmt.Printf(respErr.Error())
	}
	defer resp.Body.Close()

	resData, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(resData))
}
