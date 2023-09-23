package logger

import (
	"context"

	"github.com/hnamzian/go-mallbots/depot/internal/application"
	"github.com/hnamzian/go-mallbots/depot/internal/domain"
	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

func NewApplication(logger zerolog.Logger, app application.App) Application {
	return Application{App: app, logger: logger}
}

func (a Application) CreateShoppingList(ctx context.Context, create *application.CreateShoppingList) (err error) {
	a.logger.Info().Msgf("Depot.CreateShoppingList: %v", create)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Depot.CreateShoppingList")
		} else {
			a.logger.Info().Msg("<-- Depot.CreateShoppingList")
		}
	}()
	return a.App.CreateShoppingList(ctx, create)
}

func (a Application) CancelShoppingList(ctx context.Context, cancel *application.CancelShoppingList) (err error) {
	a.logger.Info().Msgf("Depot.CancelShoppingList: %v", cancel)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Depot.CancelShoppingList")
		} else {
			a.logger.Info().Msg("<-- Depot.CancelShoppingList")
		}
	}()
	return a.App.CancelShoppingList(ctx, cancel)
}

func (a Application) CompleteShoppingList(ctx context.Context, complete *application.CompleteShoppingList) (err error) {
	a.logger.Info().Msgf("Depot.CompleteShoppingList: %v", complete)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Depot.CompleteShoppingList")
		} else {
			a.logger.Info().Msg("<-- Depot.CompleteShoppingList")
		}
	}()
	return a.App.CompleteShoppingList(ctx, complete)
}

func (a Application) AssignBotToShoppingList(ctx context.Context, assign *application.AssignBotToShoppingList) (err error) {
	a.logger.Info().Msgf("Depot.StartBasket: %AssignBotToShoppingList", assign)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Depot.AssignBotToShoppingList")
		} else {
			a.logger.Info().Msg("<-- Depot.AssignBotToShoppingList")
		}
	}()
	return a.App.AssignBotToShoppingList(ctx, assign)
}

func (a Application) GetShoppingList(ctx context.Context, get *application.GetShoppingList) (list *domain.ShoppingList, err error) {
	a.logger.Info().Msgf("Depot.GetShoppingList: %v", get)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Depot.GetShoppingList")
		} else {
			a.logger.Info().Msg("<-- Depot.GetShoppingList")
		}
	}()
	return a.App.GetShoppingList(ctx, get)
}
