package repository

import (
	"context"

	"github.com/newUser1337/task-news/internal/entity"
	"github.com/newUser1337/task-news/internal/repository/article"
	"github.com/newUser1337/task-news/internal/repository/mgdb"
)

type ArticleRepInf interface {
	SaveArticleList(context.Context, *entity.ArticleList) ([]int64, error)
	SaveArticle(context.Context, *entity.Article) (int64, error)
	GetArticles(context.Context) ([]*entity.NewsletterNewsItem, error)
	GetArticle(context.Context, int64) (*entity.NewsArticle, error)
}

type Repository struct {
	ArticleRep ArticleRepInf
}

func NewRepository(clientDb any) *Repository {
	dbConnector := mgdb.NewMongoClient(clientDb)

	return &Repository{
		ArticleRep: article.NewArticleRep(dbConnector),
	}
}
