package serializer

import "acgfate/model"

// Article 文章序列化器
type Article struct {
	Aid       uint64 `json:"aid"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Auther    string `json:"auther"`
	State     int32  `json:"state"`
	Category  string `json:"category"`
	CreatedAt int64  `json:"created_at"`
}

// BuildArticle 序列化文章
func BuildArticle(item model.Article) Article {
	return Article{
		Aid:       item.Aid,
		Title:     item.Title,
		Summary:   item.Summary,
		Auther:    item.Auther,
		State:     item.State,
		Category:  item.Category,
		CreatedAt: item.CreatedAt.Unix(),
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
