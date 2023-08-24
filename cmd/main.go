package main

import (
	"context"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/hnamzian/go-mallbots/internal/config"
	"github.com/hnamzian/go-mallbots/internal/logger"
)

func main() {
	run()
}

func run() error {
	cfg := config.InitConfig()
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
	app.waitForWebServer(ctx)

	return nil
}
