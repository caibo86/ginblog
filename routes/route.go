package routes

import (
	v1 "github.com/caibo86/ginblog/api/v1"
	"github.com/caibo86/ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		// 用户模块的路由接口
		router.POST("users", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("users/:id", v1.EditUser)
		router.DELETE("users/:id", v1.DeleteUser)
		// 文章模块的路由接口
		// 分类模块的路由接口
	}
	r.Run(utils.HttpPort)

}
