package presentation

import (
	"context"
	"fmt"
	"net"

	"../crossutting/memoryBus"
	"../ddd_application"
	"../ddd_infrastructure/kendoDDDProto"
	"../ddd_infrastructure/os/shutdown"
	"../dddcore"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
)

//初始化RPC服务
func InitRPC() {
	beeLog := logs.GetBeeLogger()

	grpcServer := grpc.NewServer()
	kendoDDDProto.RegisterKendoGrpcServer(grpcServer, KendoRpcSer)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "127.0.0.1", beego.AppConfig.String("rpcport")))
	if err != nil {
		beeLog.Info("failed to listen: %v\n", err)
	} else {
		beeLog.Info("[RPC] running at %s:%d\n", "127.0.0.1", beego.AppConfig.String("rpcport"))
	}

	go func() {
		beeLog.Critical("failed to serve: %v\n", grpcServer.Serve(lis))
	}()

	shutdown.GracefulStop(func() {
		beeLog.Info("[RPC] shutting down...\n")

		grpcServer.GracefulStop()

		beeLog.Info("[RPC] gracefully stopped\n")
	})

}

// KendoRpcSer 统一的Rpc服务
var KendoRpcSer kendoDDDProto.KendoGrpcServer

func init() {
	eventStore := memoryBus.NewEventStore()
	eventBus := memoryBus.NewEventBus()
	commandBus := memoryBus.NewCommandBus()

	KendoRpcSer = newGrpcServer(commandBus, eventBus, eventStore)
}

//
type GRPCServer struct {
	commandBus dddcore.CommandBus
	eventBus   dddcore.EventBus
	eventStore dddcore.EventStore
}

//
func registerRpcCommandHandlers(cb dddcore.CommandBus, es dddcore.EventStore, eb dddcore.EventBus) {
	cb.Subscribe(SingnProtocolByInfo, ddd_application.OnSingnProtocol(es, eb))
	cb.SubscribeMap(CfrmProtocolByReqSn, ddd_application.OnCfrmProtocol(es, eb))
	cb.SubscribeMap(GetProtocolInfoByNo, ddd_application.OnGetProtocolByNo(es, eb))
}

///
func registerRpcEventHandlers(eb dddcore.EventBus) {
}

// DispatchCommand implements proto.UserServer interface
func (s *GRPCServer) DispatchCommand(ctx context.Context, cmd *kendoDDDProto.CommandRequest) (*kendoDDDProto.CommandResponse, error) {
	out := make(chan dddcore.BusChan)
	defer close(out)

	go func() {
		s.commandBus.Publish(ctx, cmd.GetName(), cmd.GetPayload(), out)
	}()

	select {
	case <-ctx.Done():
		rst := new(kendoDDDProto.CommandResponse)
		rst.IsSucceed = false
		rst.ErrMsg = ctx.Err().Error()
		return rst, ctx.Err()
	case msg, _ := <-out:
		rst := new(kendoDDDProto.CommandResponse)
		rst.IsSucceed = true
		rst.Infos = msg.RespMsg
		return rst, msg.ErrMsg
	}
}

//DispatchCommandMap 处理Map类型的参数
func (s *GRPCServer) DispatchCommandMap(ctx context.Context, cmd *kendoDDDProto.CommandMapRequest) (*kendoDDDProto.CommandResponse, error) {
	out := make(chan dddcore.BusChan)
	defer close(out)

	go func() {
		s.commandBus.Publish(ctx, cmd.GetName(), cmd.GetDicInfo(), out)
	}()

	select {
	case <-ctx.Done():
		rst := new(kendoDDDProto.CommandResponse)
		rst.IsSucceed = false
		rst.ErrMsg = ctx.Err().Error()
		return rst, ctx.Err()
	case msg, _ := <-out:
		rst := new(kendoDDDProto.CommandResponse)
		rst.IsSucceed = true
		rst.Infos = msg.RespMsg
		return rst, msg.ErrMsg
	}
}

// newGrpcServer returns new user server object
func newGrpcServer(cb dddcore.CommandBus, eb dddcore.EventBus, es dddcore.EventStore) kendoDDDProto.KendoGrpcServer {
	s := &GRPCServer{cb, eb, es}

	registerRpcCommandHandlers(cb, es, eb)
	registerRpcEventHandlers(eb)

	return s
}
