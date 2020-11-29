package trip

import "time"

type ITripSer interface {
	GetFlights(req FlightsReq) (flights []FlightInfo, err error)
}

type FlightsReq struct {
	Date      time.Time
	DCity     string
	DCityName string
	ACity     string
	ACityName string
	ExtData   map[string]string
}

type FlightInfo struct {
	AirlineName   string
	FlightNumber  string
	DepartureDate string
	ArrivalDate   string
	Prices        []Price
}

type Price struct {
	Price      float64
	SalePrice  float64
	PrintPrice float64
	FdPrice    float64
	Rate       float64
}
