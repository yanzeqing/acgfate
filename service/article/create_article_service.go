package article

import (
	"acgfate/model"
	"acgfate/serializer"
)

// CreateArticleService 文章投稿服务
type CreateArticleService struct {
	Title    string `form:"title" json:"title" binding:"required,min=1,max=30"`
	Summary  string `form:"summary" json:"summary" binding:"required,min=0,max=200"`
	Content  string `form:"content" json:"content" binding:"required"`
	Category uint16 `form:"category" json:"category" binding:"required"`
}

// articleValid 验证表单
func (service *CreateArticleService) articleValid() *serializer.Response {
	if _, ok := model.Category[service.Category]; !ok {
		return &serializer.Response{
			Code: 40008,
			Msg:  "分区错误",
		}
	}
	return nil
}

// Create 文章投稿
func (service *CreateArticleService) Create(user *model.User) serializer.Response {
	Article := model.Article{
		Author:   user.Uid,
		Title:    service.Title,
		Summary:  service.Summary,
		Content:  service.Content,
		Category: service.Category,
	}
	// 表单验证
	if err := service.articleValid(); err != nil {
		return *err
	}
	// 文章创建
	err := model.DB.Create(&Article).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "文章投稿失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildArticle(Article),
	}
}
