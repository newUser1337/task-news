package router

import (
	"context"

	"github.com/newUser1337/task-news/internal/controller"
	"github.com/newUser1337/task-news/internal/usecase"

	"github.com/gorilla/mux"
)

func NewRouter(ctx context.Context, uc *usecase.Usecase) *mux.Router {
	r := mux.NewRouter()

	// Articles
	{
		articles := controller.NewArticlesHandler(uc.Article)
		r.HandleFunc("/articles", articles.GetArticles).Methods("GET")
		r.HandleFunc("/articles/{id}", articles.GetArticle).Methods("GET")
	}

	return r
}
