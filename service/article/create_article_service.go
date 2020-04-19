package article

import (
	"acgfate/model"
	"acgfate/serializer"
	"github.com/gin-gonic/gin"
)

// CreateArticleService 文章投稿服务
type CreateArticleService struct {
	Title    string `form:"title" json:"title" binding:"required,min=1,max=30"`
	Summary  string `form:"summary" json:"summary" binding:"required,min=0,max=200"`
	Category string `form:"category" json:"category" binding:"required,min=1,max=5"`
}

// 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// Create 文章创建
func (service *CreateArticleService) Create(c *gin.Context) serializer.Response {
	Article := model.Article{
		Title:    service.Title,
		Summary:  service.Summary,
		Auther:   CurrentUser(c).UserName,
		Category: service.Category,
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
