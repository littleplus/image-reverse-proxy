package service

import (
	"git.runode.com/littleplus/image-reverse-proxy/utils"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

func Search(c *gin.Context) {
	url := c.Query("pic")
	if url == "" {
		c.Redirect(301, "https://www.baidu.com")
	}

	e := c.Query("engine")
	if e == "" {
		e = "saucenao"
	}

	fileName := uuid.NewV4().String() + ".jpg"
	err := utils.DownloadImage(url, fileName)
	if err != nil {
		log.Errorf("ImageDownloadError: %v", err)
		c.JSON(500, "server error")
	}

	newUrl := "http://" + c.Request.Host + utils.ImageRequestPath + "/" +
		fileName
	log.Infof("Redirect: from url(%v) to url(%v)", url, newUrl)
	c.Redirect(302, getSearchEngine(e)+newUrl)
	return
}

func getSearchEngine(code string) string {
	switch code {
	case "saucenao":
		return "https://saucenao.com/search.php?db=999&url="
	case "google":
		return "https://www.google.com/searchbyimage?image_url="
	case "yandex":
		return "https://yandex.com/images/search?rpt=imageview&url="
	default:
		return "https://saucenao.com/search.php?db=999&url="
	}
}
