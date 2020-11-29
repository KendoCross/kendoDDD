package trip

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

type ctripServer struct {
	*http.Client
}

func NewctripServer() *ctripServer {
	server := &ctripServer{
		Client: &http.Client{
			Timeout: time.Duration(3 * time.Second),
		},
	}
	return server
}

func (ser *ctripServer) GetFlights(req FlightsReq) (flights []FlightInfo, rateFlights []FlightInfo, err error) {
	uri := "https://flights.ctrip.com/itinerary/api/12808/products"

	req_payload := map[string]interface{}{
		"token":       "536b912ebd54dcd63f8cf194f526a1c1",
		"flightWay":   "Oneway",
		"classType":   "ALL",
		"hasChild":    false,
		"hasBaby":     false,
		"searchIndex": 1,
		"airportParams": []map[string]interface{}{
			{"dcity": req.DCity, "acity": req.ACity, "dcityname": req.DCityName, "acityname": req.ACityName,
				"date": req.Date.Format("2006-01-02")},
		},
	}

	body, err := json.Marshal(req_payload)
	if err != nil {
		return
	}

	request, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
		return
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	request.Header.Add("Referer", fmt.Sprintf("https://flights.ctrip.com/itinerary/oneway/%s-%s?date=%s", req.DCity, req.ACity, req.Date.Format("2006-01-02")))
	request.Header.Add("Content-Type", "application/json")

	resp, err := ser.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("HTTP StatusCode: %d", resp.StatusCode)
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	respJson := gjson.ParseBytes(data)

	if respJson.Get("status").Int() != 0 || respJson.Get("data.error.code").Str != "" {
		err = fmt.Errorf("%s %s", respJson.Get("msg").String(), respJson.Get("data.error.msg").String())
		return
	}

	rateFlights = make([]FlightInfo, 0)
	flights = make([]FlightInfo, 0)
	for _, item := range respJson.Get("data.routeList").Array() {
		legs := item.Get("legs").Array()
		if len(legs) > 0 {
			flight := FlightInfo{}
			flt := legs[0].Get("flight")
			flight.AirlineName = flt.Get("airlineName").Str
			flight.FlightNumber = flt.Get("flightNumber").Str
			flight.DepartureDate = flt.Get("departureDate").Str
			flight.ArrivalDate = flt.Get("arrivalDate").Str
			for _, cabin := range legs[0].Get("cabins").Array() {
				rate := cabin.Get("price.rate").Float()
				if rate < 0.9 {
					flight.Prices = append(flight.Prices, Price{
						Price:      cabin.Get("price.price").Float(),
						SalePrice:  cabin.Get("price.salePrice").Float(),
						PrintPrice: cabin.Get("price.printPrice").Float(),
						FdPrice:    cabin.Get("price.fdPrice").Float(),
						Rate:       rate,
					})
				}
			}

			for _, price := range flight.Prices {
				if price.Rate < 0.9 {
					rateFlights = append(rateFlights, flight)
					break
				}
			}

			flights = append(flights, flight)
		}
	}

	return
}
