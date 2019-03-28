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
func TestGetDownloadLink(t *testing.T) {
	t.Log(len("[download] "))
	t.Log(len(" has already been downloaded and merged"))
}

func TestDownloadVideo(t *testing.T) {
	filename, err := DownloadVideo("https://youtu.be/a-6VVwObr3g", "137+140")
	if err != nil {
		t.Errorf("DownloadVideo %s", err)
	}
	//if isexist := Exists("Warframe _ Nightwave - Series 1 Launch Trailer -- Out now on ALL platforms-a-6VVwObr3g.mp4"); !isexist {
	//	t.Errorf("Exists %v", isexist)
	//}
	t.Log(filename)
}

func TestUploadAccount(t *testing.T) {
	UploadAccount()
}

func TestYtubeupload(t *testing.T) {
	Ytubeupload()
}
