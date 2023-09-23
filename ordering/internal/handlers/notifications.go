package handlers

import (
	"github.com/hnamzian/go-mallbots/internal/ddd"
	"github.com/hnamzian/go-mallbots/ordering/internal/application"
	"github.com/hnamzian/go-mallbots/ordering/internal/domain"
)

func RegisterNotificationHandlers(handlers application.DomainEventHandlers, subscriber ddd.EventSubscriber) {
	subscriber.Subscribe(domain.OrderCreated{}, handlers.OnOrderCreated)
	subscriber.Subscribe(domain.OrderCancelled{}, handlers.OnOrderCancelled)
	subscriber.Subscribe(domain.OrderReadied{}, handlers.OnOrderReadied)
}