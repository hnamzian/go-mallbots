package logger

import (
	"context"

	"github.com/hnamzian/go-mallbots/customers/internal/application"
	"github.com/hnamzian/go-mallbots/customers/internal/domain"
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

func (a Application) RegisterCustomer(ctx context.Context, register application.RegisterCustomer) (err error) {
	a.logger.Info().Msg("--> Customers.RegisterCustomer")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Customers.RegisterCustomer")
		} else {
			a.logger.Info().Msg("<-- Customers.RegisterCustomer") 
		}
	}()
	return a.App.RegisterCustomer(ctx, register)
}

func (a Application) AuthorizeCustomer(ctx context.Context, authorize application.AuthorizeCustomer) (err error) {
	a.logger.Info().Msg("--> Customers.AuthorizeCustomer")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Customers.AuthorizeCustomer")
		} else {
			a.logger.Info().Msg("<-- Customers.AuthorizeCustomer") 
		}
	}()
	return a.App.AuthorizeCustomer(ctx, authorize)
}

func (a Application) GetCustomer(ctx context.Context, get application.GetCustomer) (customer *domain.Customer, err error) {
	a.logger.Info().Msg("--> Customers.GetCustomer")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Customers.GetCustomer")
		} else {
			a.logger.Info().Msg("<-- Customers.GetCustomer") 
		}
	}()
	return a.App.GetCustomer(ctx, get)
}

func (a Application) EnableCustomer(ctx context.Context, enable application.EnableCustomer) (err error) {
	a.logger.Info().Msg("--> Customers.EnableCustomer")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Customers.EnableCustomer")
		} else {
			a.logger.Info().Msg("<-- Customers.EnableCustomer") 
		}
	}()
	return a.App.EnableCustomer(ctx, enable)
}

func (a Application) DisableCustomer(ctx context.Context, disable application.DisableCustomer) (err error) {
	a.logger.Info().Msg("--> Customers.DisableCustomer")
	defer func() { 
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Customers.DisableCustomer")
		} else {
			a.logger.Info().Msg("<-- Customers.DisableCustomer") 
		}
	}()
	return a.App.DisableCustomer(ctx, disable)
}
