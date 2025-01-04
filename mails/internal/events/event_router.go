package events

import "errors"

type EventRouter struct {
	handlers map[string]func([]byte) (err error)
}

func NewEventRouter() *EventRouter {
	return &EventRouter{
		handlers: make(map[string]func([]byte) (err error)),
	}
}

func (r *EventRouter) Register(eventType string, handler func([]byte) (err error)) {
	r.handlers[eventType] = handler
}

func (r *EventRouter) Route(eventType string, payload []byte) (err error) {
	handler, exists := r.handlers[eventType]
	if !exists {
		return errors.New("no handler for event type: " + eventType)
	}

	return handler(payload)
}