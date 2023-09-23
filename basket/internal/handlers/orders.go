package handlers

import (
	"github.com/hnamzian/go-mallbots/basket/internal/application"
	"github.com/hnamzian/go-mallbots/basket/internal/domain"
	"github.com/hnamzian/go-mallbots/internal/ddd"
)

func RegisterOrderHandlers(handlers application.DomainEventHandlers, subscriber ddd.EventSubscriber) {
	subscriber.Subscribe(&domain.BasketCheckedOut{}, handlers.OnBasketCheckedOut)
}
