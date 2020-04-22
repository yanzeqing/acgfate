package article

import (
	"acgfate/model"
	"acgfate/serializer"
)

// ListArticleService 文章列表服务
type ListArticleService struct {
}

// List 文章列表
func (service *ListArticleService) List(category string) serializer.Response {
	var articles []model.Article
	err := model.DB.Find(&articles, "category = ?", category).Error
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildArticles(articles),
	}
}
