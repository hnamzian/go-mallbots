package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/hnamzian/go-mallbots/customers"
	"github.com/hnamzian/go-mallbots/internal/config"
	"github.com/hnamzian/go-mallbots/internal/logger"
	"github.com/hnamzian/go-mallbots/internal/module"
	"github.com/hnamzian/go-mallbots/internal/waiter"
	"github.com/hnamzian/go-mallbots/stores"
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

	app.waiter = waiter.NewWaiter()

	opts := []grpc.ServerOption{}
	app.rpc = grpc.NewServer(opts...)
	reflection.Register(app.rpc)

	app.mux = chi.NewMux()

	app.modules = []module.Module{
		customers.Module{},
		stores.Module{},
	}
	if err = app.startupModules(); err != nil {
		return err
	}

	app.waiter.Add(
		app.waitForWebServer,
		app.waitForRPC,
	)

	return app.waiter.Wait()
}
