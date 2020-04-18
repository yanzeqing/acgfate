package model

import "github.com/jinzhu/gorm"

// Article 文章模型
type Article struct {
	gorm.Model
	Aid      uint64
	Title    string
	Summary  string
	Auther   string
	State    int32
	Category string
}
