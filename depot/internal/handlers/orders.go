package handlers

import (
	"github.com/hnamzian/go-mallbots/depot/internal/application"
	"github.com/hnamzian/go-mallbots/depot/internal/domain"
	"github.com/hnamzian/go-mallbots/internal/ddd"
)

func RegisterOrderHandlers(handlers application.DomainEventHandlers, domainSubscribe ddd.EventSubscriber) {
	domainSubscribe.Subscribe(&domain.ShoppingListCompleted{}, handlers.OnShoppingListCompleted)
}
