package processor

import (
	"github.com/gr4y/fitbit-graphite/lib/fitbit"
)

type BodyProcessor struct {
	Body fitbit.Body
}

func (p BodyProcessor) FetchData(start_date string, period string) ([]string, error) {
	var collectedData []fitbit.TimeSeriesData
	weight_data, err := p.Body.GetWeightForDateAndPeriod(start_date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, weight_data)
	bmi_data, err := p.Body.GetBMIForDateAndPeriod(start_date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, bmi_data)
	fat_data, err := p.Body.GetFatForDateAndPeriod(start_date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, bmi_data)
	fat_data, err := p.Body.GetHeartForDateAndPeriod(start_date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, fat_data)

	return convertTimeSeriesData(collectedData), nil
}
