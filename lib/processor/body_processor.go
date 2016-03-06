package processor

import (
	"fmt"
	"github.com/gr4y/fitbit-graphite/lib/fitbit"
)

type BodyProcessor struct {
	Client fitbit.Client
}

func (p BodyProcessor) FetchData(start_date string, period string) {
	weight_data, err := p.Client.GetWeightForDateAndPeriod(start_date, period)
	if err != nil {
		panic(err)
	}
	for _, datapoint := range weight_data {
		fmt.Println(datapoint.DateTime, ": ", datapoint.Value)
	}

	bmi_data, err := p.Client.GetBMIForDateAndPeriod(start_date, period)
	if err != nil {
		panic(err)
	}
	for _, datapoint := range bmi_data {
		fmt.Println(datapoint.DateTime, ": ", datapoint.Value)
	}

	fat_data, err := p.Client.GetFatForDateAndPeriod(start_date, period)
	if err != nil {
		panic(err)
	}
	for _, datapoint := range fat_data {
		fmt.Println(datapoint.DateTime, ": ", datapoint.Value)
	}

}