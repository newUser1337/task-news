package news

import (
	gocontext "context"
	"fmt"
	"log"
	"time"

	"github.com/newUser1337/task-news/internal/config"
	"github.com/newUser1337/task-news/internal/context"
	"github.com/newUser1337/task-news/internal/httpclient"
	"github.com/newUser1337/task-news/internal/repository"
	entityrep "github.com/newUser1337/task-news/internal/repository/entity"
	"github.com/newUser1337/task-news/internal/service/news/entity"
)

const articlesListEndpoint string = "https://www.htafc.com/api/incrowd/getnewlistinformation?count=50"
const articleEndpoint string = "https://www.htafc.com/api/incrowd/getnewsarticleinformation?id=%d"

type articleRepository interface {
	SaveArticleList(gocontext.Context, []*entityrep.NewsletterNewsItemDB) ([]int64, error)
	SaveArticle(gocontext.Context, *entityrep.NewsArticleDB) (int64, error)
}

type NewsServer struct {
	appCtx  *context.AppCtx
	article articleRepository
}

func StartNewNewsServer(appCtx *context.AppCtx, cfg *config.Config, rep *repository.Repository) {
	newsServer := &NewsServer{
		appCtx:  appCtx,
		article: rep.Article,
	}

	go newsServer.maintain(cfg.Core.RefreshInterval)
}

func (n *NewsServer) maintain(refreshTime time.Duration) {
	n.appCtx.WgAdd(1)
	defer n.appCtx.WgDone()

	httpClient := httpclient.NewClient(httpclient.WithTimeout(time.Second * 5))
	ticker := time.NewTicker(refreshTime)
	defer ticker.Stop()

	if err := n.getArticles(httpClient); err != nil {
		log.Println(err)
	}

	for range ticker.C {
		select {
		case <-n.appCtx.GetGoContext().Done():
			return
		default:
		}

		if err := n.getArticles(httpClient); err != nil {
			log.Println(err)
		}
	}
}

func (n *NewsServer) getArticles(httpClient *httpclient.Client) error {
	articlesExt := &entity.ArticleListExternal{}
	if err := httpClient.FetchData(articlesListEndpoint, articlesExt); err != nil {
		return err
	}
	if articlesExt == nil {
		return nil
	}

	articlesDb := make([]*entityrep.NewsletterNewsItemDB, 0, len(articlesExt.ArticleList.Articles))
	for _, article := range articlesExt.ArticleList.Articles {
		articleDb := &entityrep.NewsletterNewsItemDB{
			ArticleURL:        article.ArticleURL,
			NewsArticleID:     article.NewsArticleID,
			PublishDate:       article.PublishDate,
			Taxonomies:        article.Taxonomies,
			TeaserText:        article.TeaserText,
			ThumbnailImageURL: article.ThumbnailImageURL,
			Title:             article.Title,
			OptaMatchId:       article.OptaMatchId,
			LastUpdateDate:    article.LastUpdateDate,
			IsPublished:       article.IsPublished,
		}

		articlesDb = append(articlesDb, articleDb)
	}
	insertedArticlesId, err := n.article.SaveArticleList(n.appCtx.GetGoContext(), articlesDb)
	if err != nil {
		return err
	}

	for _, id := range insertedArticlesId {
		articleExt := &entity.ArticleExternal{}
		if err := httpClient.FetchData(fmt.Sprintf(articleEndpoint, id), articleExt); err != nil {
			return err
		}
		if articleExt.NewsArticle == nil {
			continue
		}

		articleDb := &entityrep.NewsArticleDB{
			ArticleURL:        articleExt.NewsArticle.ArticleURL,
			NewsArticleID:     articleExt.NewsArticle.NewsArticleID,
			PublishDate:       articleExt.NewsArticle.PublishDate,
			Taxonomies:        articleExt.NewsArticle.Taxonomies,
			TeaserText:        articleExt.NewsArticle.TeaserText,
			Subtitle:          articleExt.NewsArticle.Subtitle,
			ThumbnailImageURL: articleExt.NewsArticle.ThumbnailImageURL,
			Title:             articleExt.NewsArticle.Title,
			BodyText:          articleExt.NewsArticle.BodyText,
			GalleryImageURLs:  articleExt.NewsArticle.GalleryImageURLs,
			VideoURL:          articleExt.NewsArticle.VideoURL,
			OptaMatchId:       articleExt.NewsArticle.OptaMatchId,
			LastUpdateDate:    articleExt.NewsArticle.LastUpdateDate,
			IsPublished:       articleExt.NewsArticle.IsPublished,
		}

		if _, err := n.article.SaveArticle(n.appCtx.GetGoContext(), articleDb); err != nil {
			return err
		}
	}

	return nil
}
