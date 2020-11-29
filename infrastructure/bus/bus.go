package bus

import (
	"context"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/commandhandler/bus"
	eventbus "github.com/looplab/eventhorizon/eventbus/local"
)

//commandBus 命令总线
var commandBus = bus.NewCommandHandler()

//eventBus 事件总线
var eventBus = eventbus.NewEventBus(nil)

var eventStore eh.EventStore

//aggregateStore 领域事件存储与发布
//var AggregateStore *events.AggregateStore

func InitBus() {
	//eventStore, _ = eventstore.NewEventStore("112.74.187.73:27017", "EventStore")
	//AggregateStore, _ = events.NewAggregateStore(eventStore, eventBus)
}

//RegisterHandler 注册命令的处理
func RegisterHandler(cmd eh.CommandType, cmdHandler eh.Aggregate) {
	err := commandBus.SetHandler(cmdHandler, cmd)
	if err != nil {
		panic(err)
	}
}

//HandleCommand 命令的执行
func HandleCommand(ctx context.Context, cmd eh.Command) error {
	return commandBus.HandleCommand(ctx, cmd)
}

//RegisterEventHandler 注册事件的处理
func RegisterEventHandler(evtMatcher eh.EventMatcher, evtHandler eh.EventHandler) {
	eventBus.AddHandler(context.TODO(), evtMatcher, evtHandler)
}

//RaiseEvents 异步进行事件的存储 和 发布
func RaiseEvents(ctx context.Context, aggre *events.AggregateBase, originalVersion int) error {
	_ = originalVersion
	// TODO 事件存储器未开启后，方可打开以下注释
	// go eventStore.Save(ctx, aggre.Events(), originalVersion)
	for _, event := range aggre.Events() {
		err := eventBus.HandleEvent(ctx, event)
		if err != nil {
			return err
		}
	}
	//aggre.ClearEvents()
	return nil
}
