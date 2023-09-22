package logger

import (
	"context"

	"github.com/hnamzian/go-mallbots/internal/ddd"
	"github.com/hnamzian/go-mallbots/ordering/internal/application"
	"github.com/rs/zerolog"
)

type DomainEventHandlers struct {
	application.DomainEventHandlers
	logger zerolog.Logger
}

func NewDomainEventHandlers(domainHandlers application.DomainEventHandlers, logger zerolog.Logger) *DomainEventHandlers {
	return &DomainEventHandlers{
		DomainEventHandlers: domainHandlers,
		logger:              logger,
	}
}

func (h DomainEventHandlers) OnOrderCreated(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> Ordering.OnOrderCreated")
	defer func() {
		if err != nil {
			h.logger.Error().Err(err).Msg("<-- Ordering.OnOrderCreated")
		} else {
			h.logger.Info().Msg("<-- Ordering.OnOrderCreated")
		}
	}()
	return h.DomainEventHandlers.OnOrderCreated(ctx, event)
}

func (h DomainEventHandlers) OnOrderCancelled(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> Ordering.OnOrderCreated")
	defer func() {
		if err != nil {
			h.logger.Error().Err(err).Msg("<-- Ordering.OnOrderCreated")
		} else {
			h.logger.Info().Msg("<-- Ordering.OnOrderCreated")
		}
	}()
	return h.DomainEventHandlers.OnOrderCreated(ctx, event)
}

func (h DomainEventHandlers) OnOrderReadied(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> Ordering.OnOrderCreated")
	defer func() {
		if err != nil {
			h.logger.Error().Err(err).Msg("<-- Ordering.OnOrderCreated")
		} else {
			h.logger.Info().Msg("<-- Ordering.OnOrderCreated")
		}
	}()
	return h.DomainEventHandlers.OnOrderCreated(ctx, event)
}

func (h DomainEventHandlers) OnOrderCompleted(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> Ordering.OnOrderCreated")
	defer func() {
		if err != nil {
			h.logger.Error().Err(err).Msg("<-- Ordering.OnOrderCreated")
		} else {
			h.logger.Info().Msg("<-- Ordering.OnOrderCreated")
		}
	}()
	return h.DomainEventHandlers.OnOrderCreated(ctx, event)
}
