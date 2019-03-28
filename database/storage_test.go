package database

import (
	"flag"
	"github.com/golang/glog"
	"pure-media/models"
	"testing"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "/tmp")
	flag.Set("v", "5")
	flag.Parse()

}
func TestNewDB(t *testing.T) {
	db := NewDB()
	err := db.AddColumn("tube_fid", "")
	if err != nil {
		glog.Errorf("AddColumn  tube_fid err %s \n", err)

	}
	err = db.AddColumn("new_tube_id", "")
	if err != nil {
		glog.Errorf("AddColumn  tube_fid err %s \n", err)

	}
	err = db.AddColumn("type", "COPY")
	if err != nil {
		glog.Errorf("AddColumn  tube_fid err %s \n", err)
	}
	err = db.AddColumn("copy_state", "DOWNLOADING")
	if err != nil {
		glog.Errorf("AddColumn  tube_fid err %s \n", err)
	}

}

func TestPostGreDB_IsExistColumn(t *testing.T) {

	db := NewDB()
	err := db.IsExistColumn("title")
	if err != nil {
		t.Error(err)
	}
}

func TestPostGreDB_AddColumn(t *testing.T) {
	db := NewDB()
	err := db.AddColumn("type", "COPY")
	if err != nil {
		t.Error(err)
	}
}

func TestPostGreDB_UpdateInfo(t *testing.T) {

	db := NewDB()
	//info := &models.UploadFileInfo{FileId:"111111",FileType: "image", ImageFormat: "jpg", CreatedAt: time.Now().String()}
	info := &models.VideoMsg{}
	err := db.UpdateInfo(info)
	if err != nil {
		t.Error(err)
	}
}

func TestPostGreDB_SaveUploadFile(t *testing.T) {
	db := NewDB()
	//info := &models.VideoMsg{FileId:"111111",FileType: "image", ImageFormat: "jpg", CreatedAt: time.Now().String()}
	info := &models.VideoMsg{TubeFid: "11111"}
	err := db.SaveTubeInfo(info)
	if err != nil {
		t.Error(err)
	}
}

func TestPostGreDB_Rechecking(t *testing.T) {
	db := NewDB()
	//info := &models.UploadFileInfo{FileId: "puremedia/pureY3ZQQnJZZ2tScXJtM21ldUtCYjJRcldOZU93X2ltYWdlXzI3MTcyNg", FileType: "image", ImageFormat: "jpg", CreatedAt: time.Now().String()}

	id := "5EWqQW1Y_9c"
	msg, err := db.Rechecking(id)
	if err != nil {
		t.Errorf("Rechecking %s \n", err)
	}
	t.Log(msg)
}
