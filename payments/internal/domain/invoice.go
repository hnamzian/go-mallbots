package domain

import "github.com/stackus/errors"

var (
	ErrInvoiceCannotBeChanged   = errors.Wrap(errors.ErrBadRequest, "the invoice cannot be changed")
	ErrInvoiceCannotBePaid      = errors.Wrap(errors.ErrBadRequest, "the invoice cannot be paid")
	ErrInvoiceCannotBeCancelled = errors.Wrap(errors.ErrBadRequest, "the invoice cannot be cancelled")
)

type Invoice struct {
	ID      string
	OrderID string
	Amount  float64
	Status  InvoiceStatus
}

func CreateInvoice(id, orderID string, amount float64) *Invoice {
	return &Invoice{
		ID:      id,
		OrderID: orderID,
		Amount:  amount,
		Status:  InvoicePending,
	}
}

func (i *Invoice) Adjust(amount float64) error {
	if i.Status != InvoicePending {
		return ErrInvoiceCannotBeChanged
	}
	i.Amount = amount
	return nil
}

func (i *Invoice) Pay() error {
	if i.Status != InvoicePending {
		return ErrInvoiceCannotBePaid
	}
	i.Status = InvoicePaid
	return nil
}

func (i *Invoice) Cancel() error {
	if i.Status != InvoicePending {
		return ErrInvoiceCannotBeChanged
	}
	i.Status = InvoiceCancelled
	return nil
}
