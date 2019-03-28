package database

import (
	"database/sql"
	"fmt"
	"github.com/golang/glog"
	"pure-media/config"
	"pure-media/models"
	"strings"
)

/*
upload_file_info 表结构
file_id，S3文件ID
file_size，文件大小，单位：Bytes
file_sha256，文件哈希
file_type, 文件类型image
image_format，图片类型
image_width，图片宽度
image_height，图片高度
image_color_extractor，图片颜色主色调
tags, 标签方便搜索
context，上下分，JSON类型
created_at，创建时间
*/
const (
	MaxChar = 255
)

func init() {
	//	TODO 2019.3.6数据库表结构变动
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

type PostGreDB struct {
	db *sql.DB
}

func NewDB() *PostGreDB {
	db, err := config.CreateTube()
	if err != nil {
		panic(err)
	}
	return &PostGreDB{db}
}

//select count(*)from information_schema.columns where table_name = 't_zxxs_aj' and column_name='c_ah'
//检测列是否存在
func (self *PostGreDB) IsExistColumn(column string) (err error) {
	queryStr := "select count(*) from information_schema.columns where table_schema='table_schema' and table_name ='crawler_tube' and column_name=" + column
	fmt.Println(queryStr)
	stmt, err := self.db.Prepare(queryStr)
	if err != nil {
		glog.Errorf("image Prepare err %s \n", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		glog.Errorf("image Exec err %s \n", err)
		return err
	}
	return nil

}

//alter table crawler_tube
//	add tube_fid 		varchar  default ''::character  varying  NOT NULL
func (self *PostGreDB) AddColumn(coulmn, value string) (err error) {
	//stmt, err := self.db.Query("ALTER table crawler_tube ADD COLUMN_NAME $1 VARCHAR DEFAULT $2")
	var queryStr string
	//if value == "" {
	//	queryStr = "ALTER table crawler_tube ADD " + coulmn + " varchar "
	//} else {
	queryStr = "ALTER table crawler_tube ADD " + coulmn + " varchar  default " + "'" + value + "'"
	//}
	//fmt.Println("queryStr", queryStr)
	stmt, err := self.db.Prepare(queryStr)
	if err != nil {
		glog.Errorf("image Prepare err %s \n", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		//glog.Errorf("%s", err)
		if strings.Contains(err.Error(), "already exists") {
			return nil
		}

	}
	return nil
}

//func (self *PostGreDB) SaveUploadFile(info *models.UploadFileInfo) (err error) {
//
//	defer self.db.Close()
//
//	//根据文件类型存储文件信息
//	//switch info.FileType {
//	//case config.VIDEO:
//	//	/*
//	//		youtubeId，视频ID
//	//		title，视频标题
//	//		lengthSeconds，视频播放长度（秒）
//	//		screenImage，视频截图
//	//		keywords，视频关键词
//	//		hl，视频发布语言
//	//		author，视频作者
//	//		rating，评分
//	//		viewCount，播放量
//	//		definition，视频清晰度
//	//	*/
//	//
//	//case config.IMAGE:
//	//10.23 添加HEX颜色
//	//将Tag按","为间隔转换为字符串
//	tags := strings.Join(info.Tags, ",")
//	//glog.Errorf("image UrlUploadFile %v \n", info)
//	queryStr := "INSERT INTO upload_file_info (file_id,file_size,file_sha256,file_type,image_format,image_width,image_height,color_extractor_rgb,color_extractor_hex,tags,created_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"
//	_, err = self.db.Query(queryStr,
//		info.FileId,
//		info.FileSize,
//		info.FileSha256,
//		info.FileType,
//		info.ImageFormat,
//		info.ImageWidth,
//		info.ImageHeight,
//		info.ColorExtractorRGB,
//		info.ColorExtractorHEX,
//		tags,
//		info.CreatedAt)
//	if err != nil {
//		glog.Errorf("image INSERT err %s \n", err)
//		return err
//	}
//
//	//case config.AUDIO:
//	//
//	//default:
//	//	//将Tag按","为间隔转换为字符串
//	//	glog.Errorf("default UrlUploadFile %v \n", info)
//	//	tags := strings.Join(info.Tags, ",")
//	//	_, err = self.db.Query("INSERT INTO upload_file_info (file_id,file_size,file_sha256,file_type,tags,created_at) VALUES($1,$2,$3,$4,$5,$6)",
//	//		info.FileId,
//	//		info.FileSize,
//	//		info.FileSha256,
//	//		info.FileType,
//	//		tags,
//	//		info.CreatedAt)
//	//	if err != nil {
//	//		glog.Errorf("default INSERT err %s \n", err)
//	//		return err
//	//	}
//	//}
//
//	return err
//}

func (self *PostGreDB) UpdateInfo(info *models.VideoMsg) (err error) {

	//stmt,err :=self.db.Prepare("SELECT tube_fid FROM crawler_tube WHERE col IS NULL")
	//if err != nil {
	//	glog.Errorf("db.Prepare err %s \n",err)
	//	return err
	//}
	//defer stmt.Close()
	//_,err =stmt.Exec()
	//if err != nil {
	//	glog.Errorf("image Exec err %s \n", err)
	//}

	stmt, err := self.db.Prepare("UPDATE  crawler_tube SET tube_fid =$1 WHERE video_id=$2 ")
	if err != nil {
		glog.Errorf("db.Prepare err %s \n", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(info.TubeFid, info.VideoId)
	if err != nil {
		glog.Errorf("image UPDATE err %s \n", err)
		return err
	}
	glog.V(5).Infoln("update crawler_tube success")
	return nil
}

func (self *PostGreDB) SaveTubeInfo(info *models.VideoMsg) (err error) {
	//defer self.db.Close()
	//将Tag按","为间隔转换为字符串
	KeyWords := strings.Join(info.KeyWords, ",")
	//glog.Errorf("image UrlUploadFile %v \n", len(KeyWords))
	if len(KeyWords) > MaxChar {

		KeyWords = KeyWords[:MaxChar]
	}
	queryStr := "INSERT INTO crawler_tube (video_id,title,length_seconds,screen_fid,screen_youtube_url,key_worlds,info,who,created) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	_, err = self.db.Query(queryStr,
		info.VideoId,
		info.Title,
		info.LengthSeconds,
		info.ScreenFid,
		info.ScreenYoutubeUrl,
		KeyWords,
		info.Info,
		info.Who,
		info.Created,
		info.TubeFid,
		info.NewTubeId,
		info.Type,
		info.CopyState)
	if err != nil {
		glog.Errorf("image INSERT err %s \n", err)
		return err
	}

	return err
}

//查重
func (self *PostGreDB) Rechecking(vedioId string) (info *models.VideoMsg, err error) {
	//defer self.db.Close()
	info = &models.VideoMsg{}
	var KeyWords string
	rows, err := self.db.Query("SELECT * FROM crawler_tube WHERE video_id = $1", vedioId)
	if err != nil {
		glog.Errorf("Rechecking Query %s \n", err)
	}
	//glog.Infoln("rows == ", rows)

	if rows.Next() {
		//查询结果不为空
		//err = errors.New("record is exist")
		//return err
		//	var item storage.Item
		//	err := self.db.QueryRow("SELECT url, visited, count FROM shortener where shortener=$1 limit 1", code).
		//		Scan(&item.URL, &item.Visited, &item.Count)
		err = rows.Scan(&info.Id, &info.VideoId, &info.Title, &info.LengthSeconds, &info.ScreenFid, &info.ScreenYoutubeUrl,
			&KeyWords, &info.Info, &info.Who, &info.Created, &info.TubeFid, &info.NewTubeId, &info.Type, &info.CopyState)
		if err != nil {
			glog.Errorf("rows.Scan %s \n", err)
		}
	}
	info.KeyWords = strings.Split(KeyWords, ",")

	return info, err
}

func (self *PostGreDB) Close() {
	self.db.Close()
}
