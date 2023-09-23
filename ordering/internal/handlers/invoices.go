package handlers

import (
	"github.com/hnamzian/go-mallbots/internal/ddd"
	"github.com/hnamzian/go-mallbots/ordering/internal/application"
	"github.com/hnamzian/go-mallbots/ordering/internal/domain"
)

func RegisterInvoiceHandlers(invoiceHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderReadied{}, invoiceHandlers.OnOrderReadied)
}
