package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func (a *App) waitForWebServer(ctx context.Context) error {
	server := http.Server{
		Addr: a.cfg.Http.Address(),
	}

	group, gCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		a.logger.Debug().Msg("web server started")
		defer fmt.Println("web server shut down")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start web server: %w", err)
		}
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		fmt.Println("web server to be shutdown")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})

	return group.Wait()
}
