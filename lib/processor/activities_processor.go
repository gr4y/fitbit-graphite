package processor

import (
	"fmt"
	"github.com/gr4y/fitbit-graphite/lib/fitbit"
)

type ActivitiesProcessor struct {
	Activities fitbit.Activities
}

func (p ActivitiesProcessor) FetchData(start_date string, period string) {
	_, err := p.fetchSteps(start_date, period)
	if err != nil {
		panic(err)
	}
}

func (p ActivitiesProcessor) fetchSteps(start_date string, period string) (map[string]string, error) {
	steps, err := p.Activities.GetStepsForDateAndPeriod(start_date, period)
	if err != nil {
		return nil, err
	}
	for _, datapoint := range steps {
		fmt.Println(datapoint.DateTime, ": ", datapoint.Value)
	}
	return nil, nil
}
