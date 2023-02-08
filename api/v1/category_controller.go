package v1

import (
	"github.com/caibo86/ginblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	category := &model.Category{}
	if err := c.ShouldBindJSON(category); err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	code := model.CreateCategory(category)
	RenderResult(c, code, category)
}

// IndexCategory 查询分类
func IndexCategory(c *gin.Context) {
	perPage, page := GetPaginate(c)
	code, categories := model.IndexCategory(perPage, page)
	RenderResult(c, code, categories)
}

// UpdateCategory 更新分类
func UpdateCategory(c *gin.Context) {
	category := &model.Category{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	if err = c.ShouldBindJSON(category); err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	code := model.UpdateCategory(id, category)
	RenderResult(c, code, category)
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	code := model.DeleteCategory(id)
	RenderResult(c, code, nil)
}
