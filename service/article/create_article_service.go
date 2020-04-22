package article

import (
	"acgfate/model"
	"acgfate/serializer"
)

// CreateArticleService 文章投稿服务
type CreateArticleService struct {
	Title    string `form:"title" json:"title" binding:"required,min=1,max=30"`
	Summary  string `form:"summary" json:"summary" binding:"required,min=0,max=200"`
	Category string `form:"category" json:"category" binding:"required,min=1,max=5"`
}

// Create 文章创建
func (service *CreateArticleService) Create(user *model.User) serializer.Response {
	Article := model.Article{
		Title:    service.Title,
		Summary:  service.Summary,
		Category: service.Category,
		Author:   user.Uid,
	}
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
