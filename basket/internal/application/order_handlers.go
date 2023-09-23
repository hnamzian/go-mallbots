package application

import (
	"context"

	"github.com/hnamzian/go-mallbots/basket/internal/domain"
	"github.com/hnamzian/go-mallbots/internal/ddd"
)

type OrderHandlers struct {
	orders domain.OrderRepository
	ignoreUnimplementedDomainEvents
}

var _ DomainEventHandlers = (*OrderHandlers)(nil)

func NewOrderHandlers(orders domain.OrderRepository) *OrderHandlers {
	return &OrderHandlers{
		orders: orders,
	}
}

func (h *OrderHandlers) OnBasketCheckedOut(ctx context.Context, event ddd.Event) error {
	basket := event.(*domain.BasketCheckedOut).Basket
	_, err := h.orders.Save(ctx, basket)
	return err
}
