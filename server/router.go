package server

import (
	"acgfate/api"
	"acgfate/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
			// 文章投稿
			auth.POST("articles", api.CreateArticle)
			// 文章详情
			auth.GET("article/:id", api.DetailArticle)
			// 文章列表
			auth.GET("articles", api.ListArticle)
			// 文章更新
			auth.PUT("article/:id", api.UpdateArticle)
			// 文章删除
			auth.DELETE("article/:id", api.DeleteArticle)
		}
	}
	return r
}
