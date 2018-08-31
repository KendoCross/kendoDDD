package presentation

import (
	"context"

	"../crossutting/memoryBus"
	"../ddd_application"
	"../dddcore"
)

// Server API for HTTP
// 订阅中心
type HttpServer interface {
	DispatchCommand(context.Context, string, []byte) (string, error)
	DispatchCommandMap(context.Context, string, map[string]string) (string, error)
}

var kendoSer HttpServer

type httpServer struct {
	commandBus dddcore.CommandBus
	eventBus   dddcore.EventBus
	eventStore dddcore.EventStore
}

func init() {
	eventStore := memoryBus.NewEventStore()
	eventBus := memoryBus.NewEventBus()
	commandBus := memoryBus.NewCommandBus()

	kendoSer = new(commandBus, eventBus, eventStore)
}

// DispatchCommand implements proto.UserServer interface
func (s *httpServer) DispatchCommand(ctx context.Context, command string, payload []byte) (string, error) {
	out := make(chan dddcore.BusChan)
	defer close(out)

	go func() {
		s.commandBus.Publish(ctx, command, payload, out)
	}()

	select {
	case <-ctx.Done():
		return "ctx.Done了", ctx.Err()
	case msg, _ := <-out:
		return msg.RespMsg, msg.ErrMsg
	}
}

func (s *httpServer) DispatchCommandMap(ctx context.Context, command string, payload map[string]string) (string, error) {
	out := make(chan dddcore.BusChan)
	defer close(out)

	go func() {
		s.commandBus.Publish(ctx, command, payload, out)
	}()

	select {
	case <-ctx.Done():
		return "ctx.Done了", ctx.Err()
	case msg, _ := <-out:
		return msg.RespMsg, msg.ErrMsg
	}
}

// New returns new user server object
func new(cb dddcore.CommandBus, eb dddcore.EventBus, es dddcore.EventStore) HttpServer {
	s := &httpServer{cb, eb, es}

	registerCommandHandlers(cb, es, eb)
	registerEventHandlers(eb)

	return s
}

func registerCommandHandlers(cb dddcore.CommandBus, es dddcore.EventStore, eb dddcore.EventBus) {
	cb.Subscribe(SingnProtocolByInfo, ddd_application.OnSingnProtocol(es, eb))
	cb.SubscribeMap(CfrmProtocolByReqSn, ddd_application.OnCfrmProtocol(es, eb))
	cb.SubscribeMap(GetProtocolInfoByNo, ddd_application.OnGetProtocolByNo(es, eb))
}

func registerEventHandlers(es dddcore.EventBus) {
}
