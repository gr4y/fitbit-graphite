package processor

import (
	"fmt"
	"github.com/gr4y/fitbit-graphite/lib/fitbit"
)

type BodyProcessor struct {
	Body fitbit.Body
}

func (p BodyProcessor) FetchData(start_date string, period string) {
	weight_data, err := p.Body.GetWeightForDateAndPeriod(start_date, period)
	if err != nil {
		panic(err)
	}
	for _, datapoint := range weight_data {
		fmt.Println(datapoint.DateTime, ": ", datapoint.Value)
	}

	bmi_data, err := p.Body.GetBMIForDateAndPeriod(start_date, period)
	if err != nil {
		panic(err)
	}
	for _, datapoint := range bmi_data {
		fmt.Println(datapoint.DateTime, ": ", datapoint.Value)
	}

	fat_data, err := p.Body.GetFatForDateAndPeriod(start_date, period)
	if err != nil {
		panic(err)
	}
	for _, datapoint := range fat_data {
		fmt.Println(datapoint.DateTime, ": ", datapoint.Value)
	}

}
