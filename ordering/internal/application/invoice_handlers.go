package application

import (
	"context"

	"github.com/hnamzian/go-mallbots/internal/ddd"
	"github.com/hnamzian/go-mallbots/ordering/internal/domain"
)

type InvoiceHandlers struct {
	invoices domain.InvoiceRepository
	ignoreUnimplementedDomainEvents
}

func NewInvoiceHandlers(invoices domain.InvoiceRepository) *InvoiceHandlers {
	return &InvoiceHandlers{
		invoices: invoices,
	}
}

func (h InvoiceHandlers) OnOrderReadied(ctx context.Context, event ddd.Event) error {
	order := event.(*domain.OrderReadied)
	return h.invoices.Save(ctx, order.Order.ID, order.Order.PaymentID, order.Order.GetTotal())
}
