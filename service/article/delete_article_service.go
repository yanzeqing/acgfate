package article

import (
	"acgfate/model"
	"acgfate/serializer"
)

// DeleteArticleService 文章删除服务
type DeleteArticleService struct {
	Title   string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Summary string `form:"summary" json:"summary" binding:"required,min=0,max=200"`
}

// Delete 文章删除
func (service *DeleteArticleService) Delete(aid string, user *model.User) serializer.Response {
	var article model.Article

	err := model.DB.First(&article, aid).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "文章不存在",
			Error: err.Error(),
		}
	}
	err = model.DB.First(&article, aid).Select("author").Error
	if article.Author != user.Uid {
		return serializer.UserErr("", err)
	}
	err = model.DB.Delete(&article).Error
	if err != nil {
		return serializer.Response{
			Code:  50004,
			Msg:   "文章删除失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Msg: "文章删除成功",
	}
}
