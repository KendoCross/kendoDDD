package presentation

import (
	"../dddcore"
)

type grpcServer struct {
	commandBus dddcore.CommandBus
	eventBus   dddcore.EventBus
	eventStore dddcore.EventStore
}

// func registerCommandHandlers(cb dddcore.CommandBus, es dddcore.EventStore, eb dddcore.EventBus) {
// 	cb.Subscribe(SingnProtocolByInfo, application.OnRegisterUserWithEmail(es, eb, j))
// }

// func registerEventHandlers(eb dddcore.EventBus) {
// }

// // DispatchCommand implements proto.UserServer interface
// func (s *grpcServer) DispatchCommand(ctx context.Context, cmd *proto.DispatchCommandRequest) (*proto.DispatchCommandResponse, error) {
// 	out := make(chan error)
// 	defer close(out)

// 	go func() {
// 		s.commandBus.Publish(ctx, cmd.GetName(), cmd.GetPayload(), out)
// 	}()

// 	select {
// 	case <-ctx.Done():
// 		return new(proto.DispatchCommandResponse), ctx.Err()
// 	case err := <-out:
// 		return new(proto.DispatchCommandResponse), err
// 	}
// }

// // New returns new user server object
// func New(cb dddcore.CommandBus, eb dddcore.EventBus, es dddcore.EventStore) proto.UserServer {
// 	s := &userServer{cb, eb, es}

// 	registerCommandHandlers(cb, es, eb)
// 	registerEventHandlers(eb)

// 	return s
// }
