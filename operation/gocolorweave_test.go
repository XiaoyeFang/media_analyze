package operation

import (
	"io/ioutil"
	"testing"
)

func init() {
	//flag.Set("alsologtostderr", "true")
	//flag.Set("log_dir", "/tmp")
	//flag.Set("v", "5")
	//flag.Parse()

}

func TestGetproDominant(t *testing.T) {
	//url := "https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9"
	//url := "https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX3NjcmVlbl8wXzE1Mzc0NzkzNDlfMDY1/screen-0.jpg?h=355&fakeurl=1&type=.jpg"
	//url := "http://www.52inwet.com/wp-content/uploads/2013/05/041.gif"
	//url := "https://leetcode-cn.com/problems/convert-sorted-array-to-binary-seach-tree/"
	//url := "http://imgfit.staging.winudf.com/v2/user/comment/MTc1ODExMF9TY3JlZW5zaG90XzIwMTctMDQtMDctMDItMTItNDQucG5nXzIwMTcwNDA5MTEyNTEzNzY3.pnghttp:"
	//url := "http://imgfit.staging.winudf.com/v2/user/comment/YWRtaW5fY29tbWVudF8xOTIuMTY4LjkuMTIyXzE2MDFfZmFjZWJvb2tfY29tLmZhY2Vib29rLmthdGFuYV90PTEoTmV4dXMgNCkucG5nXzE1NDA5NTU5NjQzNjE"
	//url :="https://image.winudf.com/v2/user/comment/dW5kZWZpbmVkX0lNR18yMDE4MTEyMF8yMDA1MDkuanBnXzFfMjAxODEyMDMwNjI4NDc2Njk.jpg"
	//url := "https://image.winudf.com/v2/user/comment/dW5kZWZpbmVkX1NjcmVlbnNob3RfMjAxOC0xMC0zMS0yMS01Ny00MS0xMzk1ODA3NjkwLnBuZ18xXzIwMTgxMjAzMDYyMjQwODUw.png"
	//url :="https://thumbs.gfycat.com/PhysicalBrightArcherfish-size_restricted.gif"
	//url := "https://image.winudf.com/v2/user/comment/NTkwMV9odW50ZXItaHVudGVkMS5qcGdfMV8yMDE3MDgzMTIzNDU1MDkyMQ.jpg"
	url := "http://image.winudf.com/v2/user/comment/NDc1OTAwMV9JTUdfMjAxNzEyMjFfMjEzODE3X0hEUi5qcGdfMV8yMDE3MTIyMjA3NDI1MDc2NQ.jpg"
	reply, err := GetproDominant(url)
	if err != nil {
		t.Errorf("GetproDominant %s \n", err)
	}
	t.Log(reply)
}

func TestGetDominant(t *testing.T) {
	bt, err := ioutil.ReadFile("images/icon.png")
	if err != nil {
		t.Error(err)
	}
	t.Log(len(bt))

	msg, err := GetDominant(bt)
	if err != nil {
		t.Error(err)
	}
	t.Log(msg)
}
