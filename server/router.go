package server

import (
	"os"

	"acgfate/api/v1"
	"acgfate/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序固定
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	r1 := r.Group("/api/v1")
	{
		r1.POST("ping", v1.Ping)

		// 用户注册
		r1.POST("user/register", v1.UserRegister)

		// 用户登录
		r1.POST("user/login", v1.UserLogin)
		// 需要登录保护的
		auth := r1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// 用户路由
			auth.GET("user/me", v1.UserMe)
			auth.DELETE("user/logout", v1.UserLogout)
			// 文章操作
			auth.POST("articles", v1.CreateArticle)
			auth.PUT("article/:aid", v1.UpdateArticle)
			auth.DELETE("article/:aid", v1.DeleteArticle)
		}
		// 文章列表
		r1.GET("articles/:category", v1.ListArticle)
		// 文章详情
		r1.GET("article/:aid", v1.DetailArticle)
	}
	return r
}
