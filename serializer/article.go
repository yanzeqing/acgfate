package serializer

import "acgfate/model"

// Article 文章序列化器
type Article struct {
	Aid       uint64 `json:"aid"`
	Author    User   `json:"author"`
	CreatedAt int64  `json:"created_at"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
	Category  uint16 `json:"category"`
	Status    string `json:"status"`
}

// BuildArticle 序列化文章
func BuildArticle(item model.Article) Article {
	author, _ := model.GetUser(item.Author)
	return Article{
		Aid:       item.Aid,
		Author:    BuildUser(author),
		CreatedAt: item.CreatedAt.Unix(),
		Title:     item.Title,
		Summary:   item.Summary,
		Content:   item.Content,
		Category:  item.Category,
		Status:    item.Status,
	}
}

// BuildArticles 序列化文章列表
func BuildArticles(items []model.Article) []Article {
	var articles []Article
	for _, item := range items {
		article := BuildArticle(item)
		articles = append(articles, article)
	}
	return articles
}
