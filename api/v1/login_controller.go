package v1

import (
	"github.com/caibo86/ginblog/api/base"
	"github.com/caibo86/ginblog/model"
	"github.com/caibo86/ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	user := &model.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	if err := model.CheckLogin(user.Username, user.Password); err != nil {
		base.RenderResult(c, err, nil)
		return
	}
	token, err := utils.SetToken(user.Username)
	base.RenderResult(c, err, gin.H{
		"token": token,
	})
}
