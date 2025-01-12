package events

import (
	"errors"
	"log"
)

type EventRouter struct {
	handlers map[string]func([]byte) error
}

func NewEventRouter() *EventRouter {
	return &EventRouter{
		handlers: make(map[string]func([]byte) error),
	}
}

func (r *EventRouter) Register(eventType string, handler func([]byte) error) {
	r.handlers[eventType] = handler
}

func (r *EventRouter) Route(eventType string, payload []byte) error {
	handler, exists := r.handlers[eventType]
	if !exists {
		log.Printf("No handler registered for event type: %s", eventType)
		return errors.New("no handler for event type: " + eventType)
	}

	return handler(payload)
}