package trips

import (
	"context"
	"fmt"

	"github.com/KendoCross/kendoDDD/infrastructure/bus"
	"github.com/KendoCross/kendoDDD/infrastructure/helper"
	eh "github.com/looplab/eventhorizon"
)

func init() {
	// 注册事件handler
	bus.RegisterEventHandler(eh.MatchEvents{RateFlightsEvent}, &tripsEventHandler{})

}

type tripsEventHandler struct {
}

func (a *tripsEventHandler) HandlerType() eh.EventHandlerType {
	return ""
}
func (a *tripsEventHandler) HandleEvent(ctx context.Context, event eh.Event) (err error) {
	switch evtData := event.Data().(type) {
	case *RateFlightsInfo:
		msg := ""
		for _, item := range evtData.RateFlights {
			priceStr := ""
			for _, price := range item.Prices {
				priceStr += fmt.Sprintf("%.2f--%.0f ;", price.Rate, price.Price)
			}
			msg += fmt.Sprintf("%10s %5s %s %s <br> ", item.AirlineName, "", item.DepartureDate, priceStr)
		}
		err = helper.SendMail([]string{"1225062503@qq.com"}, "VIP", msg)
	}

	return fmt.Errorf("couldn't handle Event")
}
