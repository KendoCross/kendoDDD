package domain

import (
	"github.com/KendoCross/kendoDDD/domain/k8s_info"
	"github.com/KendoCross/kendoDDD/domain/testcmd"
	"github.com/KendoCross/kendoDDD/domain/trips"
)

func StartUp() (err error) {
	err = testcmd.StartUp()
	if err != nil {
		return
	}

	err = k8s_info.StartUp()
	if err != nil {
		return
	}

	err = trips.SingleTripsAgg.RegJob()
	if err != nil {
		return
	}
	return
}
