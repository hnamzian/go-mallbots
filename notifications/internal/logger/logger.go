package logger

import (
	"context"

	"github.com/hnamzian/go-mallbots/notifications/internal/application"
	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

func NewApplication(app application.App, logger zerolog.Logger) Application {
	return Application{
		App:    app,
		logger: logger,
	}
}

func (a Application) NotifyOrderCreated(ctx context.Context, notify application.OrderCreated) (err error) {
	a.logger.Info().Msg("--> Notifications.NotifyOrderCreated")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Notifications.NotifyOrderCreated")
		} else {
			a.logger.Info().Msg("<-- Notifications.NotifyOrderCreated") 
		}
	}()
	return nil
}

func (a Application) NotifyOrderCanceled(ctx context.Context, notify application.OrderCanceled) (err error) {
	a.logger.Info().Msg("--> Notifications.NotifyOrderCanceled")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Notifications.NotifyOrderCanceled")
		} else {
			a.logger.Info().Msg("<-- Notifications.NotifyOrderCanceled") 
		}
	}()
	return nil
}

func (a Application) NotifyOrderReady(ctx context.Context, notify application.OrderReady) (err error) {
	a.logger.Info().Msg("--> Notifications.NotifyOrderReady")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Notifications.NotifyOrderReady")
		} else {
			a.logger.Info().Msg("<-- Notifications.NotifyOrderReady") 
		}
	}()
	return nil
}
