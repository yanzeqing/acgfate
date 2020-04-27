package model

import (
	"time"
)

// Article 文章模型
type Article struct {
	Aid       uint64 `gorm:"primary_key"`
	Author    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Title     string
	Summary   string
	Content   string
	Category  uint16
	Status    string
}

// Category 分区模型
var Category = map[uint16]string{
	0: "默认",
	1: "软件",
	2: "音乐",
	3: "绘画",
	4: "测试",
}
