package trips

import (
	"context"
	"fmt"
	"time"

	"github.com/KendoCross/kendoDDD/domain/services/trip"
	"github.com/KendoCross/kendoDDD/infrastructure/bus"
	"github.com/google/uuid"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
)

func init() {
	sthAgg := &tripsAggregate{
		AggregateBase: events.NewAggregateBase(TripsAggregateType, uuid.Nil),
	}

	SingleTripsAgg = sthAgg
}

//想更多的表达"继承"建议使用匿名成员。具名成员表示组合。
type tripsAggregate struct {
	*events.AggregateBase //DDD框架约束
}

//Command异步执行，不需要返回值的
func (a *tripsAggregate) HandleCommand(ctx context.Context, cmd eh.Command) (err error) {
	switch cmd := cmd.(type) {
	case *AddFileCmd:
		_ = cmd
	default:
		err = fmt.Errorf("couldn't handle command")
	}
	return
}

func (a *tripsAggregate) ApplyEvent(ctx context.Context, event eh.Event) (err error) {

	return
}

//Command同步执行，需要返回值的
func (a *tripsAggregate) DealCommand(ctx context.Context, cmd eh.Command) (interface{}, error) {
	return nil, fmt.Errorf("couldn't Dealer command")
}

/////聚合根对外开放的能力
var timer time.Ticker

func (a *tripsAggregate) RegJob() (err error) {
	// cronSpec := "@every 15s"
	// c := cron.New()

	timer := time.NewTicker(time.Minute * 15)
	go func() {
		for {
			<-timer.C
			_, rst, errIn := trip.NewctripServer().GetFlights(trip.FlightsReq{
				Date:      time.Unix(1612909552, 0),
				DCity:     "szx",
				DCityName: "深圳",
				ACity:     "sia",
				ACityName: "西安",
			})
			if errIn != nil {
				return
			}
			a.AppendEvent(RateFlightsEvent, &RateFlightsInfo{RateFlights: rst}, time.Now())
			bus.RaiseEvents(context.Background(), a.AggregateBase, 1)
		}
	}()

	// _, err = c.AddFunc(cronSpec, func() {

	// })
	return
}

///// 聚合根，对内的调度
