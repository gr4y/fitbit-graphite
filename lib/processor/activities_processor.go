package processor

import (
	"github.com/gr4y/fitbit-graphite/lib/fitbit"
)

type ActivitiesProcessor struct {
	Activities fitbit.Activities
}

func (p ActivitiesProcessor) FetchData(date string, period string) ([]string, error) {
	var collectedData []fitbit.TimeSeriesData

	steps, err := p.Activities.GetStepsForDateAndPeriod(date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, steps)

	distance, err := p.Activities.GetDistanceForDateAndPeriod(date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, distance)

	floors, err := p.Activities.GetFloorsForDateAndPeriod(date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, floors)

	elevation, err := p.Activities.GetElevationForDateAndPeriod(date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, elevation)

	minutesSedentary, err := p.Activities.GetMinutesSedentaryForDateAndPeriod(date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, minutesSedentary)

	minutesLightlyActive, err := p.Activities.GetMinutesLightlyActiveForDateAndPeriod(date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, minutesLightlyActive)

	minutesFairlyActive, err := p.Activities.GetMinutesFairlyActiveForDateAndPeriod(date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, minutesFairlyActive)

	minutesVeryActive, err := p.Activities.GetMinutesVeryActiveForDateAndPeriod(date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, minutesVeryActive)

	calories, err := p.Activities.GetCaloriesForDateAndPeriod(date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, calories)

	caloriesBMR, err := p.Activities.GetCaloriesBMRForDateAndPeriod(date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, caloriesBMR)

	activityCalories, err := p.Activities.GetActivityCaloriesForDateAndPeriod(date, period)
	if err != nil {
		return nil, err
	}
	collectedData = append(collectedData, activityCalories)

	return convertTimeSeriesData(collectedData), nil
}
