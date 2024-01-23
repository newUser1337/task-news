package usecase

import (
	"context"

	"github.com/newUser1337/task-news/internal/entity"
	"github.com/newUser1337/task-news/internal/repository"
	"github.com/newUser1337/task-news/internal/usecase/article"
)

type ArticleInf interface {
	GetArticles(context.Context) ([]*entity.NewsletterNewsItem, error)
	GetArticle(context.Context, int64) (*entity.NewsArticle, error)
}

type Usecase struct {
	Article ArticleInf
}

func NewUsecase(rep *repository.Repository) *Usecase {
	return &Usecase{
		Article: article.NewArticle(rep.ArticleRep),
	}
}
