package api

import (
	"acgfate/service/article"
	"github.com/gin-gonic/gin"
)

// CreateArticle 文章投稿接口
func CreateArticle(c *gin.Context) {
	createArticleService := article.CreateArticleService{}
	if err := c.ShouldBind(&createArticleService); err == nil {
		res := createArticleService.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DetailArticle 文章详情接口
func DetailArticle(c *gin.Context) {
	detailArticleService := article.DetailArticleService{}
	if err := c.ShouldBind(&detailArticleService); err == nil {
		res := detailArticleService.Detail(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListArticle 文章列表接口
func ListArticle(c *gin.Context) {
	listArticleService := article.ListArticleService{}
	if err := c.ShouldBind(&listArticleService); err == nil {
		res := listArticleService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateArticle 文章更新接口
func UpdateArticle(c *gin.Context) {
	updateArticleService := article.UpdateArticleService{}
	if err := c.ShouldBind(&updateArticleService); err == nil {
		res := updateArticleService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// DeleteArticle 文章删除接口
func DeleteArticle(c *gin.Context) {
	deleteArticleService := article.DeleteArticleService{}
	res := deleteArticleService.Delete(c.Param("id"))
	c.JSON(200, res)
}
