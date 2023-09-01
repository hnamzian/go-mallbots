package application

import (
	"context"

	"github.com/hnamzian/go-mallbots/payments/internal/domain"
	"github.com/stackus/errors"
)

type App interface {
	AuthorizePayment(ctx context.Context, authorize AuthorizePayment) error
	ConfirmPayment(ctx context.Context, confirm ConfirmPayment) error
	GetPayment(ctx context.Context, get GetPayment) (*domain.Payment, error)
	CreateInvoice(ctx context.Context, create CreateInvoice) error
	AdjustInvoice(ctx context.Context, adjust AdjustInvoice) error
	PayInvoice(ctx context.Context, pay PayInvoice) error
	CancelInvoice(ctx context.Context, cancel CancelInvoice) error
	GetInvoice(ctx context.Context, get GetInvoice) (*domain.Invoice, error)
}

type Application struct {
	invoices domain.InvoiceRepository
	payments domain.PaymentRepository
	orders   domain.OrderRepository
}

type (
	AuthorizePayment struct {
		ID         string
		CustomerID string
		Amount     float64
	}

	ConfirmPayment struct {
		ID string
	}

	GetPayment struct {
		ID string
	}

	CreateInvoice struct {
		ID      string
		OrderID string
		Amount  float64
	}

	AdjustInvoice struct {
		ID     string
		Amount float64
	}

	PayInvoice struct {
		ID string
	}

	CancelInvoice struct {
		ID string
	}

	GetInvoice struct {
		ID string
	}
)

func NewApplication(invoices domain.InvoiceRepository, payments domain.PaymentRepository, orders domain.OrderRepository) *Application {
	return &Application{
		invoices: invoices,
		payments: payments,
		orders:   orders,
	}
}

func (a Application) AuthorizePayment(ctx context.Context, authorize AuthorizePayment) error {
	return a.payments.Save(ctx, &domain.Payment{
		ID:         authorize.ID,
		CustomerID: authorize.CustomerID,
		Amount:     authorize.Amount,
	})
}

func (a Application) ConfirmPayment(ctx context.Context, confirm ConfirmPayment) error {
	payment, err := a.payments.Find(ctx, confirm.ID)
	if err != nil || payment == nil {
		return errors.Wrap(err, "payment not found")
	}
	return nil
}

func (a Application) GetPayment(ctx context.Context, get GetPayment) (*domain.Payment, error) {
	return a.payments.Find(ctx, get.ID)
}

func (a Application) CreateInvoice(ctx context.Context, create CreateInvoice) error {
	invoice := domain.CreateInvoice(create.ID, create.OrderID, create.Amount)
	return a.invoices.Save(ctx, invoice)
}

func (a Application) AdjustInvoice(ctx context.Context, adjust AdjustInvoice) error {
	invoice, err := a.invoices.Find(ctx, adjust.ID)
	if err != nil || invoice == nil {
		return errors.Wrap(err, "invoice not found")
	}

	err = invoice.Adjust(adjust.Amount)
	if err != nil {
		return err
	}

	return a.invoices.Update(ctx, invoice)
}

func (a Application) PayInvoice(ctx context.Context, pay PayInvoice) error {
	invoice, err := a.invoices.Find(ctx, pay.ID)
	if err != nil || invoice == nil {
		return errors.Wrap(err, "invoice not found")
	}

	if err = invoice.Pay(); err != nil {
		return err
	}

	return a.invoices.Update(ctx, invoice)	
}

func (a Application) CancelInvoice(ctx context.Context, cancel CancelInvoice) error {
	invoice, err := a.invoices.Find(ctx, cancel.ID)
	if err != nil || invoice == nil {
		return errors.Wrap(err, "invoice not found")
	}

	if err = invoice.Cancel(); err != nil {
		return err
	}

	return a.invoices.Update(ctx, invoice)
}

func (a Application) GetInvoice(ctx context.Context, get GetInvoice) (*domain.Invoice, error) {
	return a.invoices.Find(ctx, get.ID)
}
