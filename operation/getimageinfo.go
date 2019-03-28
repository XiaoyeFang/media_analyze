package operation

import (
	"bytes"
	"fmt"
	"github.com/EdlinOrg/prominentcolor"
	"github.com/golang/glog"
	"image"
	"pure-media/models"
)

func GetImageInfo(bysilce []byte) (imageInfo *models.UploadFileInfo, err error) {
	imageInfo = &models.UploadFileInfo{}
	//获取图片信息
	img, format, err := image.Decode(bytes.NewBuffer(bysilce))
	if err != nil {

		return imageInfo, err
	}
	item, err := prominentcolor.KmeansWithAll(1, img, 0, 0, nil)
	if err != nil {
		return imageInfo, err
	}
	imageInfo.ImageFormat = format
	imageInfo.ImageWidth = int64(img.Bounds().Dx())
	imageInfo.ImageHeight = int64(img.Bounds().Dy())

	//R, G, B uint32
	if len(item) != 0 {
		r := item[0].Color.R
		g := item[0].Color.G
		b := item[0].Color.B
		imageInfo.ColorExtractorRGB = fmt.Sprintf("RGB(%d,%d,%d)", r, g, b)
		//强转有时会有偏差
		imageInfo.ColorExtractorHEX = fmt.Sprintf("#%s", t2x(r)+t2x(g)+t2x(b))
	}
	glog.V(2).Infof("imageInfo = %v \n", imageInfo)

	return imageInfo, err

}
