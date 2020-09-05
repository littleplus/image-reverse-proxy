package utils

import (
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	urlTool "net/url"
	"os"
)

func DownloadImage(url string, fileName string) (err error) {
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
	newImage, err := ImageToJPG(rawImage)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(ImageDownloadDir+"/"+fileName, newImage, os.ModePerm)
	return
}
