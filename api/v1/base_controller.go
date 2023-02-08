package v1

import (
	"github.com/caibo86/ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RenderResult(c *gin.Context, code error, data interface{}) {
	if code != errmsg.Success {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": code.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": code.Error(),
			"data":    data,
		})
	}
}

func RenderError(c *gin.Context, status int, err error) {
	var code errmsg.Code
	switch status {
	case http.StatusBadRequest:
		code = errmsg.ErrBadRequest
	default:
		code = errmsg.ErrServer
	}
	c.JSON(status, gin.H{
		"status":  code,
		"message": code.With(err),
	})
}

func GetPaginate(c *gin.Context) (int, int) {
	perPage, _ := strconv.Atoi(c.Query("per_page"))
	page, _ := strconv.Atoi(c.Query("page"))
	if perPage < 1 {
		perPage = 20
	}
	if page < 1 {
		page = 1
	}
	return perPage, page
}
