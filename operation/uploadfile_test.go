package operation

import (
	"flag"
	"io/ioutil"
	"path"
	"pure-media/protos"
	"strings"
	"testing"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "/tmp")
	flag.Set("v", "5")
	flag.Parse()

}

func TestSaveFileInfo(t *testing.T) {
	urls := []string{
		//"http://img0.utuku.china.com/212x0/news/20170612/c3b2a11e-28ac-4b50-b295-8e6c6ba7a7ed.jpg",
		//"https://image.winudf.com/v2/image/Y29tLnBhcmFub2lkam95LmF0dGJfYmFubmVyXzE1NDAzMzA3NDZfMDI3/banner.jpg?w=850&fakeurl=1&type=.jpg",
		//"https://image.winudf.com/v2/image/Y29tLmF1ZXIuZGF3bmJyZWFrLm5pZ2h0LndpdGNoLnNpbmdsZS5wbGF5ZXIuZnJlZS5zZXF1ZWxfYmFubmVyXzE1MzczODY5ODBfMDgw/banner.jpg?w=850&fakeurl=1&type=.jpg",
		//"https://lh3.googleusercontent.com/2Z8pBEXe14rLRmQ1jt92i6xJQW2jaU3c9MfR5KGpIA2tUV2Fs2RnS5DSDbv6CMEZsTXl=w200-h300-rw",
		//"https://image.winudf.com/v2/image/Y29tLmF1ZXIuZGF3bmJyZWFrLm5pZ2h0LndpdGNoLnNpbmdsZS5wbGF5ZXIuZnJlZS5zZXF1ZWxfYmFubmVyXzE1MzczODY5ODBfMDgw/banner.jpg?w=850&fakeurl=1&type=.jpg",
		//"https://lh3.googleusercontent.com/2Z8pBEXe14rLRmQ1jt92i6xJQW2jaU3c9MfR5KGpIA2tUV2Fs2RnS5DSDbv6CMEZsTXl=w200-h300-rw",
		//"https://image.winudf.com/v2/image/Y29tLmF1ZXIuZGF3bmJyZWFrLm5pZ2h0LndpdGNoLnNpbmdsZS5wbGF5ZXIuZnJlZS5zZXF1ZWxfYmFubmVyXzE1MzczODY5ODBfMDgw/banner.jpg?w=850&fakeurl=1&type=.jpg",
		//"https://img-blog.csdn.net/20160413112832792?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQv/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/Center",
		//"https://image.winudf.com/v2/image/Y29tLmdhbWV2aWwuZXMyLmFuZHJvaWQuZ29vZ2xlLmdsb2JhbC5ub3JtYWxfYmFubmVyXzE1Mzk2ODg5NjNfMDQ3/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://i.ytimg.com/vi/2EuTs1Yo-Bo/maxresdefault.jpg",
	}
	c := make(chan int, 10)
	c1 := make(chan int, len(urls))
	for _, v := range urls {
		c <- 1
		req := &protos.UploadFileRequest{
			FileUrl:  v,
			FileType: "image",
			Tags:     []string{"ssssss", "eeeeeeee"},
			Bucket:   "puremedia",
			Prefix:   "pure",
		}

		go func(req *protos.UploadFileRequest) {
			reply, err := SaveFileInfo(req.FileUrl, req.FileType, req.Bucket, req.Prefix, req.Tags)
			if err != nil {
				t.Errorf("SaveFileInfo %s \n", err)
			}
			t.Logf("reply %v\n", reply)
			c1 <- 1
			<-c
		}(req)
	}

	for i := 0; i < len(urls); i++ {
		<-c1
	}

}

func TestUrlUploadFile(t *testing.T) {
	url := "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	//url := "http://imgfit2.default/v2/image/Y29tLmdhbWV2aWwuZXMyLmFuZHJvaWQuZ29vZ2xlLmdsb2JhbC5ub3JtYWxfYmFubmVyXzE1Mzk2ODg5NjNfMDQ3/banner.jpg?w=850&fakeurl=1&type=.jpg"
	fileType := "image"
	tags := []string{"ssssss", "bbbbbb"}
	bucket := "puremedia"
	prefix := "pure"

	fileInfo, err := UrlUploadFile(url, fileType, bucket, prefix, tags)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v \n", fileInfo)

}

func TestUploadFile(t *testing.T) {
	//url := "https://image.winudf.com/v2/image/Y29tLmdhbWV2aWwuZXMyLmFuZHJvaWQuZ29vZ2xlLmdsb2JhbC5ub3JtYWxfYmFubmVyXzE1Mzk2ODg5NjNfMDQ3/banner.jpg?w=850&fakeurl=1&type=.jpg"
	byteslice, err := ioutil.ReadFile("images/betaicon.png")
	if err != nil {
		t.Log(err)
	}

	fileInfo, err := UploadFile(byteslice, "image", []string{"111"})
	if err != nil {
		t.Errorf("UrlUploadFile %s \n", err)
	}
	t.Logf("fileinfo = %v \n", fileInfo)
}

func GetFileType(link string) string {
	if link == "" {
		return ""
	}

	ret := strings.Split(link, "/")
	if len(ret) < 2 {
		return ""
	}
	return path.Ext(ret[len(ret)-1])
}
