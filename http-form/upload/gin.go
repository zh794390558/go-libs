package main

import (
	"github.com/gin-gonic/gin"
)

var (
	uploadFileKey = "upload-key"
)

func main() {
	r := gin.Default()
	r.POST("/upload", uploadHandler)
	r.Run()
}

func uploadHandler(c *gin.Context) {
	header, err := c.FormFile(uploadFileKey)
	if err != nil {
		//ignore
	}
	dst := header.Filename
	// gin 简单做了封装,拷贝了文件流
	if err := c.SaveUploadedFile(header, dst); err != nil {
		// ignore
	}
}
