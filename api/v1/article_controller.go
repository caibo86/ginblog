package v1

import (
	"github.com/caibo86/ginblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateArticle(c *gin.Context) {
	article := &model.Article{}
	if err := c.ShouldBindJSON(article); err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	code := model.CreateArticle(article)
	RenderResult(c, code, article)
}

func IndexArticle(c *gin.Context) {
	perPage, page := GetPaginate(c)
	articles, code := model.IndexArticle(perPage, page)
	RenderResult(c, code, articles)
}

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

func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		RenderError(c, http.StatusBadRequest, err)
		return
	}
	code := model.DeleteArticle(id)
	RenderResult(c, code, nil)
}
