package base

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// OffsetByPage 根据页计算偏移
func OffsetByPage(perPage, page int) int {
	return (page - 1) * perPage
}

func GetPaginate(c *gin.Context) (int, int) {
	perPage, _ := strconv.Atoi(c.Query("per_page"))
	page, _ := strconv.Atoi(c.Query("page"))
	if perPage < 1 {
		perPage = 20
	}
	if page < 1 {
		page = 1
	}
	return perPage, page
}
