package presentation

import (
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

func MapParams(c *beego.Controller) map[string]string {
	m := c.Ctx.Input.Params()
	if len(m) == 0 {
		if c.Ctx.Input.Context.Request.Form == nil {
			c.Ctx.Input.Context.Request.ParseForm()
		}
		forms := c.Ctx.Input.Context.Request.Form
		for key, values := range forms {
			if len(values) > 0 {
				m[key] = values[0]
			}
		}
	}
	return m
}
