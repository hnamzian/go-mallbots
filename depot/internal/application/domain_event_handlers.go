package application

import (
	"context"

	"github.com/hnamzian/go-mallbots/internal/ddd"
)

type DomainEventHandlers interface {
	OnShoppingListCreated(context.Context, ddd.Event) error
	OnShoppingListCancelled(context.Context, ddd.Event) error
	OnShoppingListAssigned(context.Context, ddd.Event) error
	OnShoppingListCompleted(context.Context, ddd.Event) error
}

type ignoreUnimplementedDomainHandlers struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainHandlers)(nil)

func (ignoreUnimplementedDomainHandlers) OnShoppingListCreated(context.Context, ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainHandlers) OnShoppingListCancelled(context.Context, ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainHandlers) OnShoppingListAssigned(context.Context, ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainHandlers) OnShoppingListCompleted(context.Context, ddd.Event) error {
	return nil
}
