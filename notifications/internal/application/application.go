package application

import (
	"context"

	"github.com/hnamzian/go-mallbots/notifications/internal/domain"
)

type App interface {
	NotifyOrderCreated(ctx context.Context, notify OrderCreated) error
	NotifyOrderCanceled(ctx context.Context, notify OrderCanceled) error
	NotifyOrderReady(ctx context.Context, notify OrderReady) error
}

type Application struct{
	customers domain.CustomerRepository
}

type (
	OrderCreated struct {
		OrderID    string
		CustomerID string
	}

	OrderCanceled struct {
		OrderID    string
		CustomerID string
	}

	OrderReady struct {
		OrderID    string
		CustomerID string
	}
)

func NewApplication(customers domain.CustomerRepository) Application {
	return Application{
		customers,
	}
}

func (a Application) NotifyOrderCreated(ctx context.Context, notify OrderCreated) error {
	// Not Implemented
	return nil
}

func (a Application) NotifyOrderCanceled(ctx context.Context, notify OrderCanceled) error {
	// Not Implemented
	return nil
}

func (a Application) NotifyOrderReady(ctx context.Context, notify OrderReady) error {
	// Not Implemented
	return nil
}
