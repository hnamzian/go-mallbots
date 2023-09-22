package domain

import "github.com/stackus/errors"

var (
	ErrOrderHasNoItems         = errors.Wrap(errors.ErrBadRequest, "the order has no items")
	ErrOrderCannotBeCancelled  = errors.Wrap(errors.ErrBadRequest, "the order cannot be cancelled")
	ErrCustomerIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
	ErrPaymentIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
)

type Order struct {
	ID         string
	CustomerID string
	PaymentID  string
	InvoiceID  string
	ShoppingID string
	Items      []Item
	Status     OrderStatus
}

func CreateOrder(id, customerID, paymentID string, items []Item) (*Order, error) {
	if len(items) == 0 {
		return nil, ErrOrderHasNoItems
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	if paymentID == "" {
		return nil, ErrPaymentIDCannotBeBlank
	}

	order := &Order{
		ID:         id,
		CustomerID: customerID,
		PaymentID:  paymentID,
		Items:      items,
		Status:     OrderIsPending,
	}

	return order, nil
}

func (o *Order) Cancel() error {
	if o.Status != OrderIsPending {
		return ErrOrderCannotBeCancelled
	}
	o.Status = OrderIsCancelled
	return nil
}

func (o *Order) Complete(invoiceID string) error {
	o.InvoiceID = invoiceID
	o.Status = OrderIsCompleted
	return nil
}

func (o *Order) Ready() error {
	o.Status = OrderIsReady
	return nil
}

func (o *Order) GetTotal() float64 {
	total := float64(0)
	for _, item := range(o.Items) {
		total += item.Price * float64(item.Quantity) 
	}
	
	return total
}