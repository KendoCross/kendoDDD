package k8s_info

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
	v1 "k8s.io/api/core/v1"
)

func init() {
	sthAgg := &k8sInfoAggregate{
		AggregateBase: events.NewAggregateBase(K8SInfoAggregateType, uuid.Nil),
	}

	SingleK8sInfoAgg = sthAgg
}

//想更多的表达"继承"建议使用匿名成员。具名成员表示组合。
type k8sInfoAggregate struct {
	*events.AggregateBase //DDD框架约束
}

//Command异步执行，不需要返回值的
func (a *k8sInfoAggregate) HandleCommand(ctx context.Context, cmd eh.Command) (err error) {
	switch cmd := cmd.(type) {
	case *AddFileCmd:
		_ = cmd
	default:
		err = fmt.Errorf("couldn't handle command")
	}
	return
}

func (a *k8sInfoAggregate) ApplyEvent(ctx context.Context, event eh.Event) (err error) {

	return
}

//Command同步执行，需要返回值的
func (a *k8sInfoAggregate) DealCommand(ctx context.Context, cmd eh.Command) (interface{}, error) {
	return nil, fmt.Errorf("couldn't Dealer command")
}

/////聚合根对外开放的能力

func (a *k8sInfoAggregate) PodList() (pods *v1.PodList, err error) {
	en := newpodEnByOV()
	return en.PodList()
}

///// 聚合根，对内的调度
