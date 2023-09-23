package application

import (
	"context"

	"github.com/hnamzian/go-mallbots/internal/ddd"
	"github.com/hnamzian/go-mallbots/ordering/internal/domain"
)

type NotificationHandlers struct {
	notifications domain.NotificationRepository
	ignoreUnimplementedDomainEvents
}

func NewNotificationHandlers(notifications domain.NotificationRepository) *NotificationHandlers {
	return &NotificationHandlers{
		notifications: notifications,
	}
}

func (h NotificationHandlers) OnOrderCreated(ctx context.Context, event ddd.Event) error {
	order := event.(*domain.OrderCreated)
	return h.notifications.NotifyOrderCreated(ctx, order.Order.ID, order.Order.CustomerID)
}

func (h NotificationHandlers) OnOrderCancelled(ctx context.Context, event ddd.Event) error {
	order := event.(*domain.OrderCancelled)
	return h.notifications.NotifyOrderCanceled(ctx, order.Order.ID, order.Order.CustomerID)
}

func (h NotificationHandlers) OnOrderReadied(ctx context.Context, event ddd.Event) error {
	order := event.(*domain.OrderReadied)
	return h.notifications.NotifyOrderReady(ctx, order.Order.ID, order.Order.CustomerID)
}
