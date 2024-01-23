package news

import (
	gocontext "context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/newUser1337/task-news/internal/context"
	"github.com/newUser1337/task-news/internal/entity"
	"github.com/newUser1337/task-news/internal/repository"
)

const articlesListEndpoint string = "https://www.htafc.com/api/incrowd/getnewlistinformation?count=50"
const articleEndpoint string = "https://www.htafc.com/api/incrowd/getnewsarticleinformation?id=%d"

type articleRepInf interface {
	SaveArticleList(gocontext.Context, *entity.ArticleList) ([]int64, error)
	SaveArticle(gocontext.Context, *entity.Article) (int64, error)
}

type NewsServer struct {
	appCtx      *context.AppCtx
	articlesRep articleRepInf
}

func StartNewNewsServer(appCtx *context.AppCtx, rep *repository.Repository) {
	newsServer := &NewsServer{
		appCtx:      appCtx,
		articlesRep: rep.ArticleRep,
	}

	go newsServer.maintain()
}

func (n *NewsServer) maintain() {
	n.appCtx.WgAdd(1)
	defer n.appCtx.WgDone()

	refreshTime := n.appCtx.GetConfig().RefreshInterval
	timer := time.NewTimer(refreshTime)

	if err := n.getArticels(); err != nil {
		log.Printf("failed to get articels %v", err)
	}

	for {
		timer.Reset(refreshTime)
		select {
		case <-n.appCtx.GetGoContext().Done():
			if !timer.Stop() {
				<-timer.C
			}
			return
		case <-timer.C:
		}

		n.getArticels()
	}
}

func (n *NewsServer) getArticels() error {
	articles := new(entity.ArticleList)
	if err := fetchData(articlesListEndpoint, articles); err != nil {
		return err
	}

	insertedArticlesId, err := n.articlesRep.SaveArticleList(n.appCtx.GetGoContext(), articles)
	if err != nil {
		return err
	}

	for _, id := range insertedArticlesId {
		article := new(entity.Article)
		if err := fetchData(fmt.Sprintf(articleEndpoint, id), article); err != nil {
			return err
		}

		if _, err := n.articlesRep.SaveArticle(n.appCtx.GetGoContext(), article); err != nil {
			return err
		}
	}

	return nil
}

func fetchData(url string, data any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := xml.Unmarshal(body, data); err != nil {
		return err
	}
	return nil
}
