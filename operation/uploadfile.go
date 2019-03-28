package operation

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang/glog"
	_ "golang.org/x/image/webp"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"pure-media/config"
	"pure-media/models"
	pb "pure-media/protos"
	"pure-media/storage"
	"time"
)

func SaveFileInfo(fileUrl, fileType, bucket, prefix string, tags []string) (reply *pb.UploadFileReply, err error) {
	//db := database.NewDB()

	switch fileType {
	case config.IMAGE:

	case config.VIDEO:
		videoId, err := GetYouTubeId(fileUrl)
		if err != nil {
			glog.Errorf("GetYouTubeId %s \n", err)
			return reply, err
		}
		vodeiMsg, err := GetVideoMsg(videoId)
		if err != nil {
			glog.Errorf("GetVideoMsg %s \n", err)
			return reply, err
		}
		fileUrl = vodeiMsg.ScreenYoutubeUrl
		fmt.Printf("ScreenImage URL %s \n", fileUrl)
	default:

	}

	fileInfo, err := UrlUploadFile(fileUrl, fileType, bucket, prefix, tags)
	if err != nil || fileInfo.FileId == "" {
		return reply, err
	}
	//查重并存储
	//if db.Rechecking(id) == nil {
	//err = db.SaveUploadFile(fileInfo)
	//if err != nil {
	//	return reply, err
	//}
	//}

	reply = &pb.UploadFileReply{
		FileId:            fileInfo.FileId,
		FileSha256:        fileInfo.FileSha256,
		FileType:          fileInfo.FileType,
		FileSize:          fileInfo.FileSize,
		ImageFormat:       fileInfo.ImageFormat,
		ImageWidth:        fileInfo.ImageWidth,
		ImageHeight:       fileInfo.ImageHeight,
		ColorExtractorRGB: fileInfo.ColorExtractorRGB,
		ColorExtractorHEX: fileInfo.ColorExtractorHEX,
		Tags:              fileInfo.Tags,
		CreatedAt:         fileInfo.CreatedAt,
	}

	glog.V(4).Infof("fileInfo %v \n", fileInfo)
	return reply, err
}

//根据buff上传文件
func UploadFile(file []byte, fileType string, tags []string) (fileInfo *models.UploadFileInfo, err error) {
	//glog.Errorf("UrlUploadFile %s \n", fileType)
	//fileInfo = &models.UploadFileInfo{}
	//if len(file) == 0 {
	//	err = errors.New(config.NOT_FOUND)
	//	return fileInfo, err
	//}
	//
	////根据唯一哈希生成S3key
	//fileInfo.FileSha256 = base64.RawURLEncoding.EncodeToString(file[:20])
	//fmt.Println("fileInfo.FileSha256 = ", fileInfo.FileSha256)
	//
	////上传文件到S3获取fid
	//fid, err := S3Upload(file, int64(len(file)), fileType, fileInfo.FileSha256)
	//if err != nil {
	//	return fileInfo, err
	//}
	////文件大小
	//fileInfo.FileId = fid
	//fileInfo.FileSize = int64(len(file))
	//fileInfo.Tags = tags
	//fileInfo.CreatedAt = time.Now().String()
	//switch fileType {
	//case config.IMAGE:
	//	//获取图片信息
	//	img, format, err := image.Decode(bytes.NewReader(file))
	//
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "%s \n", err)
	//
	//		return fileInfo, err
	//	}
	//	item, err := prominentcolor.KmeansWithAll(1, img, 0, 0, nil)
	//	if err != nil {
	//		return fileInfo, err
	//	}
	//	fileInfo.FileType = fileType
	//	fileInfo.ImageFormat = format
	//	fileInfo.ImageWidth = int64(img.Bounds().Dx())
	//	fileInfo.ImageHeight = int64(img.Bounds().Dy())
	//
	//	//R, G, B uint32
	//	if len(item) != 0 {
	//		r := item[0].Color.R
	//		g := item[0].Color.G
	//		b := item[0].Color.B
	//		fileInfo.ColorExtractorRGB = fmt.Sprintf("RGB(%d,%d,%d)", r, g, b)
	//	}
	//case config.VIDEO:
	//
	//case config.AUDIO:
	//
	//default:
	//
	//}

	return fileInfo, err
}

//根据URL上传文件到S3并获取文件信息
func UrlUploadFile(fileUrl, fileType, bucket, prefix string, tags []string) (fileInfo *models.UploadFileInfo, err error) {
	fileInfo = &models.UploadFileInfo{}
	if fileUrl == "" {
		err = errors.New(config.INVALID_URL)
		return fileInfo, err
	}
	//purl, _ := url.Parse(config.MediaConfig.ProxyHttp)
	//proxy := http.ProxyURL(purl)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: nil,
		},
	}

	req, err := MakeCheckUrl(fileUrl)
	if err != nil {
		return fileInfo, err

	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		glog.V(0).Infof("client.Get %s\n", err)
		return fileInfo, err
	}
	glog.V(5).Infof("open %s ret: %v", fileUrl, resp)

	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("%d", resp.StatusCode))
		return fileInfo, err
	}

	b := bytes.Buffer{}
	b.ReadFrom(resp.Body)
	//根据唯一哈希生成S3key
	h := sha1.New()
	h.Write([]byte(fileUrl))
	fileInfo.FileSha256 = base64.RawURLEncoding.EncodeToString(h.Sum([]byte{}))
	glog.V(2).Infof("fileInfo.FileSha256 = %s \n", fileInfo.FileSha256)

	switch fileType {
	case config.IMAGE:
		//获取图片信息
		//上传文件到S3获取fid
		glog.V(3).Infof("bucket %s  prefix %s\n", bucket, prefix)

		fid, err := s3Upload(b.Bytes(), int64(len(b.Bytes())), fileType, fileInfo.FileSha256, bucket, prefix)
		if err != nil {
			return fileInfo, err
		}
		imageInfo, err := GetImageInfo(b.Bytes())
		if err != nil {
			glog.V(2).Infof("fileInfo = %v \n", fileInfo)
			return fileInfo, err

		}
		//文件大小
		fileInfo = &models.UploadFileInfo{
			FileId:            fid,
			FileSha256:        fileInfo.FileSha256,
			FileType:          fileType,
			FileSize:          int64(len(b.Bytes())),
			ImageFormat:       imageInfo.ImageFormat,
			ImageWidth:        imageInfo.ImageWidth,
			ImageHeight:       imageInfo.ImageHeight,
			ColorExtractorRGB: imageInfo.ColorExtractorRGB,
			ColorExtractorHEX: imageInfo.ColorExtractorHEX,
			Tags:              tags,
			CreatedAt:         time.Now().String(),
		}

		glog.V(2).Infof("fileInfo = %v \n", fileInfo)
	case config.VIDEO:

	case config.AUDIO:

	default:

	}

	return fileInfo, err
}

//上传文件
func s3Upload(file []byte, len int64, fileType, fileSha256, bucket, prefix string) (fid string, err error) {

	//config.MediaConfig.S3Conf.Bucket = bucket
	//config.MediaConfig.S3Conf.Prefix = prefix
	//glog.V(3).Infof("bucket %s  prefix %s\n", bucket, prefix)

	fs := storage.NewS3Storage(&config.MediaConfig.S3ConfImage)

	//fs.Conf.Prefix = prefix
	//fs.Conf.Bucket = bucket
	//检查bucket是否存在
	_, err = fs.LsDir(bucket)
	if err != nil {
		glog.Errorf("fs.LsDir %s \n", err)
	}
	glog.Errorf("files == %v \n", fs.Svc.Bucket(bucket))

	key := fmt.Sprintf("%s_%s_%03d", fileSha256, fileType, len)
	key = base64.RawURLEncoding.EncodeToString([]byte(key))
	//fid, err = fs.CopyFileByUrl(url, key)
	fid, err = fs.CopyFileByReader(bytes.NewReader(file), len, fileType, key)
	if err != nil {
		return "", err
	}
	if fs.Conf.Bucket == "" {
		//never do
		glog.Errorf("fs err:%v", fs)
	}
	fid = fmt.Sprintf("%s/%s", fs.Conf.Bucket, fid)
	glog.V(4).Infof("key:%v fid:%s", key, fid)
	return fid, err
}

func MakeCheckUrl(link string) (*http.Request, error) {
	if !IsValidUrl(link) {
		err := errors.New(config.INVALID_URL)
		return &http.Request{}, err
	}

	req, err := http.NewRequest(http.MethodGet, link, nil)

	if err != nil {
		return nil, err
	}
	//req.Header.Set("Accept", config.HTTP_ACCEPT)
	//req.Header.Set("User-Agent", config.HTTP_USER_AGENT)

	return req, err

}
