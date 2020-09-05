package main

import (
	"git.runode.com/littleplus/image-reverse-proxy/service"
	"git.runode.com/littleplus/image-reverse-proxy/utils"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	_ = os.MkdirAll(utils.ImageDownloadDir, os.ModeDir)
}

func main() {
	r := gin.Default()
	r.GET("/s", service.Search)
	r.Static(utils.ImageRequestPath, utils.ImageDownloadDir)
	err := r.Run()
	if err != nil {
		panic(err.Error())
	}
}
