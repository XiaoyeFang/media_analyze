package gogrpc

import (
	"flag"
	"golang.org/x/net/context"
	"pure-media/protos"
	"testing"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "/tmp")
	flag.Set("v", "5")
	flag.Parse()

}
func TestPureMedia_ImgAnalyze(t *testing.T) {
	var puremedia PureMedia
	var ctx context.Context
	req := &protos.ImgRequest{
		//DownloadUrl:"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX3NjcmVlbl8wXzE1Mzc0NzkzNDlfMDY1/screen-0.jpg?h=355&fakeurl=1&type=.jpg",
		//DownloadUrl:"https://apk.302e.com:3443/uploads/-/system/user/avatar/29/avatar.png",
		//DownloadUrl:"http://image.winudf.com/v2/user/comment/MTQ4M19JTUdfMjAxNzEwMTVfMTk0MTQ2LmpwZ18yXzIwMTcxMDMwMTcxOTQ0Njk5.jpg",
		//DownloadUrl: "http://pic1.win4000.com/wallpaper/6/5243c949aadb6.jpg",
		//DownloadUrl:"https://img.pc841.com/2018/0730/20180730081702510.jpg",
		//DownloadUrl:"https://image.winudf.com/v2/user/comment/NTMyNzEwOF9TY3JlZW5zaG90X9mi2aDZodmo2aHZodmi2act2aDZptmg2aTZpNmgLnBuZ18xXzIwMTgxMTI3MDMwNDU4NDQy.jpg",
		//DownloadUrl:"https://image.winudf.com/v2/image/Y29tLmR0cy5mcmVlZmlyZXRoX3NjcmVlbl8wXzE1Mzk2ODA0MjJfMDE5/screen-0.jpg?h=355&fakeurl=1&type=.jpg",
		//大型图片出错url
		DownloadUrl: "http://up.36992.com/pic_source/5a/f0/19/5af01952e75f06cdeb5dedccfba77c70.jpg",
	}
	reply, err := puremedia.ImgAnalyze(ctx, req)
	if err != nil {
		t.Log(err)
	}
	t.Log(reply)
}
func TestPureMedia_ListImgAnalyze(t *testing.T) {
	var puremedia PureMedia
	var ctx context.Context

	req := &protos.ImgRequest{
		//DownloadUrl:"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX3NjcmVlbl8wXzE1Mzc0NzkzNDlfMDY1/screen-0.jpg?h=355&fakeurl=1&type=.jpg",
		DownloadUrl: "https://apk.302e.com:3443/uploads/-/system/user/avatar/29/avatar.png",
	}
	reqs := &protos.ListImgRequest{}
	reqs.Request = append(reqs.Request, req)
	reply, err := puremedia.ListImgAnalyze(ctx, reqs)
	if err != nil {
		t.Log(err)
	}
	t.Log(reply)
}

func TestPureMedia_YoutubeVideo(t *testing.T) {
	var puremedia PureMedia
	var ctx context.Context

	req := &protos.YoutubeVideoRequest{
		Link: "https://www.youtube.com/watch?v=_XwO73wafp4",
		//Link: "https://www.youtube.com/watch?feature=youtu.be&v=skHbLHkS5LE",
	}
	reply, err := puremedia.YoutubeVideo(ctx, req)
	if err != nil {
		t.Log(err)
	}
	t.Log("reply =", reply.ScreenFid)
}
