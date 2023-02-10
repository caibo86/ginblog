package base

import (
	"github.com/caibo86/ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RenderResult(c *gin.Context, err error, data gin.H) {
	if err == nil {
		h := gin.H{
			"status":  0,
			"message": "OK",
		}
		if data != nil {
			h["data"] = data
		}
		c.JSON(http.StatusOK, h)
	} else if code, ok := err.(errmsg.Code); ok {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": code.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ErrServer,
			"message": errmsg.ErrServer.With(err),
		})
	}
}

func RenderError(c *gin.Context, status int, err error) {
	var code errmsg.Code
	var ok bool
	if code, ok = err.(errmsg.Code); !ok {
		switch status {
		case http.StatusBadRequest:
			code = errmsg.ErrBadRequest
		case http.StatusUnauthorized:
			code = errmsg.ErrUnauthorized
		case http.StatusInternalServerError:
			code = errmsg.ErrServer
		default:
			code = errmsg.ErrServer
		}
	} else {
		err = nil
	}
	c.JSON(status, gin.H{
		"status":  code,
		"message": code.With(err),
	})
}
