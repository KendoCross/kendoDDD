/*业务逻辑无关，很薄的一层。只是作为计算机领域
到业务领域的过渡，应用型功能*/
package ddd_application

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego/logs"

	"../ddd_domain/protocol"
	"../ddd_infrastructure"
	"../dddcore"
)

func fromJSON(c *protocol.SignModel, payload json.RawMessage) error {
	return json.Unmarshal(payload, c)
}

// creates command handler
func OnSingnProtocol(es dddcore.EventStore, eb dddcore.EventBus) dddcore.CommandHandler {
	esRepo := ddd_infrastructure.NewESRepository(fmt.Sprintf("%T", protocol.New()), es, eb)
	return func(ctx context.Context, payload []byte, out chan<- dddcore.BusChan) {

		log.Println(string(payload))

		pJs := new(protocol.SignModel)
		fromJSON(pJs, payload)
		p := protocol.New()
		p.SignVm = pJs

		reqSn, err := p.Sign()
		if err != nil {
			out <- dddcore.BusChan{RespMsg: "", ErrMsg: err}
			return
		}

		log.Println(reqSn)
		out <- dddcore.BusChan{RespMsg: reqSn, ErrMsg: nil}

		esRepo.Save(ctx, p)
	}
}

func OnCfrmProtocol(es dddcore.EventStore, eb dddcore.EventBus) dddcore.CommandHandlerMap {
	// esRepo := ddd_infrastructure.NewESRepository(fmt.Sprintf("%T", protocol.New()), es, eb)
	return func(ctx context.Context, payload map[string]string, out chan<- dddcore.BusChan) {

		beeLog := logs.GetBeeLogger()
		beeLog.Notice("签约确认参数：%v", payload)

		p := protocol.New()
		retStr, err := p.CfrmSign(payload["reqSn"], payload["verCode"])
		if err != nil {
			out <- dddcore.BusChan{RespMsg: "", ErrMsg: err}
			return
		}

		beeLog.Notice(retStr)
		out <- dddcore.BusChan{RespMsg: retStr, ErrMsg: nil}
		// esRepo.Save(ctx, p)
	}
}
