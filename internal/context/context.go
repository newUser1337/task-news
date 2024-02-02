package context

import (
	"context"
	"sync"
)

type AppCtx struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	wg         *sync.WaitGroup
}

func NewAppCtx() (*AppCtx, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &AppCtx{
		ctx:        ctx,
		cancelFunc: cancelFunc,
		wg:         &sync.WaitGroup{},
	}, nil
}

func (appCtx *AppCtx) GetGoContext() context.Context {
	return appCtx.ctx
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
