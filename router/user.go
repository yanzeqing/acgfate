package router

import (
	v1 "acgfate/api/v1"
	"acgfate/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("user").Use(middleware.AuthRequired())
	{
		UserRouter.GET("me", v1.UserMe)
		UserRouter.DELETE("logout", v1.UserLogout)
	}
}
