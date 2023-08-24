package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/hnamzian/go-mallbots/internal/config"
	"github.com/hnamzian/go-mallbots/internal/logger"
	"github.com/hnamzian/go-mallbots/internal/waiter"
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

	w := waiter.NewWaiter()
	w.Add(
		app.waitForWebServer,
		app.waitForRPC,
	)
	return w.Wait()
}
