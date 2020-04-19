package model

import (
	"time"
)

// Article 文章模型
type Article struct {
	Aid       uint64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Title     string
	Summary   string
	Auther    string
	State     int32
	Category  string
}
