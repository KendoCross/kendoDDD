package routers

import (
	"../presentation"
	"github.com/astaxie/beego"
)

//服务初始化，包括HTTP和RPC服务两种
func init() {
	if beego.AppConfig.String("isHTTPSer") == "yes" {
		beego.Router("/", &presentation.MainController{})
		beego.Router("/Protocol/Sign", &presentation.ProtocolController{}, "Post:Sign")
		beego.Router("/Protocol/SignCfrm", &presentation.ProtocolController{}, "*:SignCfrm")
		beego.Router("/Protocol/Sign", &presentation.ProtocolController{}, "Get:ProtocolInfo")
	} else {
		presentation.InitRPC()
	}
}
