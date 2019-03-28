package models

import (
	"image/color"
	"strconv"
	"sync/atomic"
	"time"
)

// Config contains the configuration of the url shortener.
type Config struct {
	GrpcListen      string   `yaml:"grpc_listen" json:"grpc_listen"`
	ProxyHttp       []string `yaml:"proxy_http" json:"proxy_http"`
	ProxyHttpPrefix string   `yaml:"proxy_http_prefix" json:"proxy_http_prefix"`
	LogLevel        string   `yaml:"log_level" json:"log_level"`
	Postgres        struct {
		Host     string `yaml:"host" json:"host"`
		Port     string `yaml:"port" json:"port"`
		User     string `yaml:"user" json:"user"`
		Password string `yaml:"password" json:"password"`
		DB       string `yaml:"db" json:"db"`
	} `yaml:"postgres" json:"postgres"`
	AccessKeyId     string   `yaml:"access_key_id" json:"access_key_id"`
	AccessHeySecret string   `yaml:"access_key_secret" json:"access_key_secret"`
	YtApiKey        string   `yaml:"yt_api_key" json:"yt_api_key"`
	YtApiPart       string   `yaml:"yt_api_part" json:"yt_api_part"`
	YtApiFields     string   `yaml:"yt_api_fields" json:"yt_api_fields"`
	BurstLimit      int      `yaml:"burst_limit" json:"burst_limit"`
	S3ConfImage     S3Config `yaml:"s3_conf_image" json:"s3_conf_image"`
	S3ConfYoutube   S3Config `yaml:"s3_conf_youtube" json:"s3_conf_youtube"`
	QiniuApi        QiniuApi `yaml:"qiniu_api" json:"qiniu_api"`
}

type S3Config struct {
	Endpoint  string `yaml:"endpoint" json:"endpoint"`
	AccessKey string `yaml:"access_key" json:"access_key"`
	SecretKey string `yaml:"secret_key" json:"secret_key"`
	Bucket    string `yaml:"bucket" json:"bucket"`
	Prefix    string `yaml:"prefix" json:"prefix"`
	S3Acl     string `yaml:"s3acl" json:"s3acl"`
	ChunkSize string `yaml:"chunk_size" json:"chunk_size"`
}

type QiniuApi struct {
	QiniuAkey string `yaml:"qiniu_akey" json:"qiniu_akey"`
	QiniuSkey string `yaml:"qiniu_skey" json:"qiniu_skey"`
	QiniuHost string `yaml:"qiniu_host" json:"qiniu_host"`
	Bucket    string `yaml:"bucket" json:"bucket"`
	Expires   uint32 `yaml:"expires" json:"expires"`
}

type UploadFileInfo struct {
	FileId            string   //S3文件ID
	FileSize          int64    //文件大小，单位：Bytes
	FileSha256        string   //文件哈希
	FileType          string   //文件类型image
	ImageFormat       string   //图片类型
	ImageWidth        int64    //图片宽度
	ImageHeight       int64    //图片高度
	ColorExtractorRGB string   //图片颜色主色调RGB
	ColorExtractorHEX string   //图片颜色主色调HEX
	Tags              []string //标签方便搜索
	//Context           []byte    //上下文，JSON类型
	CreatedAt string //创建时间
}

type ImgMsg struct {
	/*
		size，文件大小，单位：Bytes
		format，图片类型，如png、jpeg、gif、bmp等。
		width，图片宽度，单位：像素(px)。
		height，图片高度，单位：像素(px)。
		colorExtractor，图片颜色主色调
		colorModel，彩色空间，如palette16、ycbcr等。
		frameNumber，帧数，gif 图片会返回此项。
		exif，图片EXIF信息（DateTime、ExposureBiasValue、ExposureTime、Model、ISOSpeedRatings、
		ResolutionUnit等） 参考技术白皮书 http://www.cipa.jp/std/documents/e/DC-008-2012_E.pdf
	*/

	Size           int64       `yaml:"size" json:"size"`
	Format         string      `yaml:"format" json:"format"`
	Width          int64       `yaml:"width" json:"width"`
	Height         int64       `yaml:"height" json:"height"`
	ColorExtractor string      `yaml:"colorExtractor" json:"colorExtractor"`
	ColorModel     color.Model `yaml:"colorModel" json:"colorModel"`
	FrameNumber    int         `yaml:"frameNumber" json:"frameNumber"`
	Exif           Exif        `yaml:"exif" json:"exif"`
}

type Exif struct {
	DateTime          string `yaml:"dateTime" json:"dateTime"`
	ExposureBiasValue string `yaml:"exposureBiasValue" json:"exposureBiasValue"`
	ExposureTime      string `yaml:"exposureTime" json:"exposureTime"`
	Model             string `yaml:"model" json:"model"`
	ISOSpeedRatings   string `yaml:"iSOSpeedRatings" json:"iSOSpeedRatings"`
	ResolutionUnit    string `yaml:"resolutionUnit" json:"resolutionUnit"`
}

type VideoMsg struct {
	//Id               int64            `orm:"pk;column(id);auto"`
	//	VideoId          string           `orm:"cloumn(video_id);size(11)" json:"video_id"`
	//	Title            string           `orm:"cloumn(title);" json:"title"`
	//	LengthSeconds    string           `orm:"cloumn(length_seconds)" json:"length_seconds"`
	//	ScreenFid        string           `orm:"cloumn(screen_fid)" json:"screen_fid"`
	//	ScreenYoutubeUrl string           `orm:"cloumn(screen_youtube_url)" json:"screen_youtube_url"`
	//	KeyWorlds        SliceStringField `orm:"cloumn(key_worlds)" json:"key_worlds"`
	//	Info             string           `orm:"type(jsonb)" json:"info"`
	//	Who              SliceStringField `orm:"cloumn(who)" json:"who"`
	//	Created          time.Time
	Id               int64    `yaml:"id" json:"id"`
	VideoId          string   `yaml:"video_id" json:"video_id"`             //视频ID
	Title            string   `yaml:"title" json:"title"`                   //视频标题
	LengthSeconds    string   `yaml:"length_seconds" json:"length_seconds"` //视频播放长度（秒）
	ScreenFid        string   `yaml:"screen_fid" json:"screen_fid"`
	ScreenYoutubeUrl string   `yaml:"screen_youtube_url" json:"screen_youtube_url"` //视频截图
	KeyWords         []string `yaml:"key_worlds" json:"key_worlds"`                 //视频关键词
	Info             string   `yaml:"info" json:"info"`                             //
	Who              string   `yaml:"who" json:"who"`
	//HL               string   `yaml:"hl" json:"hl"`                 //视频发布语言
	//Author           string   `yaml:"author" json:"author"`         //视频作者
	//Rating           string   `yaml:"rating" json:"rating"`         //评分
	//ViewCount        string   `yaml:"viewCount" json:"viewCount"`   //播放量
	//Definition       string   `yaml:"definition" json:"definition"` //视频清晰度
	Created time.Time `yaml:"created" json:"created"`

	TubeFid   string `yaml:"tube_fid" json:"tube_fid"`
	NewTubeId string `yaml:"new_tube_id" json:"new_tube_id"`
	Type      string `yaml:"type" json:"type"`
	CopyState string `yaml:"copy_state" json:"copy_state"`
}

// An AtomicInt is an int64 to be accessed atomically.
type AtomicInt int64

// Add atomically adds n to i.
func (i *AtomicInt) Add(n int64) {
	atomic.AddInt64((*int64)(i), n)
}

// Add atomically adds n to i.
func (i *AtomicInt) Set(n int64) {
	atomic.StoreInt64((*int64)(i), n)
}

// Get atomically gets the value of i.
func (i *AtomicInt) Get() int64 {
	return atomic.LoadInt64((*int64)(i))
}

func (i *AtomicInt) String() string {
	return strconv.FormatInt(i.Get(), 10)
}

type YouTubeVideoInfo struct {
	kind             string `json:"kind"`
	Id               int64  `orm:"pk;column(id);auto"`
	VideoId          string `orm:"cloumn(video_id);size(11)" json:"video_id"`
	Title            string `orm:"cloumn(title);" json:"title"`
	LengthSeconds    string `orm:"cloumn(length_seconds)" json:"length_seconds"`
	ScreenFid        string `orm:"cloumn(screen_fid)" json:"screen_fid"`
	ScreenYoutubeUrl string `orm:"cloumn(screen_youtube_url)" json:"screen_youtube_url"`
	//KeyWorlds        SliceStringField `orm:"cloumn(key_worlds)" json:"key_worlds"`
	Info string `orm:"type(jsonb)" json:"info"`
	//Who              SliceStringField `orm:"cloumn(who)" json:"who"`
	Created time.Time `orm:"auto_now_add;type(date)`
}

/*
{
         "kind": "youtube#videoListResponse",
         "etag": "\"XI7nbFXulYBIpL0ayR_gDh3eu1k/QVkH6oXGvMJMAoUUax8M4veB7qg\"",
         "pageInfo": {
          "totalResults": 1,
          "resultsPerPage": 1
         },
         "items": [
          {
           "kind": "youtube#video",
           "etag": "\"XI7nbFXulYBIpL0ayR_gDh3eu1k/r3rFnj_2cuJmFGsReaqSYHP_UZA\"",
           "id": "2EuTs1Yo-Bo",
           "snippet": {
            "publishedAt": "2014-11-18T21:24:06.000Z",
            "channelId": "UC6LHN4GnT2ui5wRBQanNKxQ",
            "title": "【刺客教條：大革命】- PC特效全開中文劇情電影60FPS - 第七集 -  Episode 7 - Assassin's Creed：Unity - 刺客信條 ： 大革命 - 最強無損畫質影片",
            "description": "全集播放清單：https://www.youtube.com/playlist?list=PLFbpBUhLEeCyzI-RzcRcCvzJvZNka6hyf\n全程特效全開 1080P 60FPS 錄製，給您最佳的電影級享受！\n【刺客教條：大革命】- PC特效全開中文劇情電影60FPS - 第七集\nPC主機配備 :\nDriver 344.65\n-i7-3770k OC 4.5Ghz\n-Sandisk 1600mhz 16G RAM (8GX2)\n-EVGA GTX780  ACX SC3GD5\n-Sandisk Extreme SSD 120G\n-SuperFlower 850W silver 80+\n\n遊戲簡介：\n《刺客教條：大革命》遊戲的背景設定在18世紀的巴黎，法國大革命時期。藉由新世代繪圖硬體呈現更細緻寫實的場景與角色。\n本作主角名“亞諾”，來自日爾曼語裡的鷹，設定為18世紀的法國，Ezio時代的莊園系統回歸，主角將與巴黎刺客兄弟會一同進行戰鬥，支援4人合作模式。",
            "thumbnails": {
             "default": {
              "url": "https://i.ytimg.com/vi/2EuTs1Yo-Bo/default.jpg",
              "width": 120,
              "height": 90
             },
             "medium": {
              "url": "https://i.ytimg.com/vi/2EuTs1Yo-Bo/mqdefault.jpg",
              "width": 320,
              "height": 180
             },
             "high": {
              "url": "https://i.ytimg.com/vi/2EuTs1Yo-Bo/hqdefault.jpg",
              "width": 480,
              "height": 360
             },
             "standard": {
              "url": "https://i.ytimg.com/vi/2EuTs1Yo-Bo/sddefault.jpg",
              "width": 640,
              "height": 480
             },
             "maxres": {
              "url": "https://i.ytimg.com/vi/2EuTs1Yo-Bo/maxresdefault.jpg",
              "width": 1280,
              "height": 720
             }
            },
            "channelTitle": "Semenix Gaming",
            "tags": [
             "full movie",
             "pc",
             "gameplay",
             "1080",
             "ultra settings",
             "max settings",
             "攻略",
             "xbox one",
             "ps4",
             "episode 1",
             "walkthrough",
             "流程",
             "影片",
             "實況",
             "trailer",
             "xbox360",
             "ps3",
             "full gameplay",
             "part 1",
             "Semenix",
             "TGN",
             "VISO",
             "刺客教條 大革命",
             "刺客信條 大革命",
             "Assassin's Creed Unity",
             "Assassin's Creed (Video Game Series)",
             "Unity (Software)",
             "ezio",
             "亞諾",
             "anno",
             "刺客教條5",
             "刺客信條5",
             "Episode"
            ],
            "categoryId": "20",
            "liveBroadcastContent": "none",
            "localized": {
             "title": "【刺客教條：大革命】- PC特效全開中文劇情電影60FPS - 第七集 -  Episode 7 - Assassin's Creed：Unity - 刺客信條 ： 大革命 - 最強無損畫質影片",
             "description": "全集播放清單：https://www.youtube.com/playlist?list=PLFbpBUhLEeCyzI-RzcRcCvzJvZNka6hyf\n全程特效全開 1080P 60FPS 錄製，給您最佳的電影級享受！\n【刺客教條：大革命】- PC特效全開中文劇情電影60FPS - 第七集\nPC主機配備 :\nDriver 344.65\n-i7-3770k OC 4.5Ghz\n-Sandisk 1600mhz 16G RAM (8GX2)\n-EVGA GTX780  ACX SC3GD5\n-Sandisk Extreme SSD 120G\n-SuperFlower 850W silver 80+\n\n遊戲簡介：\n《刺客教條：大革命》遊戲的背景設定在18世紀的巴黎，法國大革命時期。藉由新世代繪圖硬體呈現更細緻寫實的場景與角色。\n本作主角名“亞諾”，來自日爾曼語裡的鷹，設定為18世紀的法國，Ezio時代的莊園系統回歸，主角將與巴黎刺客兄弟會一同進行戰鬥，支援4人合作模式。"
            }
           },
           "contentDetails": {
            "duration": "PT31M44S",
            "dimension": "2d",
            "definition": "hd",
            "caption": "false",
            "licensedContent": true,
            "projection": "rectangular"
           },
           "statistics": {
            "viewCount": "56371",
            "likeCount": "162",
            "dislikeCount": "5",
            "favoriteCount": "0",
            "commentCount": "49"
           }
          }
         ]
        }
*/

//youtube/v3 API返回结果
type YTApiV3Video struct {
	Page  PageInfo    `yaml:"pageInfo" json:"pageInfo"` //可以自定义cron的执行时区
	Items []*ItemInfo `yaml:"items,flow" json:"items"`
}

type ItemInfo struct {
	Kind           string             `yaml:"kind" json:"kind"`
	Id             string             `yaml:"id" json:"id"`                              // crontab
	Snippet        *SnipInfo          `yaml:"snippet,flow" json:"snippet"`               //s3 配置
	Statistics     *StatisticInfo     `yaml:"statistics,flow" json:"statistics"`         //s3 配置
	ContentDetails *ContentDetailInfo `yaml:"contentDetails,flow" json:"contentDetails"` //s3 配置
}

type SnipInfo struct {
	PublishedAt  string         `yaml:"publishedAt" json:"publishedAt"`
	ChannelId    string         `yaml:"channelId" json:"channelId"`
	Description  string         `yaml:"description" json:"description"`
	Title        string         `yaml:"title" json:"title"`
	CategoryId   string         `yaml:"categoryId" json:"categoryId"`
	ChannelTitle string         `yaml:"channelTitle" json:"channelTitle"`
	Tags         []string       `yaml:"tags" json:"tags"`
	Thumbnails   *ThumbnailInfo `json:"thumbnails"`
}

type ThumbnailInfo struct {
	Default  *Thumbnail `json:"default"`
	Medium   *Thumbnail `json:"medium"`
	High     *Thumbnail `json:"high"`
	Standard *Thumbnail `json:"standard"`
	Maxres   *Thumbnail `json:"maxres"`
}

type Thumbnail struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

//json 中的key,适配youtube返回结果进行调整
type PageInfo struct {
	TotalResult    int `json:"totalResults"`
	TesultsPerPage int `json:"resultsPerPage"`
}

//视频统计信息
type StatisticInfo struct {
	ViewCount     string `json:"viewCount"`
	LikeCount     string `json:"likeCount"`
	DislikeCount  string `json:"dislikeCount"`
	FavoriteCount string `json:"favoriteCount"`
	CommentCount  string `json:"commentCount"`
}

//视频播放信息
type ContentDetailInfo struct {
	Duration        string `json:"duration"`
	Dimension       string `json:"dimension"`
	Definition      string `json:"definition"`
	Caption         string `json:"caption"`
	LicensedContent bool   `json:"licensedContent"`
	Projection      string `json:"projection"`
}
