package router

import (
	"os"

	"acgfate/middleware"
	"github.com/gin-gonic/gin"
)

// 初始化路由
func NewRouter() *gin.Engine {
	var router = gin.Default()

	// 中间件，顺序固定
	router.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	router.Use(middleware.Cors())
	router.Use(middleware.CurrentUser())

	// 路由组
	r := router.Group("/api/v1")

	InitBaseRouter(r)    // 注册基础功能路由，不做鉴权
	InitUserRouter(r)    // 注册用户路由
	InitArticleRouter(r) // 注册文章路由

	return router
}
