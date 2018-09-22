package ddd_application

import (
	"context"

	"github.com/astaxie/beego/logs"

	"../ddd_domain/tokens"
	"../dddcore"
)

//
func OnGetWsToken(es dddcore.EventStore, eb dddcore.EventBus) dddcore.CommandHandlerMap {
	return func(ctx context.Context, payload map[string]string, out chan<- dddcore.BusChan) {

		beeLog := logs.GetBeeLogger()
		beeLog.Notice("获取长连接Token参数：%v", payload)

		ts := tokens.New()
		token, err := ts.CreatToken(payload["userID"])
		if err != nil {
			out <- dddcore.BusChan{RespMsg: "", ErrMsg: err}
			return
		}
		beeLog.Notice("获取长连接Token成功！")
		out <- dddcore.BusChan{RespMsg: token, ErrMsg: nil}
	}
}
