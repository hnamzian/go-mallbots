package ddd

import (
	"context"
	"sync"
)

type EventSubscriber interface {
	Subscribe(event Event, handler EventHandler)
}

type EventPublisher interface {
	Publish(ctx context.Context, events ...Event) error
}

type EventDispatcher struct {
	mu sync.RWMutex
	handlers map[string][]EventHandler
}

var _ interface {
	EventSubscriber
	EventPublisher
} = (*EventDispatcher)(nil)

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (d *EventDispatcher) Subscribe(event Event, handler EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.handlers[event.EventName()] = append(d.handlers[event.EventName()], handler)
}

func (d *EventDispatcher) Publish(ctx context.Context, events ...Event) error {
	for _, event := range events {
		for _, handler := range d.handlers[event.EventName()] {
			if err := handler(ctx, event); err != nil {
				return err
			}
		}
	}

	return nil
}
