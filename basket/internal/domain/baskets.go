package domain

import (
	"sort"

	"github.com/stackus/errors"
)

type BasketStatus string

var (
	ErrBasketHasNoItems         = errors.Wrap(errors.ErrBadRequest, "the basket has no items")
	ErrBasketCannotBeModified   = errors.Wrap(errors.ErrBadRequest, "the basket cannot be modified")
	ErrBasketCannotBeCancelled  = errors.Wrap(errors.ErrBadRequest, "the basket cannot be cancelled")
	ErrQuantityCannotBeNegative = errors.Wrap(errors.ErrBadRequest, "the item quantity cannot be negative")
	ErrBasketIDCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the basket id cannot be blank")
	ErrPaymentIDCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
	ErrCustomerIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
)

const (
	BasketUnknown    BasketStatus = "Unknown"
	BasketOpen       BasketStatus = "Open"
	BasketCancelled  BasketStatus = "Cancelled"
	BasketCheckedOut BasketStatus = "CheckedOut"
)

type Basket struct {
	ID         string
	CustomerID string
	PaymentID  string
	Items      []Item
	Status     BasketStatus
}

func StartBasket(id, customerID string) *Basket {
	return &Basket{
		ID:         id,
		CustomerID: customerID,
		Status:     BasketOpen,
	}
}

func (b *Basket) IsOpen() bool {
	return b.Status == BasketOpen
}

func (b *Basket) IsCancellable() bool {
	return b.Status == BasketOpen
}

func (b *Basket) Cancel() error {
	if b.IsCancellable() {
		b.Status = BasketCancelled
		return nil
	}
	return ErrBasketCannotBeCancelled
}

func (b *Basket) Checkout(paymentID string) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if len(b.Items) == 0 {
		return ErrBasketHasNoItems
	}

	if paymentID == "" {
		return ErrPaymentIDCannotBeBlank
	}

	b.PaymentID = paymentID
	b.Status = BasketCheckedOut

	return nil
}

func (b *Basket) AddItem(store *Store, product *Product, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	for i, item := range(b.Items) {
		if item.StoreID == store.ID && item.ProductID == product.ID {
			b.Items[i].Quantity += quantity
			return nil
		}
	}

	b.Items = append(b.Items, Item{
		StoreID:      store.ID,
		ProductID:    product.ID,
		StoreName:    store.Name,
		ProductName:  product.Name,
		ProductPrice: product.Price,
		Quantity:     quantity,
	})

	sort.Slice(b.Items, func(i, j int) bool {
		return b.Items[i].ProductName < b.Items[j].ProductName && b.Items[i].StoreName <= b.Items[j].StoreName
	})

	return nil
}

func (b *Basket) RemoveItem(product *Product, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	for i, item := range(b.Items) {
		if item.ProductID == product.ID && item.StoreID == product.StoreID {
			b.Items[i].Quantity -= quantity

			if b.Items[i].Quantity < 1 {
				b.Items = append(b.Items[:i], b.Items[i+1:]...)
			}
			return nil
		}
	}

	return nil
}
