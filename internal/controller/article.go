package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/newUser1337/task-news/internal/entity"

	"github.com/gorilla/mux"
)

type articleInf interface {
	GetArticles(context.Context) ([]*entity.NewsletterNewsItem, error)
	GetArticle(context.Context, int64) (*entity.NewsArticle, error)
}

type ArticleHandler struct {
	article articleInf
	ctx     context.Context
}

func NewArticlesHandler(ctx context.Context, article articleInf) *ArticleHandler {
	return &ArticleHandler{
		article: article,
		ctx:     ctx,
	}
}

func (a *ArticleHandler) GetArticles(w http.ResponseWriter, req *http.Request) {
	articles, err := a.article.GetArticles(a.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("failed to get articles", err)
		return
	}
	response, err := json.Marshal(articles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("failed to marshal articles", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (a *ArticleHandler) GetArticle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("failed to parse parameter int", err)
		return
	}

	article, err := a.article.GetArticle(a.ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to get article with key %d: %v\n", id, err)
		return
	}

	response, err := json.Marshal(article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("failed to marshal articles", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
