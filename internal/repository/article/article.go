package article

import (
	"context"
	"fmt"
	"log"

	"github.com/newUser1337/task-news/internal/entity"
	"github.com/newUser1337/task-news/internal/repository/mgdb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleRep struct {
	articleTitle *mongo.Collection
	article      *mongo.Collection
}

func NewArticleRep(client *mgdb.MongoDriver) *ArticleRep {
	return &ArticleRep{
		articleTitle: client.CreateCollection("news", "articleTitle"),
		article:      client.CreateCollection("news", "article"),
	}
}

func (a *ArticleRep) SaveArticleList(ctx context.Context, articleList *entity.ArticleList) ([]int64, error) {
	articles := make([]interface{}, 0, len(articleList.ArticleList.Articles))
	for _, article := range articleList.ArticleList.Articles {
		articles = append(articles, article)
	}
	insertedIds, err := a.articleTitle.InsertMany(ctx, articles)
	if err != nil {
		return nil, err
	}
	result := make([]int64, 0, len(insertedIds.InsertedIDs))
	for _, insertedId := range insertedIds.InsertedIDs {
		id, ok := insertedId.(int64)
		if !ok {
			log.Printf("unexpected type of id %v", insertedId)
			continue
		}
		result = append(result, int64(id))
	}
	return result, err
}

func (a *ArticleRep) SaveArticle(ctx context.Context, article *entity.Article) (int64, error) {
	result, err := a.article.InsertOne(ctx, article.NewsArticle)
	if err != nil {
		return 0, err
	}

	id, ok := result.InsertedID.(int64)
	if !ok {
		log.Printf("unexpected type of id %v", result.InsertedID)
		return 0, fmt.Errorf("unexpected type of id %v", result.InsertedID)
	}
	return id, nil
}

func (a *ArticleRep) GetArticles(ctx context.Context) ([]*entity.NewsletterNewsItem, error) {
	cur, err := a.articleTitle.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	articles := []*entity.NewsletterNewsItem{}
	for cur.Next(ctx) {
		article := new(entity.NewsletterNewsItem)
		if err := cur.Decode(article); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (a *ArticleRep) GetArticle(ctx context.Context, id int64) (*entity.NewsArticle, error) {
	article := new(entity.NewsArticle)
	if err := a.article.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(article); err != nil {
		return nil, err
	}
	return article, nil
}
