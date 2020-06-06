package router

import (
	v1 "acgfate/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(router *gin.RouterGroup) {
	BaseRouter := router.Group("")
	{
		BaseRouter.POST("user/register", v1.UserRegister)             // 用户注册
		BaseRouter.POST("user/login", v1.UserLogin)                   // 用户登录
		BaseRouter.GET("articles/category/:category", v1.ListArticle) // 文章分区列表
		BaseRouter.GET("articles/author/:author", v1.AuthorArticle)   // 文章作者列表
		BaseRouter.GET("article/:aid", v1.DetailArticle)              // 文章详情
	}
}
