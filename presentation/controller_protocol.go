package presentation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	//context "golang.org/x/net/context"
)

type ProtocolController struct {
	beego.Controller
}

func (c *ProtocolController) Sign() {
	ctx := context.Background()
	result, err := kendoHttpSer.DispatchCommand(ctx, SingnProtocolByInfo, c.Ctx.Input.RequestBody)
	if err != nil {
		errStr := fmt.Sprintf("出错啦:%v\n", err)
		http.Error(c.Ctx.ResponseWriter, errStr, 500)
		return
	}
	c.Ctx.Output.Body([]byte(result))
}

func (this *ProtocolController) SignCfrm() {
	ctx := context.Background()
	result, err := kendoHttpSer.DispatchCommandMap(ctx, CfrmProtocolByReqSn, MapParams(&this.Controller))
	if err != nil {
		errStr := fmt.Sprintf("出错啦:%v\n", err)
		http.Error(this.Ctx.ResponseWriter, errStr, 500)
		return
	}
	this.Ctx.Output.Body([]byte(result))
}

func (this *ProtocolController) ProtocolInfo() {
	ctx := context.Background()
	result, err := kendoHttpSer.DispatchCommandMap(ctx, GetProtocolInfoByNo, MapParams(&this.Controller))
	if err != nil {
		errStr := fmt.Sprintf("出错啦:%v\n", err)
		http.Error(this.Ctx.ResponseWriter, errStr, 500)
		return
	}
	this.Ctx.Output.Body([]byte(result))
}
