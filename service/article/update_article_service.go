package article

import (
	"acgfate/model"
	"acgfate/serializer"
)

// UpdateArticleService 文章更新服务
type UpdateArticleService struct {
	Title   string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Summary string `form:"summary" json:"summary" binding:"required,min=0,max=200"`
}

// Update 文章更新
func (service *UpdateArticleService) Update(id string) serializer.Response {
	var article model.Article
	err := model.DB.First(&article, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "文章不存在",
			Error: err.Error(),
		}
	}

	article.Title = service.Title
	article.Summary = service.Summary

	err = model.DB.Save(&article).Error
	if err != nil {
		return serializer.Response{
			Code:  50003,
			Msg:   "文章更新失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildArticle(article),
	}
}
