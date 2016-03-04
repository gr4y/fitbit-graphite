package plugins

import (
	"fmt"
	"github.com/gr4y/fitbit-graphite/lib/fitbit"
)

type StepsPlugin struct {
	Client fitbit.Client
}

func (p StepsPlugin) FetchData(start_date string, period string) {
	data, err := p.Client.GetStepsForDateAndPeriod(start_date, period)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

	// for _, item := range data.Items {
	// 	fmt.Println(item)
	// }
}

func (p StepsPlugin) WriteData() {

}
