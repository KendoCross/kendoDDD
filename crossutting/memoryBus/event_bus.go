package memoryBus

import (
	"context"

	dddcore "../../dddcore"

	"github.com/vardius/golog"
	messagebus "github.com/vardius/message-bus"
)

type eventBus struct {
	messageBus messagebus.MessageBus
}

func (bus *eventBus) Publish(ctx context.Context, eventType string, event dddcore.Event) {
	bus.messageBus.Publish(eventType, ctx, event)
}

func (bus *eventBus) Subscribe(eventType string, fn dddcore.EventHandler) error {
	return bus.messageBus.Subscribe(eventType, fn)
}

func (bus *eventBus) Unsubscribe(eventType string, fn dddcore.EventHandler) error {
	return bus.messageBus.Unsubscribe(eventType, fn)
}

func NewEventBus() dddcore.EventBus {
	return &eventBus{messagebus.New()}
}

type loggableEventBus struct {
	serverName string
	eventBus   dddcore.EventBus
	logger     golog.Logger
}

func (bus *loggableEventBus) Publish(ctx context.Context, eventType string, event dddcore.Event) {
	bus.logger.Debug(ctx, "[%s EventBus|Publish]: %s %s\n", bus.serverName, eventType, event.Payload)
	bus.eventBus.Publish(ctx, eventType, event)
}

func (bus *loggableEventBus) Subscribe(eventType string, fn dddcore.EventHandler) error {
	bus.logger.Info(nil, "[%s EventBus|Subscribe]: %s\n", bus.serverName, eventType)
	return bus.eventBus.Subscribe(eventType, fn)
}

func (bus *loggableEventBus) Unsubscribe(eventType string, fn dddcore.EventHandler) error {
	bus.logger.Info(nil, "[%s EventBus|Unsubscribe]: %s\n", bus.serverName, eventType)
	return bus.eventBus.Unsubscribe(eventType, fn)
}

func LogEventBus(serverName string, parent dddcore.EventBus, log golog.Logger) dddcore.EventBus {
	return &loggableEventBus{serverName, parent, log}
}
