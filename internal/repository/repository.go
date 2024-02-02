package repository

import (
	"github.com/newUser1337/task-news/internal/repository/article"
	"github.com/newUser1337/task-news/internal/repository/mgdb"
)

type Repository struct {
	Article *article.ArticleRepository
}

func NewRepository(clientDb any) *Repository {
	dbConnector := mgdb.NewMongoClient(clientDb)

	return &Repository{
		Article: article.NewArticleRep(dbConnector),
	}
}
