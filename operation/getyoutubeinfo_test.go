package operation

import (
	"flag"
	"testing"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "/tmp")
	flag.Set("v", "4")
	flag.Parse()

}
func TestGetYoutubeWatchReply(t *testing.T) {
	reply, err := GetYoutubeWatchReply("https://youtu.be/a-6VVwObr3g")
	if err != nil {
		t.Error(err)
	}
	t.Log(reply)

}

func TestGetYouTubeId(t *testing.T) {
	tds := []struct {
		url string
		id  string
		err error
	}{
		{
			//https://www.youtube.com/watch?v=ThY7vEq6DY8
			url: "https://www.youtube.com/watch?v=Uj1EHFY21G4",
			id:  "klo7ZlWV1kU",
			err: nil,
		},
		//{
		//	url: "https://www.youtube.com/watch?feature=youtu.be&v=skHbLHkS5LE",
		//	id:  "skHbLHkS5LE",
		//	err: nil,
		//},
		//{
		//	url: "https://www.youtube.com/embed/skHbLHkS5LE",
		//	id:  "skHbLHkS5LE",
		//	err: nil,
		//},
		//{
		//	url: "https://youtu.be/skHbLHkS5LE",
		//	id:  "skHbLHkS5LE",
		//	err: nil,
		//},
		//{
		//	url: "https://youtu.be/skHbLHkS5LE",
		//	id:  "skHbLHkS5LE",
		//	err: nil,
		//},
		//{
		//	url: "https://youtu.be/v/skHbLHkS5Lr",
		//	id:  "skHbLHkS5Lr",
		//	err: nil,
		//},
		//{
		//	url: "https://www.youtube.com/watch?v=2EuTs1Yo-Bo&index=8&list=PLFbpBUhLEeCyzI-RzcRcCvzJvZNka6hyf&t=0s",
		//},
	}

	for _, v := range tds {
		id, err := GetYouTubeId(v.url)
		if err != nil {
			t.Errorf("GetYouTubeId %s \n", err)
		}
		t.Logf("id %s \n", id)
	}
}

//https://youtu.be/klo7ZlWV1kU
func TestGetRespYouTubeById(t *testing.T) {
	info, err := GetRespYouTubeById("klo7ZlWV1kU")
	if err != nil {
		t.Errorf("err %s \n", err)
	}
	t.Log(len(info))
}

func TestGetVideoMsg(t *testing.T) {
	video, err := GetVideoMsg("5EWqQW1Y_9c")
	if err != nil {
		t.Errorf("GetVideoInfo %s \n", err)
	}
	t.Logf("video %v \n", video.Info)
}

func TestUploadSceenToS3(t *testing.T) {
	//url := "https://i.ytimg.com/vi/skHbLHkS5LE/maxresdefault.jpg"
	url := "https://pic2.zhimg.com/v2-701e1f34a559b273ac178f3c2056c38b_1200x500.jpg"
	id := "pic2.zhimg.com"
	fid, err := UploadSceenToS3(url, id)

	if err != nil {
		t.Error(err)
	}
	t.Log(fid)
}

func TestParseduration(t *testing.T) {
	t.Log(Parseduration("PT29M23S"))
}

func TestHalfSearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	mid := HalfSearch(nums, 0, len(nums)-1, 4)
	t.Log(mid)
}

func TestFirstUniqChar(t *testing.T) {
	str := "aadadaad"
	t.Log(FirstUniqChar(str))
}

func TestIsValidUrl(t *testing.T) {

	url := "https://youtu.be/NC_g066qN9Q"

	t.Log(IsValidUrl(url))
}
