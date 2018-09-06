package routers

import (
	"context"
	"fmt"
	"net"

	"../ddd_infrastructure/kendoDDDProto"
	"../presentation"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
)

//服务初始化，包括HTTP和RPC服务两种
func init() {
	initRPC()
	beego.Router("/", &presentation.MainController{})
	beego.Router("/Protocol/Sign", &presentation.ProtocolController{}, "Post:Sign")
	beego.Router("/Protocol/SignCfrm", &presentation.ProtocolController{}, "*:SignCfrm")
	beego.Router("/Protocol/Sign", &presentation.ProtocolController{}, "Get:ProtocolInfo")
}

//初始化RPC服务
func initRPC() {
	ctx := context.Background()
	beeLog := logs.GetBeeLogger()

	grpcServer := grpc.NewServer()
	kendoDDDProto.RegisterKendoGrpcServer(grpcServer, presentation.kendoRpcSer)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "127.0.0.1", beego.AppConfig.String("rpcport")))
	if err != nil {
		beeLog.Info(ctx, "failed to listen: %v\n", err)
	} else {
		beeLog.Info(ctx, "[user] running at %s:%d\n", "127.0.0.1", beego.AppConfig.String("rpcport"))
	}

	go func() {
		beeLog.Critical(ctx, "failed to serve: %v\n", grpcServer.Serve(lis))
	}()

	shutdown.GracefulStop(func() {
		beeLog.Info(ctx, "[user] shutting down...\n")

		grpcServer.GracefulStop()

		beeLog.Info(ctx, "[user] gracefully stopped\n")
	})

}
