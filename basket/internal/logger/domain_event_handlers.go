package logger

import (
	"context"

	"github.com/hnamzian/go-mallbots/basket/internal/application"
	"github.com/hnamzian/go-mallbots/internal/ddd"
	"github.com/rs/zerolog"
)

type DomainEventHandlers struct {
	application.DomainEventHandlers
	logger zerolog.Logger
}

func NewDomainEventHandlers(handlers application.DomainEventHandlers, logger zerolog.Logger) *DomainEventHandlers {
	return &DomainEventHandlers{
		DomainEventHandlers: handlers,
		logger:              logger,
	}
}

func (h DomainEventHandlers) OnBasketStarted(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> Basket.OnBasketStarted")
	defer func() {
		if err != nil {
			h.logger.Error().Err(err).Msg("<-- Basket.OnBasketStarted")
		} else {
			h.logger.Info().Msg("<-- Basket.OnBasketStarted")
		}
	}()
	return h.DomainEventHandlers.OnBasketStarted(ctx, event)
}

func (h DomainEventHandlers) OnBasketItemAdded(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> Basket.OnBasketItemAdded")
	defer func() {
		if err != nil {
			h.logger.Error().Err(err).Msg("<-- Basket.OnBasketItemAdded")
		} else {
			h.logger.Info().Msg("<-- Basket.OnBasketItemAdded")
		}
	}()
	return h.DomainEventHandlers.OnBasketItemAdded(ctx, event)
}

func (h DomainEventHandlers) OnBasketItemRemoved(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> Basket.OnBasketItemRemoved")
	defer func() {
		if err != nil {
			h.logger.Error().Err(err).Msg("<-- Basket.OnBasketItemRemoved")
		} else {
			h.logger.Info().Msg("<-- Basket.OnBasketItemRemoved")
		}
	}()
	return h.DomainEventHandlers.OnBasketItemRemoved(ctx, event)
}

func (h DomainEventHandlers) OnBasketCanceled(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> Basket.OnBasketCanceled")
	defer func() {
		if err != nil {
			h.logger.Error().Err(err).Msg("<-- Basket.OnBasketCanceled")
		} else {
			h.logger.Info().Msg("<-- Basket.OnBasketCanceled")
		}
	}()
	return h.DomainEventHandlers.OnBasketCanceled(ctx, event)
}

func (h DomainEventHandlers) OnBasketCheckedOut(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> Basket.OnBasketCheckedOut")
	defer func() {
		if err != nil {
			h.logger.Error().Err(err).Msg("<-- Basket.OnBasketCheckedOut")
		} else {
			h.logger.Info().Msg("<-- Basket.OnBasketCheckedOut")
		}
	}()
	return h.DomainEventHandlers.OnBasketCheckedOut(ctx, event)
}
