package context

import (
	"context"
	"sync"

	"github.com/newUser1337/task-news/internal/config"
)

type AppCtx struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	config     *config.Config
	wg         *sync.WaitGroup
}

func NewAppCtx() (*AppCtx, error) {
	config, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &AppCtx{
		ctx:        ctx,
		cancelFunc: cancelFunc,
		config:     config,
		wg:         &sync.WaitGroup{},
	}, nil
}

func (appCtx *AppCtx) GetGoContext() context.Context {
	return appCtx.ctx
}

func (appCtx *AppCtx) GetConfig() *config.Config {
	return appCtx.config
}

func (appCtx *AppCtx) Cancel() {
	appCtx.cancelFunc()
}

func (appCtx *AppCtx) WgAdd(delta int) {
	appCtx.wg.Add(delta)
}

func (appCtx *AppCtx) WgDone() {
	appCtx.wg.Done()
}

func (appCtx *AppCtx) WgWait() {
	appCtx.wg.Wait()
}
