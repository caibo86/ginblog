package v1

import (
	"github.com/caibo86/ginblog/api/base"
	"github.com/caibo86/ginblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ShowProfile 查询个人配置信息
func ShowProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	profile, err := model.ShowProfile(id)
	base.RenderResult(c, err, gin.H{
		"profile": profile,
	})
}

// UpdateProfile 更新用户配置信息
func UpdateProfile(c *gin.Context) {
	profile := &model.Profile{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	if err = c.ShouldBindJSON(profile); err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	err = model.UpdateProfile(id, profile)
	base.RenderResult(c, err, gin.H{
		"profile": profile,
	})
}
