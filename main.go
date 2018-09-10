package main

import (
	"./presentation"
	"./routers"
	"github.com/astaxie/beego"
)

func main() {
	//根据配置决定启动的服务类型。
	if beego.AppConfig.String("isHTTPSer") == "yes" {
		routers.InitHTTP()
	} else {
		presentation.InitRPC()
	}
}
