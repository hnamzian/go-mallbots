package handlers

import (
	"github.com/hnamzian/go-mallbots/depot/internal/application"
	"github.com/hnamzian/go-mallbots/depot/internal/domain"
	"github.com/hnamzian/go-mallbots/internal/ddd"
)

func RegisterOrderHandlers(domainSubscribe ddd.EventSubscriber, handlers application.OrderEventHandlers) {
	domainSubscribe.Subscribe(&domain.ShoppingListCompleted{}, handlers.OnShoppingListCompleted)
}
