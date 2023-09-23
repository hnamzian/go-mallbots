package logger

import (
	"context"

	"github.com/hnamzian/go-mallbots/depot/internal/application"
	"github.com/hnamzian/go-mallbots/internal/ddd"
	"github.com/rs/zerolog"
)

type DomainEventHandlers struct {
	application.DomainEventHandlers
	logger zerolog.Logger
}

func NewDomainEventHandlers(handlers application.DomainEventHandlers, logger zerolog.Logger) DomainEventHandlers {
	return DomainEventHandlers{
		DomainEventHandlers: handlers,
		logger:              logger,
	}
}

func (h DomainEventHandlers) OnShoppingListCreated(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> depot.OnShoppingListCreated")
	defer func() {
		if err != nil {
			h.logger.Error().Err(err).Msg("<-- depot.OnShoppingListCreated")
		} else {
			h.logger.Info().Msg("<-- depot.OnShoppingListCreated")
		}
	}()
	return h.DomainEventHandlers.OnShoppingListCreated(ctx, event)
}

func (h DomainEventHandlers) OnShoppingListCancelled(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> depot.OnShoppingListCancelled")
	defer func() {
		if err != nil {
			h.logger.Error().Err(err).Msg("<-- depot.OnShoppingListCancelled")
		} else {
			h.logger.Info().Msg("<-- depot.OnShoppingListCancelled")
		}
	}()
	return h.DomainEventHandlers.OnShoppingListCancelled(ctx, event)
}

func (h DomainEventHandlers) OnShoppingListAssigned(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> depot.OnShoppingListAssigned")
	defer func() {
		if err != nil {
			h.logger.Error().Err(err).Msg("<-- depot.OnShoppingListAssigned")
		} else {
			h.logger.Info().Msg("<-- depot.OnShoppingListAssigned")
		}
	}()
	return h.DomainEventHandlers.OnShoppingListAssigned(ctx, event)
}

func (h DomainEventHandlers) OnShoppingListCompleted(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> depot.OnShoppingListCompleted")
	defer func() {
		if err != nil {
			h.logger.Error().Err(err).Msg("<-- depot.OnShoppingListCompleted")
		} else {
			h.logger.Info().Msg("<-- depot.OnShoppingListCompleted")
		}
	}()
	return h.DomainEventHandlers.OnShoppingListCompleted(ctx, event)
}