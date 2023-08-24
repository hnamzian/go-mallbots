package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/sync/errgroup"

	"github.com/hnamzian/go-mallbots/internal/config"
	"github.com/hnamzian/go-mallbots/internal/logger"
)

func main() {
	run()
}

func run() error {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	app := &App{
		cfg: cfg,
	}
	app.logger = logger.NewLogger(logger.LoggerConfig{
		Level: logger.Level(cfg.LogLevel),
	})

	if err := app.connectDB(); err != nil {
		return err
	}
	defer app.closeDB()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ctx, cancel = signal.NotifyContext(ctx, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		<-gCtx.Done()
		cancel()
		return nil
	})
	group.Go(func() error {
		return app.waitForWebServer(gCtx)
	})
	group.Go(func() error {
		return app.waitForRPC(gCtx)
	})

	return group.Wait()
}
