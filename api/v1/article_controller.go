package v1

import (
	"github.com/caibo86/ginblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateArticle 创建文章
func CreateArticle(c *gin.Context) {
	article := &model.Article{}
	if err := c.ShouldBindJSON(article); err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	code := model.CreateArticle(article)
	RenderResult(c, code, article)
}

// IndexArticle 查询文章
func IndexArticle(c *gin.Context) {
	var categoryID int
	var err error
	perPage, page := GetPaginate(c)
	if c.Query("category_id") != "" {
		categoryID, err = strconv.Atoi(c.Query("category_id"))
		if err != nil {
			RenderError(c, http.StatusBadRequest, err)
			return
		}
	}
	articles, total, code := model.IndexArticle(perPage, page, categoryID)
	RenderResultWithTotal(c, code, articles, total)
}

// ShowArticle 查询单个文章
func ShowArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	article, code := model.ShowArticle(id)
	RenderResult(c, code, article)
}

// UpdateArticle 更新文章
func UpdateArticle(c *gin.Context) {
	article := &model.Article{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	if err = c.ShouldBindJSON(article); err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	code := model.UpdateArticle(id, article)
	RenderResult(c, code, article)
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	code := model.DeleteArticle(id)
	RenderResult(c, code, nil)
}
