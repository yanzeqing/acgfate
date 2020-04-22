package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User 用户模型
type User struct {
	Uid            uint64 `gorm:"primary_key" uri:"uid"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
	UserName       string
	PasswordDigest string
	Nickname       string
	Status         string
	Avatar         string `gorm:"size:1000"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用UID获取用户
func GetUser(UID interface{}) (User, error) {
	var user User
	result := DB.First(&user, UID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

// CheckStatus 账号状态判断
func (user *User) CheckStatus() byte {
	if user.Status == Suspend {
		return 1
	}
	if user.Status == Inactive {
		return 2
	}
	return 0
}
