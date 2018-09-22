package main

import (
	"./presentation"
	"./routers"
	"github.com/astaxie/beego"
)

func main() {
	//根据配置决定启动的服务类型。 #是HTTP服务或RPC服务或WebSocket
	serTp := beego.AppConfig.String("SerType")
	switch serTp {
	case "RPC":
		presentation.InitRPC()
	case "HTTP":
	default:
		routers.InitHTTP()
		beego.Run()
	}
}
