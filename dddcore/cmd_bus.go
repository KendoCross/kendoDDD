package dddcore

import (
	"context"
)

type BusChan struct {
	RespMsg string
	ErrMsg  error
}

// CommandHandler function
type CommandHandler func(ctx context.Context, payload []byte, out chan<- BusChan)
type CommandHandlerMap func(ctx context.Context, payload map[string]string, out chan<- BusChan)

// CommandBus allows to subscribe/dispatch commands
type CommandBus interface {
	// 发布
	Publish(ctx context.Context, command string, payload interface{}, out chan<- BusChan)
	// 订阅
	Subscribe(command string, fn CommandHandler) error
	SubscribeMap(command string, fn CommandHandlerMap) error
	// 取消
	Unsubscribe(command string, fn CommandHandler) error
	UnsubscribeMap(command string, fn CommandHandlerMap) error
}
