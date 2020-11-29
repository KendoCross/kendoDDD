package trips

import (
	"github.com/KendoCross/kendoDDD/domain/services/trip"
	eh "github.com/looplab/eventhorizon"
)

const (
	RateFlightsEvent eh.EventType = "RateFlightsEvent"
)

func init() {
	// Only the event for creating an invite has custom data.
	eh.RegisterEventData(RateFlightsEvent, func() eh.EventData {
		return &RateFlightsInfo{}
	})

}

type RateFlightsInfo struct {
	RateFlights []trip.FlightInfo
}
