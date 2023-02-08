package v1

import (
	"github.com/caibo86/ginblog/model"
	"github.com/caibo86/ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateUser 添加用户
func CreateUser(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		code := errmsg.ErrBadRequest
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"message": code.With(err),
		})
		return
	}
	code := model.CheckUser(data.Username)
	if code == errmsg.Success {
		code = model.CreateUser(&data)
	}
	RenderResult(c, code, data)
}

// IndexUser 查询用户列表
func IndexUser(c *gin.Context) {
	perPage, page := GetPaginate(c)
	var users []*model.User
	code := model.IndexUser(perPage, page, users)
	RenderResult(c, code, users)
}

// UpdateUser 编辑用户
func UpdateUser(c *gin.Context) {
	u := &model.User{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	if err = c.ShouldBindJSON(u); err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	code := model.CheckUser(u.Username)
	if code == errmsg.Success {
		code = model.UpdateUser(id, u)
	}
	RenderResult(c, code, u)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	code := model.DeleteUser(id)
	RenderResult(c, code, nil)
}
