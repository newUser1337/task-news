package article

import (
	"context"
	"fmt"
	"log"

	"github.com/newUser1337/task-news/internal/newserror"
	"github.com/newUser1337/task-news/internal/repository/entity"
	"github.com/newUser1337/task-news/internal/repository/mgdb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleRepository struct {
	articleTitle *mongo.Collection
	article      *mongo.Collection
}

func NewArticleRep(client *mgdb.MongoDriver) *ArticleRepository {
	return &ArticleRepository{
		articleTitle: client.GetCollection("news", "articleTitle"),
		article:      client.GetCollection("news", "article"),
	}
}

func (a *ArticleRepository) SaveArticleList(ctx context.Context, articleList []*entity.NewsletterNewsItemDB) ([]int64, error) {
	articles := make([]interface{}, 0, len(articleList))
	for _, article := range articleList {
		articles = append(articles, article)
	}
	insertedIds, err := a.articleTitle.InsertMany(ctx, articles)
	if err != nil {
		return nil, newserror.NewErrorNews(
			newserror.DbErr,
			"repository: failed to save article list",
			err.Error(),
		)
	}
	result := make([]int64, 0, len(insertedIds.InsertedIDs))
	for _, insertedId := range insertedIds.InsertedIDs {
		id, ok := insertedId.(int64)
		if !ok {
			log.Printf("repository: unexpected type of id %v", insertedId)
			continue
		}
		result = append(result, int64(id))
	}
	return result, nil
}

func (a *ArticleRepository) SaveArticle(ctx context.Context, article *entity.NewsArticleDB) (int64, error) {
	result, err := a.article.InsertOne(ctx, article)
	if err != nil {
		return 0, newserror.NewErrorNews(
			newserror.DbErr,
			fmt.Sprintf("repository: failed to instert article with id %d", article.NewsArticleID),
			"",
		)
	}

	id, ok := result.InsertedID.(int64)
	if !ok {
		return 0, newserror.NewErrorNews(
			newserror.DbErr,
			fmt.Sprintf("repository: failed to get atricle's id unexpected type %v", result.InsertedID),
			"",
		)
	}
	return id, nil
}

func (a *ArticleRepository) GetArticles(ctx context.Context) ([]*entity.NewsletterNewsItemDB, error) {
	cur, err := a.articleTitle.Find(ctx, bson.D{})
	if err != nil {
		return nil, newserror.NewErrorNews(
			newserror.DbErr,
			"repository: failed to find articleTitle",
			err.Error(),
		)
	}
	defer cur.Close(ctx)

	articles := []*entity.NewsletterNewsItemDB{}
	for cur.Next(ctx) {
		article := &entity.NewsletterNewsItemDB{}
		if err := cur.Decode(article); err != nil {
			return nil, newserror.NewErrorNews(
				newserror.InternalErr,
				"repository: failed to decode article",
				err.Error(),
			)
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (a *ArticleRepository) GetArticle(ctx context.Context, id int64) (*entity.NewsArticleDB, error) {
	article := &entity.NewsArticleDB{}
	if err := a.article.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(article); err != nil {
		return nil, newserror.NewErrorNews(
			newserror.InternalErr,
			fmt.Sprintf("repository: failed to find article with id %d", id),
			err.Error(),
		)
	}
	return article, nil
}
