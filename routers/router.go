package routers

import (
	"../presentation"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &presentation.MainController{})
	beego.Router("/Protocol/Sign", &presentation.ProtocolController{}, "Post:Sign")
	beego.Router("/Protocol/SignCfrm", &presentation.ProtocolController{}, "*:SignCfrm")
	beego.Router("/Protocol/Sign", &presentation.ProtocolController{}, "Get:ProtocolInfo")
}
