package v1

import (
	"github.com/caibo86/ginblog/api/base"
	"github.com/caibo86/ginblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateArticle 创建文章
func CreateArticle(c *gin.Context) {
	article := &model.Article{}
	if err := c.ShouldBindJSON(article); err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	err := model.CreateArticle(article)
	base.RenderResult(c, err, gin.H{
		"article": article,
	})
}

// IndexArticle 查询文章
func IndexArticle(c *gin.Context) {
	var categoryID int
	var err error
	pageSize, page := base.GetPaginate(c)
	if c.Query("category_id") != "" {
		categoryID, err = strconv.Atoi(c.Query("category_id"))
		if err != nil {
			base.RenderError(c, http.StatusBadRequest, err)
			return
		}
	}
	title := c.Query("title")
	articles, total, err := model.IndexArticle(pageSize, page, categoryID, title)
	base.RenderResult(c, err, gin.H{
		"articles": articles,
		"total":    total,
	})
}

// ShowArticle 查询单个文章
func ShowArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	article, err := model.ShowArticle(id)
	base.RenderResult(c, err, gin.H{
		"article": article,
	})
}

// UpdateArticle 更新文章
func UpdateArticle(c *gin.Context) {
	article := &model.Article{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	if err = c.ShouldBindJSON(article); err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	err = model.UpdateArticle(id, article)
	base.RenderResult(c, err, gin.H{
		"article": article,
	})
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		base.RenderError(c, http.StatusBadRequest, err)
		return
	}
	code := model.DeleteArticle(id)
	base.RenderResult(c, code, nil)
}
