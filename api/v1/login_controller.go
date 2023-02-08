package v1

import (
	"github.com/caibo86/ginblog/model"
	"github.com/caibo86/ginblog/utils"
	"github.com/caibo86/ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	user := &model.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	code := model.CheckLogin(user.Username, user.Password)
	if code != errmsg.OK {
		RenderResult(c, code, nil)
		return
	}
	token, code := utils.SetToken(user.Username)
	RenderResult(c, code, token)
}
