package processor

import (
	"github.com/gr4y/fitbit-graphite/lib/fitbit"
)

type SleepProcessor struct {
	Sleep fitbit.Sleep
}

func (p SleepProcessor) FetchData(start_date string, period string) ([]string, error) {
	var collectedData []fitbit.TimeSeriesData

	startTime, err := p.Sleep.GetStartTimeForDateAndPeriod(start_date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, startTime)

	timeInBed, err := p.Sleep.GetTimeInBedForDateAndPeriod(start_date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, timeInBed)

	minutesAsleep, err := p.Sleep.GetMinutesAsleepForDateAndPeriod(start_date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, minutesAsleep)

	awakeningsCount, err := p.Sleep.GetAwakeningsCountForDateAndPeriod(start_date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, awakeningsCount)

	minutesAwake, err := p.Sleep.GetMinutesAwakeForDateAndPeriod(start_date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, minutesAwake)

	minutesToFallAsleep, err := p.Sleep.GetMinutesToFallAsleepForDateAndPeriod(start_date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, minutesToFallAsleep)

	minutesAfterWakeup, err := p.Sleep.GetMinutesAfterWakeupForDateAndPeriod(start_date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, minutesAfterWakeup)

	efficiency, err := p.Sleep.GetSleepEfficiencyForDateAndPeriod(start_date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, efficiency)

	return convertTimeSeriesData(collectedData), nil
}
