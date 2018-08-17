package memoryBus

import (
	"context"

	dddcore "../../dddcore"

	"github.com/vardius/golog"
	messagebus "github.com/vardius/message-bus"
)

type commandBus struct {
	messageBus messagebus.MessageBus
}

/******************************/
//内存里（进程内）总线BUS

func (bus *commandBus) Publish(ctx context.Context, command string, payload interface{}, out chan<- dddcore.BusChan) {
	bus.messageBus.Publish(command, ctx, payload, out)
}
func (bus *commandBus) Subscribe(command string, fn dddcore.CommandHandler) error {
	return bus.messageBus.Subscribe(command, fn)
}
func (bus *commandBus) Unsubscribe(command string, fn dddcore.CommandHandler) error {
	return bus.messageBus.Unsubscribe(command, fn)
}

func (bus *commandBus) SubscribeMap(command string, fn dddcore.CommandHandlerMap) error {
	return bus.messageBus.Subscribe(command, fn)
}
func (bus *commandBus) UnsubscribeMap(command string, fn dddcore.CommandHandlerMap) error {
	return bus.messageBus.Unsubscribe(command, fn)
}

/******************************/

// New creates in memory command bus
func NewCommandBus() dddcore.CommandBus {
	return &commandBus{messagebus.New()}
}

type loggableCommandBus struct {
	serverName string
	commandBus dddcore.CommandBus
	logger     golog.Logger
}

func (bus *loggableCommandBus) Publish(ctx context.Context, command string, payload interface{}, out chan<- dddcore.BusChan) {
	bus.logger.Debug(ctx, "[%s CommandBus|Publish]: %s %s\n", bus.serverName, command, payload)
	bus.commandBus.Publish(ctx, command, payload, out)
}

func (bus *loggableCommandBus) Subscribe(command string, fn dddcore.CommandHandler) error {
	bus.logger.Info(nil, "[%s CommandBus|Subscribe]: %s\n", bus.serverName, command)
	return bus.commandBus.Subscribe(command, fn)
}

func (bus *loggableCommandBus) Unsubscribe(command string, fn dddcore.CommandHandler) error {
	bus.logger.Info(nil, "[%s CommandBus|Unsubscribe]: %s\n", bus.serverName, command)
	return bus.commandBus.Unsubscribe(command, fn)
}

func (bus *loggableCommandBus) SubscribeMap(command string, fn dddcore.CommandHandlerMap) error {
	bus.logger.Info(nil, "[%s CommandBus|Subscribe]: %s\n", bus.serverName, command)
	return bus.commandBus.SubscribeMap(command, fn)
}
func (bus *loggableCommandBus) UnsubscribeMap(command string, fn dddcore.CommandHandlerMap) error {
	bus.logger.Info(nil, "[%s CommandBus|Unsubscribe]: %s\n", bus.serverName, command)
	return bus.commandBus.UnsubscribeMap(command, fn)
}

// WithLogger creates loggable in memory command bus
func LogCommandBus(serverName string, parent dddcore.CommandBus, log golog.Logger) dddcore.CommandBus {
	return &loggableCommandBus{serverName, parent, log}
}
