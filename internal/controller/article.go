package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/newUser1337/task-news/internal/usecase/entity"

	"github.com/gorilla/mux"
)

type articleUsecase interface {
	GetArticles(context.Context) ([]*entity.NewsletterNewsItemResponse, error)
	GetArticle(context.Context, int64) (*entity.NewsArticleResponse, error)
}

type ArticleHandler struct {
	article articleUsecase
}

func NewArticlesHandler(article articleUsecase) *ArticleHandler {
	return &ArticleHandler{
		article: article,
	}
}

func (a *ArticleHandler) GetArticles(w http.ResponseWriter, req *http.Request) {
	articles, err := a.article.GetArticles(req.Context())
	if err != nil {
		log.Println("failed to get articles", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	response, err := json.Marshal(articles)
	if err != nil {
		log.Println("failed to marshal articles", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response); err != nil {
		log.Printf("GetArticles: failed to write response %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func (a *ArticleHandler) GetArticle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("failed to parse parameter int", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	article, err := a.article.GetArticle(req.Context(), id)
	if err != nil {
		log.Printf("failed to get article with key %d: %v\n", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := json.Marshal(article)
	if err != nil {
		log.Println("failed to marshal articles", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response); err != nil {
		log.Printf("GetArticle: failed to write response %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
