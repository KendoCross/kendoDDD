package routers

import (
	"../presentation"
	"github.com/astaxie/beego"
)

//InitHTTP 服务初始化，包括HTTP和RPC服务两种
//好像Golang一个进程只能监听一个端口
func InitHTTP() {
	beego.Router("/", &presentation.MainController{})
	beego.Router("/Protocol/Sign", &presentation.ProtocolController{}, "Post:Sign")
	beego.Router("/Protocol/SignCfrm", &presentation.ProtocolController{}, "*:SignCfrm")
	beego.Router("/Protocol/Sign", &presentation.ProtocolController{}, "Get:ProtocolInfo")

	beego.Router("/ws", &presentation.WebSocketController{}, "Get:Get")
	beego.Router("/WS/Token", &presentation.WebSocketController{}, "Get:GetToken")
	beego.Router("/WS/Join", &presentation.WebSocketController{}, "Get:Join")
}
