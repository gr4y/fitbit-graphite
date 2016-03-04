package plugins

import (
	"fmt"
	"github.com/gr4y/fitbit-graphite/lib/fitbit"
)

type WeightPlugin struct {
	Client fitbit.Client
}

func (p WeightPlugin) FetchData(start_date string, period string) {
	weight_data, err := p.Client.GetWeightForDateAndPeriod(start_date, period)
	if err != nil {
		panic(err)
	}
	fmt.Println("Weight: ", weight_data)

	bmi_data, err := p.Client.GetBMIForDateAndPeriod(start_date, period)
	if err != nil {
		panic(err)
	}
	fmt.Println("BMI: ", bmi_data)

	fat_data, err := p.Client.GetFatForDateAndPeriod(start_date, period)
	if err != nil {
		panic(err)
	}
	fmt.Println("Fat: ", fat_data)

}

func (p WeightPlugin) WriteData() {
	fmt.Println("WriteData")
}
