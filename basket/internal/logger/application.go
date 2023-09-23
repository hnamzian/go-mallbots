package logger

import (
	"context"

	"github.com/hnamzian/go-mallbots/basket/internal/application"
	"github.com/hnamzian/go-mallbots/basket/internal/domain"
	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

func NewApplication(logger zerolog.Logger, app application.App) *Application {
	return &Application{App: app, logger: logger}
}

func (a Application) StartBasket(ctx context.Context, start *application.StartBasket) (err error) {
	a.logger.Info().Msgf("Baskets.StartBasket: %v", start)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Baskets.StartBasket")
		} else {
			a.logger.Info().Msg("<-- Baskets.StartBasket")
		}
	}()
	return a.App.StartBasket(ctx, start)
}

func (a Application) CancelBasket(ctx context.Context, cancel *application.CancelBasket) (err error) {
	a.logger.Info().Msgf("Baskets.CancelBasket: %v", cancel)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Baskets.CancelBasket")
		} else {
			a.logger.Info().Msg("<-- Baskets.CancelBasket")
		}
	}()
	return a.App.CancelBasket(ctx, cancel)
}

func (a Application) CheckoutBasket(ctx context.Context, checkout application.CheckoutBasket) (err error) {
	a.logger.Info().Msgf("Baskets.CheckoutBasket: %v", checkout)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Baskets.CheckoutBasket")
		} else {
			a.logger.Info().Msg("<-- Baskets.CheckoutBasket")
		}
	}()
	return a.App.CheckoutBasket(ctx, checkout)
}

func (a Application) AddItem(ctx context.Context, add *application.AddItem) (err error) {
	a.logger.Info().Msgf("Baskets.CheckoutBasket: %v", add)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Baskets.CheckoutBasket")
		} else {
			a.logger.Info().Msg("<-- Baskets.CheckoutBasket")
		}
	}()
	return a.App.AddItem(ctx, add)
}

func (a Application) RemoveItem(ctx context.Context, remove *application.RemoveItem) (err error) {
	a.logger.Info().Msgf("Baskets.CheckoutBasket: %v", remove)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Baskets.CheckoutBasket")
		} else {
			a.logger.Info().Msg("<-- Baskets.CheckoutBasket")
		}
	}()
	return a.App.RemoveItem(ctx, remove)
}

func (a Application) GetBasket(ctx context.Context, get *application.GetBasket) (basket *domain.Basket, err error) {
	a.logger.Info().Msgf("Baskets.CheckoutBasket: %v", get)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Baskets.CheckoutBasket")
		} else {
			a.logger.Info().Msg("<-- Baskets.CheckoutBasket")
		}
	}()
	return a.App.GetBasket(ctx, get)
}
