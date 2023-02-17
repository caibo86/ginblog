package v1

import (
	"github.com/caibo86/ginblog/api/base"
	"github.com/caibo86/ginblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	category := &model.Category{}
	err := c.ShouldBindJSON(category)
	if err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	//exist, err := model.CheckCategoryExist(category.Name)
	//if err != nil {
	//	base.RenderError(c, http.StatusInternalServerError, err)
	//	return
	//}
	//if exist {
	//	base.RenderResult(c, errmsg.ErrNameUsed, nil)
	//	return
	//}
	err = model.CreateCategory(category)
	base.RenderResult(c, err, gin.H{
		"category": category,
	})
}

// IndexCategory 查询分类
func IndexCategory(c *gin.Context) {
	pageSize, page := base.GetPaginate(c)
	categories, total, err := model.IndexCategory(pageSize, page)
	base.RenderResult(c, err, gin.H{
		"categories": categories,
		"total":      total,
	})
}

// ShowCategory 获取单个分类信息
func ShowCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	category, err := model.ShowCategory(id)
	base.RenderResult(c, err, gin.H{
		"category": category,
	})
}

// UpdateCategory 更新分类
func UpdateCategory(c *gin.Context) {
	category := &model.Category{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	if err = c.ShouldBindJSON(category); err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	//exist, err := model.CheckCategoryExist(category.Name)
	//if err != nil {
	//	base.RenderError(c, http.StatusInternalServerError, err)
	//	return
	//}
	//if exist {
	//	base.RenderResult(c, errmsg.ErrNameUsed, nil)
	//	return
	//}
	err = model.UpdateCategory(id, category)
	base.RenderResult(c, err, gin.H{
		"category": category,
	})
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	err = model.DeleteCategory(id)
	base.RenderResult(c, err, nil)
}
