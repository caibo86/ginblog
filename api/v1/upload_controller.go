package v1

import (
	"github.com/caibo86/ginblog/api/base"
	"github.com/caibo86/ginblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	fileSize := fileHeader.Size
	url, err := model.UploadFile(file, fileSize)
	if err != nil {
		base.RenderError(c, http.StatusInternalServerError, err)
		return
	}
	base.RenderResult(c, nil, gin.H{
		"url": url,
	})
}
