package application

import (
	"context"

	"github.com/hnamzian/go-mallbots/internal/ddd"
)

type DomainEventHandlers interface {
	OnOrderCreated(context.Context, ddd.Event) error
	OnOrderCancelled(context.Context, ddd.Event) error
	OnOrderReadied(context.Context, ddd.Event) error
	OnOrderCompleted(context.Context, ddd.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (ignoreUnimplementedDomainEvents) OnOrderCreated(context.Context, ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnOrderCancelled(context.Context, ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnOrderReadied(context.Context, ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnOrderCompleted(context.Context, ddd.Event) error {
	return nil
}
