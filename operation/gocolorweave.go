package operation

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/EdlinOrg/prominentcolor"
	"github.com/golang/glog"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/webp"
	"image"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	_ "net/url"
	"os"
	"pure-media/config"
	pb "pure-media/protos"
	"strconv"
)

func GetproDominant(imgUrl string) (msg *pb.ImgReply, err error) {
	msg = &pb.ImgReply{}
	req, _ := http.NewRequest(
		"GET",
		imgUrl,
		nil,
	)

	//2018.11.26  暂时关闭代理
	resp, err := config.NoProxyHttpClient.Do(req)
	if err != nil {
		glog.V(4).Infof("url: %s,http err: %s\n", imgUrl, err)
		err = errors.New(config.HTTPERRORCODE)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		glog.V(4).Infof("url: %s,resp.StatusCode: %d\n", imgUrl, resp.StatusCode)
		err = errors.New(config.HTTPERRORCODE)
		return nil, err
	}

	byslice, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		glog.V(4).Infof("url: %s ,ioutil.ReadAll: %s \n", imgUrl, err)
		err = errors.New(config.IMAGEERRORCODE)
		return nil, err
	}
	img, format, err := image.Decode(bytes.NewReader(byslice))
	if err != nil {
		glog.V(4).Infof("url: %s,image.Decode err : %s \n", imgUrl, err)
		err = errors.New(config.IMAGEERRORCODE)
		return nil, err
	}
	item, err := prominentcolor.KmeansWithAll(1, img, 0, 0, nil)
	if err != nil {
		return nil, err
	}
	msg.Url = imgUrl
	msg.Size = int64(len(byslice))
	msg.Width = int64(img.Bounds().Dx())
	msg.Height = int64(img.Bounds().Dy())
	msg.Format = format
	//R, G, B uint32
	if len(item) != 0 {
		r := item[0].Color.R
		g := item[0].Color.G
		b := item[0].Color.B
		msg.ColorExtractorRGB = fmt.Sprintf("RGB(%d,%d,%d)", r, g, b)
		//color.RGBToCMYK()
		//强转有时会有偏差
		msg.ColorExtractorHEX = fmt.Sprintf("#%s", t2x(r)+t2x(g)+t2x(b))
	}

	msg.ColorModel = fmt.Sprint(img.Bounds().ColorModel())

	if format == "gif" {
		//msg.FrameNumber
		gif, err := gif.DecodeAll(bytes.NewReader(byslice))
		if err != nil {
			glog.Errorf("url: %s,gif.DecodeAll: %s \n", imgUrl, err)
			return msg, err
		}
		//fmt.Printf("HZ %d \n", gif.LoopCount)
		msg.FrameNumber = int64(gif.LoopCount)

	}
	//glog.V(2).Infof("msg %v \n", msg)
	return msg, err
}

func t2x(t uint32) string {
	result := strconv.FormatInt(int64(t), 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

func GetDominant(picture []byte) (msg *pb.ImgReply, err error) {
	msg = &pb.ImgReply{}

	reader := bytes.NewReader(picture)
	img, format, err := image.Decode(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)

		return msg, err
	}
	item, err := prominentcolor.KmeansWithAll(1, img, 0, 0, nil)
	if err != nil {
		return msg, err
	}

	msg.Size = int64(len(picture))
	msg.Width = int64(img.Bounds().Dx())
	msg.Height = int64(img.Bounds().Dy())
	msg.Format = format
	//R, G, B uint32
	if len(item) != 0 {
		r := item[0].Color.R
		g := item[0].Color.G
		b := item[0].Color.B
		msg.ColorExtractorRGB = fmt.Sprintf("RGB(%d,%d,%d)", r, g, b)
	}

	msg.ColorModel = fmt.Sprint(img.Bounds().ColorModel())

	if format == "gif" {
		//msg.exif

		gif, err := gif.DecodeAll(bytes.NewReader(picture))
		if err != nil {
			glog.Errorf("gif.DecodeAll %s \n", err)
		}
		//glog.Errorf("HZ %d \n", gif.LoopCount)
		msg.FrameNumber = int64(gif.LoopCount)

		return msg, err

	}

	return msg, err
}
