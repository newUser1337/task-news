package article

import (
	"context"

	entityrep "github.com/newUser1337/task-news/internal/repository/entity"
	"github.com/newUser1337/task-news/internal/usecase/entity"
)

type ArticleRepository interface {
	GetArticles(context.Context) ([]*entityrep.NewsletterNewsItemDB, error)
	GetArticle(context.Context, int64) (*entityrep.NewsArticleDB, error)
}

type Article struct {
	article ArticleRepository
}

func NewArticle(article ArticleRepository) *Article {
	return &Article{
		article: article,
	}
}

func (a *Article) GetArticles(ctx context.Context) ([]*entity.NewsletterNewsItemResponse, error) {
	articlesDb, err := a.article.GetArticles(ctx)
	if err != nil {
		return nil, err
	}
	articles := make([]*entity.NewsletterNewsItemResponse, 0, len(articlesDb))
	for _, articleDb := range articlesDb {
		article := &entity.NewsletterNewsItemResponse{
			ArticleURL:        articleDb.ArticleURL,
			NewsArticleID:     articleDb.NewsArticleID,
			PublishDate:       articleDb.PublishDate,
			Taxonomies:        articleDb.Taxonomies,
			TeaserText:        articleDb.TeaserText,
			ThumbnailImageURL: articleDb.ThumbnailImageURL,
			Title:             articleDb.Title,
			OptaMatchId:       articleDb.OptaMatchId,
			LastUpdateDate:    articleDb.LastUpdateDate,
			IsPublished:       articleDb.IsPublished,
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (a *Article) GetArticle(ctx context.Context, id int64) (*entity.NewsArticleResponse, error) {
	articleDb, err := a.article.GetArticle(ctx, id)
	if err != nil {
		return nil, err
	}
	article := &entity.NewsArticleResponse{
		ArticleURL:        articleDb.ArticleURL,
		NewsArticleID:     articleDb.NewsArticleID,
		PublishDate:       articleDb.PublishDate,
		Taxonomies:        articleDb.Taxonomies,
		Subtitle:          articleDb.Subtitle,
		ThumbnailImageURL: articleDb.ThumbnailImageURL,
		Title:             articleDb.Title,
		BodyText:          articleDb.BodyText,
		GalleryImageURLs:  articleDb.GalleryImageURLs,
		VideoURL:          articleDb.VideoURL,
		OptaMatchId:       articleDb.OptaMatchId,
		LastUpdateDate:    articleDb.LastUpdateDate,
		IsPublished:       articleDb.IsPublished,
	}

	return article, nil
}
