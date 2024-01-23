package article

import (
	"context"

	"github.com/newUser1337/task-news/internal/entity"
)

type ArticleRepInf interface {
	GetArticles(context.Context) ([]*entity.NewsletterNewsItem, error)
	GetArticle(context.Context, int64) (*entity.NewsArticle, error)
}

type Article struct {
	articleRep ArticleRepInf
}

func NewArticle(articleRep ArticleRepInf) *Article {
	return &Article{
		articleRep: articleRep,
	}
}

func (a *Article) GetArticles(ctx context.Context) ([]*entity.NewsletterNewsItem, error) {
	return a.articleRep.GetArticles(ctx)
}

func (a *Article) GetArticle(ctx context.Context, id int64) (*entity.NewsArticle, error) {
	return a.articleRep.GetArticle(ctx, id)
}
