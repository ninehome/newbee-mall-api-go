package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/utils"
	"net/http"
)

func Upload(c *gin.Context) {
	file, fileheader, _ := c.Request.FormFile("file")
	filesize := fileheader.Size
	url, code := utils.UploadToQiNiu(file, filesize)

	c.JSONP(http.StatusOK, gin.H{
		"status":  code,
		"message": "",
		"url":     url,
	})
}
