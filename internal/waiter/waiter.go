package waiter

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type WaitFunc func(ctx context.Context) error

type Waiter interface {
	Add()
	Wait()
}

type waiter struct {
	ctx    context.Context
	cancel context.CancelFunc
	fns    []WaitFunc
}

func NewWaiter() *waiter {
	w := &waiter{
		fns: []WaitFunc{},
	}

	ctx := context.Background()
	w.ctx, w.cancel = context.WithCancel(ctx)
	w.ctx, w.cancel = signal.NotifyContext(ctx, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	return w
}

func (w *waiter) Add(fns ...WaitFunc) {
	w.fns = append(w.fns, fns...)
}

func (w *waiter) Wait() error {
	group, gCtx := errgroup.WithContext(w.ctx)

	group.Go(func() error {
		<-w.ctx.Done()
		w.cancel()
		return nil
	})

	for _, fn := range w.fns {
		fn := fn
		group.Go(func() error { return fn(gCtx) })
	}

	return group.Wait()
}
