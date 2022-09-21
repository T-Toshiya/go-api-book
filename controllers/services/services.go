package services

import "go-api-book/models"

type MyAppServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	ListArticleService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)

	PostCommentService(comment models.Comment) (models.Comment, error)
}
