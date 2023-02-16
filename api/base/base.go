package base

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// OffsetByPage 根据页计算偏移
func OffsetByPage(pageSize, page int) int {
	return (page - 1) * pageSize
}

func GetPaginate(c *gin.Context) (int, int) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	page, _ := strconv.Atoi(c.Query("page"))
	if pageSize < 1 {
		pageSize = 20
	}
	if page < 1 {
		page = 1
	}
	return pageSize, page
}
