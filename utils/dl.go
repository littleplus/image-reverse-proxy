package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	urlTool "net/url"
	"os"
)

func DownloadImage(url string) (fileName string, err error) {
	log.Infof("start to download image(%v)", url)
	client := resty.New()
	u, err := urlTool.Parse(url)
	if err != nil {
		return
	}

	req, err := client.R().SetHeader("Referrer", u.Host).Get(url)
	if err != nil {
		return
	}

	rawImage := req.Body()
	//newImage, err := ImageToJPG(rawImage)
	log.Infof("compute image(%v) md5", url)
	fileName = fmt.Sprintf("%x", md5.Sum(rawImage)) + "." +
		GetImageExtFromHeader(req.RawResponse.Header.Get("Content-Type"))
	if err != nil {
		return
	}

	log.Infof("write image(%v) file", fileName)
	err = ioutil.WriteFile(ImageDownloadDir+"/"+fileName, rawImage, os.ModePerm)
	log.Infof("finish download image(%v)", fileName)
	return
}
