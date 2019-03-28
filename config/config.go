package config

import (
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	iot "io/ioutil"
	"pure-media/common"
	"pure-media/models"
	"strings"
)

const (
	HTTP_USER_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.90 Safari/537.36"
	HTTP_ACCEPT     = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
	IMAGE           = "image"
	VIDEO           = "video"
	AUDIO           = "audio"
	INVALID_URL     = "INVALID URL"
	NOT_FOUND       = "NOT_FOUND"
	HTTPERRORCODE   = "100"
	IMAGEERRORCODE  = "110"
)

//无代理HC
var NoProxyHttpClient = common.NewHttpClient("http://172.16.0.18:10081")
var MediaConfig *models.Config
var configStr = []byte(`grpc_listen: :5003
proxy_http: ["http://172.16.0.18:10081"]
loglevel: 5
postgres:
  host: '127.0.0.1'
  port: '5432'
  user: 'postgres'
  password: 'postgres'
  db: 'postgres'
access_key_id: ""
access_key_secret: ""
yt_api_key: AIzaSyCe1fzSMGP2z16W2_FVp3JA-cJ7_gbbG6k
yt_api_part: snippet,statistics,contentDetails
yt_api_fields: '*'
s3_conf_image:
  endpoint: "http://s3.staging.xfreeapp.com:80"
  access_key: ""
  secret_key: ""
  bucket: "puremedia"
  prefix: "pure"
  s3acl: "authenticated-read"
  chunk_size: 100MB
s3_conf_youtube:
  endpoint: "http://s3.staging.xfreeapp.com:80"
  access_key: "DO7ZFVUDDPP5D3IU7DQE"
  secret_key: "8SpTQhLTbO25XCSz9KhLpsbyGm2yS4iD8178PB1u"
  bucket: user
  prefix: youtube
  s3acl: private
  chunk_size: 100MB
qiniu_api:
  qiniu_url:  "http://ai.qiniuapi.com/v3/image/censor"
  qiniu_akey: "cbZkHxsv04xB-AsBajbFAUW4ZFk0GCAV0WtTPlP2"
  qiniu_skey: "Ohwl8s8zsXEkF3do610pDuOjyTzm7Ty1DTVwqTkn"
  qiniu_host: "ai.qiniuapi.com"
  bucket: "pure"`)

func init() {
	//glog.Infoln("当前版本 2.0.3")
	var err error

	if MediaConfig == nil {
		MediaConfig, err = LoadConf("/conf/app.yml")
		if err != nil {
			panic(err)
		}
	}
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "/tmp")
	flag.Set("v", MediaConfig.LogLevel)
	flag.Parse()

	//	TODO 数据库表结构变动

}

func CreateUpfile() (*sql.DB, error) {
	//MediaConfig := CreateConfig()
	//fmt.Println("MediaConfig", MediaConfig)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		MediaConfig.Postgres.Host, MediaConfig.Postgres.Port, MediaConfig.Postgres.User, MediaConfig.Postgres.Password, MediaConfig.Postgres.DB)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	strQuery := `CREATE TABLE IF NOT EXISTS upload_file_info (id serial PRIMARY KEY NOT NULL, file_id VARCHAR  NOT NULL,file_size INTEGER not NULL,
file_sha256 VARCHAR,file_type VARCHAR,image_format VARCHAR ,image_width INTEGER,image_height INTEGER , color_extractor_rgb VARCHAR,color_extractor_hex VARCHAR,
tags VARCHAR not NULL,created_at VARCHAR not NULL,count INTEGER DEFAULT 0);`
	_, err = db.Exec(strQuery)
	if err != nil {
		fmt.Println("db.Exec  ", err)
	}

	glog.V(5).Infoln("Successfully connected!")
	return db, nil
}

func CreateTube() (*sql.DB, error) {
	//MediaConfig := CreateConfig()
	//fmt.Println("MediaConfig", MediaConfig)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		MediaConfig.Postgres.Host, MediaConfig.Postgres.Port, MediaConfig.Postgres.User, MediaConfig.Postgres.Password, MediaConfig.Postgres.DB)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// video_id DESC, title ASC, length_seconds ASC, screen_fid DESC, screen_youtube_url ASC, key_worlds ASC, info ASC, who ASC, created DESC
	strQuery := `CREATE TABLE IF NOT EXISTS crawler_tube (
	id serial 		PRIMARY KEY NOT NULL,
	video_id 		varchar(11) default ''::character  varying  NOT NULL,
	tube_fid 		varchar  default ''::character  varying  NOT NULL,
	title 			varchar default ''::varchar 		NOT NULL,
	length_seconds 	varchar default ''::varchar 		NOT NULL,
	screen_fid 		varchar default ''::varchar		NOT NULL,
	screen_youtube_url 	varchar default ''::varchar	NOT NULL,
	key_worlds 			varchar default ''::varchar 				NOT NULL,
	info 				varchar default '{}' 						NOT NULL,
	who 				varchar default ''::varchar				NOT NULL,
	created 			timestamp with time zone 		NOT NULL,
 	new_tube_id  varchar default ''::varchar 		NOT NULL,
 	video_type  varchar default 'COPY'::varchar 		NOT NULL,
 	copy_state varchar default 'DOWNLOADING'::varchar 		NOT NULL);`

	_, err = db.Exec(strQuery)
	if err != nil {
		fmt.Println("db.Exec  ", err)
		return nil, err
	}

	glog.V(5).Infoln("Successfully connected!")
	return db, nil
}
func LoadConf(filepath string) (*models.Config, error) {
	if filepath == "" {
		return nil, errors.New("filepath is empty, must use --config xxx.yml/json")
	}

	data, err := iot.ReadFile(filepath)
	if err != nil {
		data = configStr
		glog.Infoln("debug mode,yaml is not use")
	}

	var cfg models.Config
	if strings.HasSuffix(filepath, ".json") {
		err = json.Unmarshal(data, &cfg)
	} else if strings.HasSuffix(filepath, ".yml") || strings.HasSuffix(filepath, ".yaml") {
		err = yaml.Unmarshal(data, &cfg)
	} else {
		return nil, errors.New("you config file must be json/yml")
	}

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
