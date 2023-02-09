package v1

import (
	"github.com/caibo86/ginblog/model"
	"github.com/caibo86/ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	fileSize := fileHeader.Size
	url, err := model.UploadFile(file, fileSize)
	if err != nil {
		RenderError(c, http.StatusInternalServerError, err)
		return
	}
	RenderResult(c, errmsg.OK, url)
}
