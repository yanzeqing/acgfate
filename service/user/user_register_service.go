package user

import (
	"acgfate/model"
	"acgfate/serializer"
)

// RegisterService 管理用户注册服务
type RegisterService struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required,ascii,min=3,max=10"`
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=1,max=10"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=16"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,eqfield=Password,min=8,max=16"`
}

// valid 验证表单
func (service *RegisterService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已被注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *RegisterService) Register() serializer.Response {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   model.Active,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	return serializer.BuildUserResponse(user)
}
