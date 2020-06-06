package v1

import (
	"acgfate/service/article"
	"github.com/gin-gonic/gin"
)

// CreateArticle 文章投稿
func CreateArticle(c *gin.Context) {
	user := CurrentUser(c)
	createArticleService := article.CreateArticleService{}
	if err := c.ShouldBind(&createArticleService); err == nil {
		res := createArticleService.Create(user)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DetailArticle 文章详情
func DetailArticle(c *gin.Context) {
	detailArticleService := article.DetailArticleService{}
	if err := c.ShouldBind(&detailArticleService); err == nil {
		res := detailArticleService.Detail(c.Param("aid"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListArticle 文章分区列表
func ListArticle(c *gin.Context) {
	listArticleService := article.ListArticleService{}
	if err := c.ShouldBind(&listArticleService); err == nil {
		res := listArticleService.CategoryList(c.Param("category"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AuthorArticle 文章作者列表
func AuthorArticle(c *gin.Context) {
	userPostService := article.ListArticleService{}
	if err := c.ShouldBind(&userPostService); err == nil {
		res := userPostService.AuthorList(c.Param("author"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateArticle 文章更新
func UpdateArticle(c *gin.Context) {
	updateArticleService := article.UpdateArticleService{}
	if err := c.ShouldBind(&updateArticleService); err == nil {
		res := updateArticleService.Update(c.Param("aid"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// DeleteArticle 文章删除
func DeleteArticle(c *gin.Context) {
	user := CurrentUser(c)
	deleteArticleService := article.DeleteArticleService{}
	res := deleteArticleService.Delete(c.Param("aid"), user)
	c.JSON(200, res)
}
