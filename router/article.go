package router

import (
	v1 "acgfate/api/v1"
	"acgfate/middleware"
	"github.com/gin-gonic/gin"
)

func InitArticleRouter(router *gin.RouterGroup) {
	ArticleRouter := router.Group("article").Use(middleware.AuthRequired())
	{
		ArticleRouter.POST("upload", v1.CreateArticle)
		ArticleRouter.PUT(":aid", v1.UpdateArticle)
		ArticleRouter.DELETE(":aid", v1.DeleteArticle)
	}
}
