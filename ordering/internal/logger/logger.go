package logger

import (
	"context"

	"github.com/hnamzian/go-mallbots/ordering/internal/application"
	"github.com/hnamzian/go-mallbots/ordering/internal/domain"
	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger *zerolog.Logger
}

func NewApplication(app application.App, logger *zerolog.Logger) *Application {
	return &Application{App: app, logger: logger}
}

func (a Application) CreateOrder(ctx context.Context, create application.CreateOrder) (err error) {
	a.logger.Info().Msg("--> Ordering.CreateOrder")
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Ordering.CreateOrder")
		} else {
			a.logger.Info().Msg("<-- Ordering.CreateOrder")
		}
	}()
	return a.App.CreateOrder(ctx, create)
}
func (a Application) GetOrder(ctx context.Context, get application.GetOrder) (order *domain.Order, err error) {
	a.logger.Info().Msg("--> Ordering.GetOrder")
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Ordering.GetOrder")
		} else {
			a.logger.Info().Msg("<-- Ordering.GetOrder")
		}
	}()
	return a.App.GetOrder(ctx, get)
}
func (a Application) CancelOrder(ctx context.Context, cancel application.CancelOrder) (err error) {
	a.logger.Info().Msg("--> Ordering.CancelOrder")
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Ordering.CancelOrder")
		} else {
			a.logger.Info().Msg("<-- Ordering.CancelOrder")
		}
	}()
	return a.App.CancelOrder(ctx, cancel)
}
func (a Application) ReadyOrder(ctx context.Context, ready application.ReadyOrder) (err error) {
	a.logger.Info().Msg("--> Ordering.ReadyOrder")
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Ordering.ReadyOrder")
		} else {
			a.logger.Info().Msg("<-- Ordering.ReadyOrder")
		}
	}()
	return a.App.ReadyOrder(ctx, ready)
}
func (a Application) CompletedOrder(ctx context.Context, complete application.CompletedOrder) (err error) {
	a.logger.Info().Msg("--> Ordering.CompletedOrder")
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Ordering.CompletedOrder")
		} else {
			a.logger.Info().Msg("<-- Ordering.CompletedOrder")
		}
	}()
	return a.App.CompletedOrder(ctx, complete)
}
