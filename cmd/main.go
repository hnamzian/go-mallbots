package main

import (
	"context"

	"github.com/hnamzian/go-mallbots/internal/config"
	"github.com/hnamzian/go-mallbots/internal/logger"
	"github.com/rs/zerolog"
)

type App struct {
	cfg    *config.AppConfig
	logger zerolog.Logger
}

func main() {
	run()
}

func run() {
	cfg := config.InitConfig()
	app := &App{
		cfg: cfg,
	}
	app.logger = logger.NewLogger(logger.LoggerConfig{
		Level: logger.Level(cfg.LogLevel),
	})

	ctx := context.Background()
	app.waitForWebServer(ctx)
}