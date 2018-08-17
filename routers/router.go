package routers

import (
	"../presentation"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func init() {
	beego.Router("/", &MainController{})
	beego.Router("/Protocol/Sign", &presentation.ProtocolController{}, "Post:Sign")
	beego.Router("/Protocol/SignCfrm", &presentation.ProtocolController{}, "*:SignCfrm")
}
