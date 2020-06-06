package article

import (
	"acgfate/model"
	"acgfate/serializer"
)

// ListArticleService 文章列表服务
type ListArticleService struct {
}

// CategoryList 按分区展示文章列表
func (service *ListArticleService) CategoryList(category string) serializer.Response {
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

// AuthorList 按作者展示文章列表
func (service *ListArticleService) AuthorList(author string) serializer.Response {
	var articles []model.Article
	count := 0
	err := model.DB.Find(&articles, "author = ?", author).Count(&count).Error
	if count == 0 {
		return serializer.Response{
			Code:  404,
			Msg:   "当前用户无投稿",
			Error: err.Error(),
		}
	}
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
