package presentation

import (
	"context"

	"../crossutting/memoryBus"
	"../ddd_application"
	"../dddcore"
)

// Server API for HTTP
type HttpServer interface {
	DispatchCommand(context.Context, string, []byte) (string, error)
	DispatchCommandMap(context.Context, string, map[string]string) (string, error)
}

//统一的HTTP服务
var kendoHttpSer HttpServer

type httpServer struct {
	commandBus dddcore.CommandBus
	eventBus   dddcore.EventBus
	eventStore dddcore.EventStore
}

func init() {
	eventStore := memoryBus.NewEventStore()
	eventBus := memoryBus.NewEventBus()
	commandBus := memoryBus.NewCommandBus()

	kendoHttpSer = newHttpServer(commandBus, eventBus, eventStore)
}

// 中转命令
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

//中转命令
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

// newHttpServer returns new server object
func newHttpServer(cb dddcore.CommandBus, eb dddcore.EventBus, es dddcore.EventStore) HttpServer {
	s := &httpServer{cb, eb, es}

	registerHttpCommandHandlers(cb, es, eb)
	registerHttpEventHandlers(eb)

	return s
}

//registerHttpCommandHandlers注册所有命令
func registerHttpCommandHandlers(cb dddcore.CommandBus, es dddcore.EventStore, eb dddcore.EventBus) {
	cb.Subscribe(SingnProtocolByInfo, ddd_application.OnSingnProtocol(es, eb))
	cb.SubscribeMap(CfrmProtocolByReqSn, ddd_application.OnCfrmProtocol(es, eb))
	cb.SubscribeMap(GetProtocolInfoByNo, ddd_application.OnGetProtocolByNo(es, eb))
	cb.SubscribeMap(GetWSTokenByUser, ddd_application.OnGetWsToken(es, eb))
}

//registerHttpEventHandlers注册所有事件
func registerHttpEventHandlers(es dddcore.EventBus) {
}
