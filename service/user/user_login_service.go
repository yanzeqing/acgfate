package user

import (
	"acgfate/model"
	"acgfate/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginService 管理用户登录的服务
type LoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=2,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (service *LoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("uid", user.Uid)
	_ = s.Save()
}

// Login 用户登录函数
func (service *LoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if user.Status == model.Suspend {
		return serializer.ParamErr("账号被封禁", nil)
	}
	if user.Status == model.Inactive {
		return serializer.ParamErr("账号未激活", nil)
	}

	// 设置session
	service.setSession(c, user)

	return serializer.BuildUserResponse(user)
}
