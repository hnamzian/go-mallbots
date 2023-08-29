package logger

import (
	"context"

	"github.com/hnamzian/go-mallbots/stores/internal/application"
	"github.com/hnamzian/go-mallbots/stores/internal/domain"
	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

func NewApplication(app application.App, logger zerolog.Logger) Application {
	return Application{
		App: app,
		logger: logger,
	}
}

func (a Application) CreateStore(ctx context.Context, id, name, location string) (err error) {
	a.logger.Info().Msg("--> Stores.CreateStore")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Stores.CreateStore")
		} else {
			a.logger.Info().Msg("<-- Stores.CreateStore") 
		}
	}()
	return a.App.CreateStore(ctx, id, name, location)
}

func (a Application) GetStore(ctx context.Context, id string) (store *domain.Store, err error) {
	a.logger.Info().Msg("--> Stores.GetStore")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Stores.GetStore")
		} else {
			a.logger.Info().Msg("<-- Stores.GetStore") 
		}
	}()
	return a.App.GetStore(ctx, id)
}

func (a Application) GetStores(ctx context.Context) (stores []*domain.Store, err error) {
	a.logger.Info().Msg("--> Stores.GetStores")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Stores.GetStores")
		} else {
			a.logger.Info().Msg("<-- Stores.GetStores") 
		}
	}()
	return a.App.GetStores(ctx)
}

func (a Application) EnableParticipation(ctx context.Context, id string) (err error) {
	a.logger.Info().Msg("--> Stores.EnableParticipation")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Stores.EnableParticipation")
		} else {
			a.logger.Info().Msg("<-- Stores.EnableParticipation") 
		}
	}()
	return a.App.EnableParticipation(ctx, id)
}

func (a Application) DisableParticipation(ctx context.Context, id string) (err error) {
	a.logger.Info().Msg("--> Stores.DisableParticipation")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Stores.DisableParticipation")
		} else {
			a.logger.Info().Msg("<-- Stores.DisableParticipation") 
		}
	}()
	return a.App.DisableParticipation(ctx, id)
}

func (a Application) GetParticipatingStores(ctx context.Context) (stores []*domain.Store, err error) {
	a.logger.Info().Msg("--> Stores.GetParticipatingStores")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Stores.GetParticipatingStores")
		} else {
			a.logger.Info().Msg("<-- Stores.GetParticipatingStores") 
		}
	}()
	return a.App.GetParticipatingStores(ctx)
}

func (a Application) CreateProduct(ctx context.Context, id string, storeID string, name string, description string, sku string, price float64) (product *domain.Product, err error) {
	a.logger.Info().Msg("--> Stores.CreateProduct")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Stores.CreateProduct")
		} else {
			a.logger.Info().Msg("<-- Stores.CreateProduct") 
		}
	}()
	return a.App.CreateProduct(ctx, id, storeID, name, description, sku, price)
}

func (a Application) GetProduct(ctx context.Context, id string) (product *domain.Product, err error) {
	a.logger.Info().Msg("--> Stores.GetProduct")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Stores.GetProduct")
		} else {
			a.logger.Info().Msg("<-- Stores.GetProduct") 
		}
	}()
	return a.App.GetProduct(ctx, id)
}

func (a Application) GetCatalog(ctx context.Context, storeID string) (products []*domain.Product, err error) {
	a.logger.Info().Msg("--> Stores.GetCatalog")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Stores.GetCatalog")
		} else {
			a.logger.Info().Msg("<-- Stores.GetCatalog") 
		}
	}()
	return a.App.GetCatalog(ctx, storeID)
}

func (a Application) DeleteProduct(ctx context.Context, id string) (err error) {
	a.logger.Info().Msg("--> Stores.DeleteProduct")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Stores.DeleteProduct")
		} else {
			a.logger.Info().Msg("<-- Stores.DeleteProduct") 
		}
	}()
	return a.App.DeleteProduct(ctx, id)
}
