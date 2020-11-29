package ddd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/KendoCross/kendoDDD/infrastructure/bus"
	eh "github.com/looplab/eventhorizon"
)

//应用层 1.主要的解析表现层数据，2.组装为正确的命令，并进行校验，3.由命令总线发布出去，不用管订阅者到底是谁。
// err 取值范围[409:errorext.CodeError,400:errorext.ListMsgError,400:validator.ValidationErrors,500:其他error]
func HandCommand(ctx context.Context, postBody []byte, commandKey string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	cmd, err := eh.CreateCommand(eh.CommandType(commandKey))
	if err != nil {
		err = fmt.Errorf("could not create command: %w", err)
		return err
	}
	if err := json.Unmarshal(postBody, &cmd); err != nil {
		err = fmt.Errorf("could not decode Json: %w", err)
		return err
	}

	if vldt, ok := cmd.(Validator); ok {
		err = vldt.Verify()
		if err != nil {
			return
		}
	} else if vldt, ok := cmd.(ValidatorCtx); ok {
		_, err = vldt.Verify(ctx)
		if err != nil {
			return
		}
	}

	if err = bus.HandleCommand(ctx, cmd); err != nil {
		return err
	}
	return
}

// err 取值范围[409:errorext.CodeError,400:errorext.ListMsgError,400:validator.ValidationErrors,500:其他error]
func DealCommand(ctx context.Context, postBody []byte, commandKey string) (rst interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
			return
		}
	}()

	cmd, err := eh.CreateCommand(eh.CommandType(commandKey))
	if err != nil {
		err = fmt.Errorf("could not create command: %w", err)
		return
	}
	if err = json.Unmarshal(postBody, &cmd); err != nil {
		err = fmt.Errorf("could not decode Json: %w", err)
		return
	}
	if vldt, ok := cmd.(Validator); ok {
		err = vldt.Verify()
		if err != nil {
			return
		}
	} else if vldt, ok := cmd.(ValidatorCtx); ok {
		_, err = vldt.Verify(ctx)
		if err != nil {
			return
		}
	}

	rst, err = bus.DealerCommand(ctx, cmd)
	if err != nil {
		return
	}
	return
}
