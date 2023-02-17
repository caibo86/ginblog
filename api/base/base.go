package base

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// OffsetByPage 根据页计算偏移
func OffsetByPage(pageSize, page int) int {
	if pageSize == -1 {
		return -1
	}
	return (page - 1) * pageSize
}

func GetPaginate(c *gin.Context) (int, int) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	page, _ := strconv.Atoi(c.Query("page"))
	if pageSize == 0 {
		pageSize = 20
	} else if pageSize < 0 {
		pageSize = -1
	}
	if page < 1 {
		page = 1
	}
	return pageSize, page
}
