package v1

import (
	"github.com/caibo86/ginblog/api/base"
	"github.com/caibo86/ginblog/model"
	"github.com/caibo86/ginblog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateUser 添加用户
func CreateUser(c *gin.Context) {
	user := &model.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	err := validator.Validate(user)
	if err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	// 利用DB唯一键 暂时不用这段代码
	//exist, err := model.CheckUserExist(user.Username)
	//if err != nil {
	//	base.RenderError(c, http.StatusInternalServerError, err)
	//	return
	//}
	//if exist {
	//	base.RenderResult(c, errmsg.ErrNameUsed, nil)
	//	return
	//}
	err = model.CreateUser(user)
	base.RenderResult(c, err, gin.H{
		"user": user,
	})
}

// IndexUser 查询用户列表
func IndexUser(c *gin.Context) {
	perPage, page := base.GetPaginate(c)
	users, total, err := model.IndexUser(perPage, page)
	base.RenderResult(c, err, gin.H{
		"users": users,
		"total": total,
	})
}

// UpdateUser 编辑用户
func UpdateUser(c *gin.Context) {
	user := &model.User{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	if err = c.ShouldBindJSON(user); err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	//exist, err := model.CheckUserExist(user.Username)
	//if err != nil {
	//	base.RenderError(c, http.StatusInternalServerError, err)
	//	return
	//}
	//if exist {
	//	base.RenderResult(c, errmsg.ErrNameUsed, nil)
	//	return
	//}
	err = model.UpdateUser(id, user)
	base.RenderResult(c, err, gin.H{
		"user": user,
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	err = model.DeleteUser(id)
	base.RenderResult(c, err, nil)
}
