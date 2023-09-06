package application

import (
	"context"

	"github.com/hnamzian/go-mallbots/ordering/internal/domain"
	"github.com/stackus/errors"
)

type App interface {
	CreateOrder(ctx context.Context, create CreateOrder) error
	GetOrder(ctx context.Context, get GetOrder) (*domain.Order, error)
	CancelOrder(ctx context.Context, cancel CancelOrder) error
	ReadyOrder(ctx context.Context, ready ReadyOrder) error
	CompletedOrder(ctx context.Context, complete CompletedOrder) error
}

type Application struct {
	orders        domain.OrderRepository
	customers     domain.CustomerRepository
	invoices      domain.InvoiceRepository
	shoppings     domain.ShoppingRepository
	payments      domain.PaymentRepository
	notifications domain.NotificationRepository
}

type (
	CreateOrder struct {
		ID         string
		CustomerID string
		PaymentID  string
		Items      []domain.Item
	}

	GetOrder struct {
		ID string
	}

	CancelOrder struct {
		ID string
	}

	ReadyOrder struct {
		ID string
	}

	CompletedOrder struct {
		ID        string
		InvoiceID string
	}
)

func NewOrderingApplication(orders domain.OrderRepository,
	customers domain.CustomerRepository,
	invoices domain.InvoiceRepository,
	shoppings domain.ShoppingRepository,
	payments domain.PaymentRepository,
	notifications domain.NotificationRepository) *Application {
	return &Application{
		orders,
		customers,
		invoices,
		shoppings,
		payments,
		notifications,
	}
}

func (a *Application) CreateOrder(ctx context.Context, create CreateOrder) error {
	order, err := domain.CreateOrder(create.ID, create.CustomerID, create.PaymentID, create.Items)
	if err != nil {
		return err
	}

	// Authorize Customer
	if err = a.customers.Authorize(ctx, order.CustomerID); err != nil {
		return errors.Wrap(err, "authorizing customer")
	}

	// Confirm Payment
	if err = a.payments.Confirm(ctx, order.PaymentID); err != nil {
		return errors.Wrap(err, "confirming payment")
	}

	// Schedule Shopping
	if order.ShoppingID, err = a.shoppings.Create(ctx, order); err != nil {
		return errors.Wrap(err, "creating shopping")
	}

	// Notfiy Order Created
	if err = a.notifications.NotifyOrderCreated(ctx, order.ID, order.CustomerID); err != nil {
		return errors.Wrap(err, "notifying order created")
	}

	err = a.orders.Save(ctx, order)
	if err != nil {
		return errors.Wrap(err, "saving order")
	}

	return nil
}

func (a *Application) GetOrder(ctx context.Context, get GetOrder) (*domain.Order, error) {
	return a.orders.Find(ctx, get.ID)
}

func (a *Application) CancelOrder(ctx context.Context, cancel CancelOrder) error {
	order, err := a.orders.Find(ctx, cancel.ID)
	if err != nil {
		return errors.Wrap(err, "finding order")
	}

	if err = order.Cancel(); err != nil {
		return errors.Wrap(err, "cancelling order")
	}

	if err = a.shoppings.Cancel(ctx, order.ShoppingID); err != nil {
		return errors.Wrap(err, "cancelling shopping")
	}

	if err = a.orders.Update(ctx, order); err != nil {
		return errors.Wrap(err, "updating order")
	}

	// Notify Order Cancelled
	if err = a.notifications.NotifyOrderCanceled(ctx, order.ID, order.CustomerID); err != nil {
		return errors.Wrap(err, "notifying order cancelled")
	}

	return nil
}

func (a *Application) ReadyOrder(ctx context.Context, ready ReadyOrder) error {
	order, err := a.orders.Find(ctx, ready.ID)
	if err != nil {
		return errors.Wrap(err, "finding order")
	}

	if err = order.Ready(); err != nil {
		return errors.Wrap(err, "readying order")
	}

	if err = a.orders.Update(ctx, order); err != nil {
		return errors.Wrap(err, "updating order")
	}

	// Notify Order Ready
	if err = a.notifications.NotifyOrderReady(ctx, order.ID, order.CustomerID); err != nil {
		return errors.Wrap(err, "notifying order ready")
	}

	return nil
}

func (a *Application) CompletedOrder(ctx context.Context, complete CompletedOrder) error {
	order, err := a.orders.Find(ctx, complete.ID)
	if err != nil {
		return errors.Wrap(err, "finding order")
	}

	if err = order.Complete(complete.InvoiceID); err != nil {
		return errors.Wrap(err, "completing order")
	}

	if err = a.orders.Update(ctx, order); err != nil {
		return errors.Wrap(err, "updating order")
	}

	return nil
}
