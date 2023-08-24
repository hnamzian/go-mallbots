package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/hnamzian/go-mallbots/internal/config"
	"github.com/rs/zerolog"
)

type App struct {
	cfg    *config.AppConfig
	rpc    *grpc.Server
	db     *sql.DB
	logger zerolog.Logger
}

func (a *App) connectDB() error {
	a.logger.Info().Msg("Connecting to database...")
	db, err := sql.Open("pgx", a.cfg.PG.Connection())
	if err != nil {
		a.logger.Error().Msg(
			fmt.Sprintf("failed to connect to database: %s", err),
		)
		return err
	}
	a.db = db
	a.logger.Info().Msg("Connected to database.")

	return nil
}

func (a *App) closeDB() error {
	a.logger.Warn().Msg("closing database connection")
	if err := a.db.Close(); err != nil {
		return err
	}
	a.logger.Warn().Msg("database connection closed")
	return nil
}

func (a *App) waitForWebServer(ctx context.Context) error {
	server := http.Server{
		Addr: a.cfg.Http.Address(),
	}

	group, gCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		a.logger.Info().Msg("web server started")
		defer a.logger.Warn().Msg("web server shut down")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("failed to start web server: %w", err)
		}
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		a.logger.Warn().Msg("web server to be shutdown")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			a.logger.Error().Msg("failed to shutdown web server")
			return err
		}
		return nil
	})

	return group.Wait()
}

func (a *App) waitForRPC(ctx context.Context) error {
	listener, err := net.Listen("tcp", a.cfg.Grpc.Address())
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	group, gCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		a.logger.Info().Msg("rpc server started")
		defer a.logger.Warn().Msg("rpc server shutdown")
		if err = a.rpc.Serve(listener); err != nil {
			return fmt.Errorf("failed to serve: %w", err)
		}
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		a.logger.Warn().Msg("rpc server to be shutdown")
		timer := time.NewTimer(5 * time.Second)
		stopped := make(chan struct{}, 1)
		go func() {
			a.rpc.GracefulStop()
			close(stopped)
		}()

		select {
		case <-timer.C:
			a.rpc.Stop()
			return fmt.Errorf("failed to gracefully shutdown: %w", ctx.Err())
		case <-stopped:
			a.logger.Warn().Msg("rpc server gracefully shutdown")
			return nil
		}
	})

	return group.Wait()
}
