package routes

import (
	"github.com/caibo86/ginblog/api/middleware"
	v1 "github.com/caibo86/ginblog/api/v1"
	"github.com/caibo86/ginblog/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	//r.Use(gin.Logger())
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	router := r.Group("api/v1")
	router.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		router.POST("users", v1.CreateUser)
		router.GET("users", v1.IndexUser)
		router.GET("users/:id", v1.ShowUser)
		router.PUT("users/:id", v1.UpdateUser)
		router.DELETE("users/:id", v1.DeleteUser)
		// 文章模块的路由接口
		router.POST("articles", v1.CreateArticle)
		router.GET("articles", v1.IndexArticle)
		router.GET("articles/:id", v1.ShowArticle)
		router.PUT("articles/:id", v1.UpdateArticle)
		router.DELETE("articles/:id", v1.DeleteArticle)
		// 分类模块的路由接口
		router.POST("categories", v1.CreateCategory)
		router.GET("categories", v1.IndexCategory)
		router.GET("categories/:id", v1.ShowCategory)
		router.PUT("categories/:id", v1.UpdateCategory)
		router.DELETE("categories/:id", v1.DeleteCategory)
		// 文件上传
		router.POST("upload", v1.Upload)
	}
	public := r.Group("api/v1")
	{
		// 用户模块的路由接口

		public.POST("login", v1.Login)
	}
	if err := r.Run(utils.HttpPort); err != nil {
		log.Fatal(err)
	}

}
