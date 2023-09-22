package application

import (
	"context"

	"github.com/hnamzian/go-mallbots/depot/internal/domain"
	"github.com/hnamzian/go-mallbots/internal/ddd"
)

type OrderEventHandlers struct {
	orders domain.OrderRepository
}

func NewOrderEventHandlers(orders domain.OrderRepository) *OrderEventHandlers {
	return &OrderEventHandlers{
		orders: orders,
	}
}

func (h *OrderEventHandlers) OnShoppingListCompleted(ctx context.Context, event ddd.Event) error {
	shopping := event.(*domain.ShoppingListCompleted)
	return h.orders.Ready(ctx, shopping.ShoppingList.ID)
}
