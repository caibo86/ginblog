package v1

import (
	"github.com/caibo86/ginblog/model"
	"github.com/caibo86/ginblog/utils/errmsg"
	"github.com/caibo86/ginblog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateUser 添加用户
func CreateUser(c *gin.Context) {
	user := &model.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	msg, err := validator.Validate(user)
	if err != nil {
		RenderError(c, http.StatusBadRequest, err, msg)
		return
	}
	code := model.CheckUser(user.Username)
	if code == errmsg.OK {
		code = model.CreateUser(user)
	}
	RenderResult(c, code, user)
}

// IndexUser 查询用户列表
func IndexUser(c *gin.Context) {
	perPage, page := GetPaginate(c)
	users, total, code := model.IndexUser(perPage, page)
	RenderResultWithTotal(c, code, users, total)
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
	if code == errmsg.OK {
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
