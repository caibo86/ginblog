package v1

import (
	"github.com/caibo86/ginblog/utils"
	"github.com/caibo86/ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func RenderResult(c *gin.Context, code error, data interface{}) {
	if code != errmsg.OK {
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
	case http.StatusUnauthorized:
		code = errmsg.ErrUnauthorized
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

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		a := c.Request.Header.Get("Authorization")
		if a == "" {
			RenderError(c, http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		ss := strings.SplitN(a, " ", 2)
		if len(ss) != 2 && ss[0] != "Bearer" {
			RenderError(c, http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		claims, code := utils.CheckToken(ss[1])
		if code != errmsg.OK {
			RenderError(c, http.StatusUnauthorized, code)
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}
