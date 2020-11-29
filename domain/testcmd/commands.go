package testcmd

import (
	"context"

	"github.com/google/uuid"
	eh "github.com/looplab/eventhorizon"
)

func init() {
	eh.RegisterCommand(func() eh.Command { return &TestOnlyCmd{} })
}

const (
	TestOnlyCmdType eh.CommandType = "TestOnlyCmd"
)

//方便开发使用的 处理机相关的命令
type TestOnlyCmd struct {
	Id      uuid.UUID
	CmdType string            `json:"cmd_type"`
	Parms   map[string]string `json:"parms"`
}

func (c *TestOnlyCmd) AggregateID() uuid.UUID {
	return c.Id
}
func (c *TestOnlyCmd) CommandType() eh.CommandType     { return TestOnlyCmdType }
func (c *TestOnlyCmd) AggregateType() eh.AggregateType { return "" }

func (cmd *TestOnlyCmd) Verify(ctx context.Context) (code int, err error) {
	return
}
