package article

import (
	"acgfate/model"
	"acgfate/serializer"
)

// DetailArticleService 文章详情服务
type DetailArticleService struct {
}

// Detail 文章详情
func (service *DetailArticleService) Detail(id string) serializer.Response {
	var article model.Article
	err := model.DB.First(&article, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "找不到内容",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildArticle(article),
	}
}
