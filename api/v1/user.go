package v1

import (
	"acgfate/serializer"
	"acgfate/service/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var registerService user.RegisterService
	if err := c.ShouldBind(&registerService); err == nil {
		res := registerService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var loginService user.LoginService
	if err := c.ShouldBind(&loginService); err == nil {
		res := loginService.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	currentUser := CurrentUser(c)
	res := serializer.BuildUserResponse(*currentUser)
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	_ = s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
