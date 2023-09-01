package logger

import (
	"context"

	"github.com/hnamzian/go-mallbots/payments/internal/application"
	"github.com/hnamzian/go-mallbots/payments/internal/domain"
	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

func NewApplication(logger zerolog.Logger, app application.App) *Application {
	return &Application{App: app, logger: logger}
}

func (a Application) AuthorizePayment(ctx context.Context, authorize application.AuthorizePayment) (err error) {
	a.logger.Info().Msgf("Payments.AuthorizePayment: %v", authorize)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Payments.AuthorizePayment")
		} else {
			a.logger.Info().Msg("<-- Payments.AuthorizePayment")
		}
	}()
	return a.App.AuthorizePayment(ctx, authorize)
}

func (a Application) ConfirmPayment(ctx context.Context, confirm application.ConfirmPayment) (err error) {
	a.logger.Info().Msgf("Payments.ConfirmPayment: %v", confirm)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Payments.ConfirmPayment")
		} else {
			a.logger.Info().Msg("<-- Payments.ConfirmPayment")
		}
	}()
	return a.App.ConfirmPayment(ctx, confirm)
}

func (a Application) GetPayment(ctx context.Context, get application.GetPayment) (payment *domain.Payment, err error) {
	a.logger.Info().Msgf("Payments.GetPayment: %v", get)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Payments.GetPayment")
		} else {
			a.logger.Info().Msg("<-- Payments.GetPayment")
		}
	}()
	return a.App.GetPayment(ctx, get)
}

func (a Application) CreateInvoice(ctx context.Context, create application.CreateInvoice) (err error) {
	a.logger.Info().Msgf("Payments.CreateInvoice: %v", create)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Payments.CreateInvoice")
		} else {
			a.logger.Info().Msg("<-- Payments.CreateInvoice")
		}
	}()
	return a.App.CreateInvoice(ctx, create)
}

func (a Application) AdjustInvoice(ctx context.Context, adjust application.AdjustInvoice) (err error) {
	a.logger.Info().Msgf("Payments.AdjustInvoice: %v", adjust)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Payments.AdjustInvoice")
		} else {
			a.logger.Info().Msg("<-- Payments.AdjustInvoice")
		}
	}()
	return a.App.AdjustInvoice(ctx, adjust)
}

func (a Application) PayInvoice(ctx context.Context, pay application.PayInvoice) (err error) {
	a.logger.Info().Msgf("Payments.PayInvoice: %v", pay)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Payments.PayInvoice")
		} else {
			a.logger.Info().Msg("<-- Payments.PayInvoice")
		}
	}()
	return a.App.PayInvoice(ctx, pay)
}

func (a Application) CancelInvoice(ctx context.Context, cancel application.CancelInvoice) (err error) {
	a.logger.Info().Msgf("Payments.CancelInvoice: %v", cancel)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Payments.CancelInvoice")
		} else {
			a.logger.Info().Msg("<-- Payments.CancelInvoice")
		}
	}()
	return a.App.CancelInvoice(ctx, cancel)
}

func (a Application) GetInvoice(ctx context.Context, get application.GetInvoice) (invoice *domain.Invoice, err error) {
	a.logger.Info().Msgf("Payments.GetInvoice: %v", get)
	defer func() {
		if err != nil {
			a.logger.Error().Err(err).Msg("<-- Payments.GetInvoice")
		} else {
			a.logger.Info().Msg("<-- Payments.GetInvoice")
		}
	}()
	return a.App.GetInvoice(ctx, get)
}
