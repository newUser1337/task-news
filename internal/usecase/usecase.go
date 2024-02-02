package usecase

import (
	"github.com/newUser1337/task-news/internal/repository"
	"github.com/newUser1337/task-news/internal/usecase/article"
)

type Usecase struct {
	Article *article.Article
}

func NewUsecase(rep *repository.Repository) *Usecase {
	return &Usecase{
		Article: article.NewArticle(rep.Article),
	}
}
