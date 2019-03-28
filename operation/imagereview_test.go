package operation

import (
	"flag"
	"testing"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "/tmp")
	flag.Set("v", "5")
	flag.Parse()

}

func TestImgScanPorn(t *testing.T) {
	//正常
	//normalUrl := "http://imgfit.staging.winudf.com/v2/user/comment/YWRtaW5fY29tbWVudF8xOTIuMTY4LjkuMTIyXzE2MDFfZmFjZWJvb2tfY29tLmZhY2Vib29rLmthdGFuYV90PTEoTmV4dXMgNCkucG5nXzE1NDA5NTU5NjQzNjE"
	//normalUrl:="https://image.winudf.com/v2/user/comment/dW5kZWZpbmVkXzE1NDMyODY5Njk2MTkuanBnXzFfMjAxODExMjcwMjUwMDk0NTE.png"
	//normalUrl := "https://pic2.zhimg.com/v2-529d6c4340c5a97f77b50fbddf5f3be1_200x112.jpeg"
	//normalUrl :="https://image.winudf.com/v2/user/comment/dW5kZWZpbmVkX21jcGVDZW50ZXJfMjAxODExMjcwOTUzMjYuanBnXzFfMjAxODExMjcwMjU0MjUzNjU.jpg"
	//性感
	//sexyUrl := "https://sjbz-fd.zol-img.com.cn/t_s208x312c5/g5/M00/07/03/ChMkJljlp7mIVS74AAZe51VcP4AAAbZEQJ0SDoABl7_286.jpg"
	//色情
	//pornUrl := "https://img5.njqyjlyh.com/html5/xin/vip1/3.jpg"
	//pornUrl := "https://pic4.zhimg.com/80/de301939599712c8b4684bfedb309edc_hd.jpg"
	//reply, err := ImgScanPorn(normalUrl)
	//if err != nil {
	//	t.Errorf("ImgScanPorn %s \n", err)
	//}
	//t.Logf("reply = %v \n", reply)
}

func TestAuditPorn(t *testing.T) {
	ember := "https://huiji-thumb.huijistatic.com/warframe/uploads/thumb/3/39/EmberPrimeNewLook.png/260px-EmberPrimeNewLook.png"
	ember = "https://pic3.zhimg.com/v2-e5d27fab4d17bb40c6fdb5682c802e23_b.jpg"
	AuditPorn(ember, "ember.png")
}

func TestCreateToken(t *testing.T) {
	ember := "https://huiji-thumb.huijistatic.com/warframe/uploads/thumb/3/39/EmberPrimeNewLook.png/260px-EmberPrimeNewLook.png"
	CreateToken(ember)
}

func TestUrltoPorn(t *testing.T) {
	ember := "https://huiji-thumb.huijistatic.com/warframe/uploads/thumb/3/39/EmberPrimeNewLook.png/260px-EmberPrimeNewLook.png"

	UrltoPorn(ember)
}
