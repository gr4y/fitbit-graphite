package fitbit

const (
	RESOURCE_ACTIVITIES_CALORIES               = "activities/calories"
	RESOURCE_ACTIVITIES_CALORIES_BMR           = "activities/caloriesBMR"
	RESOURCE_ACTIVITIES_STEPS                  = "activities/steps"
	RESOURCE_ACTIVITIES_DISTANCE               = "activities/distance"
	RESOURCE_ACTIVITIES_FLOOR                  = "activities/floors"
	RESOURCE_ACTIVITIES_ELEVATION              = "activities/elevation"
	RESOURCE_ACTIVITIES_MINUTES_SEDENTARY      = "activities/minutesSedentary"
	RESOURCE_ACTIVITIES_MINUTES_LIGHTLY_ACTIVE = "activities/minutesLightlyActive"
	RESOURCE_ACTIVITIES_MINUTES_FAIRLY_ACTIVE  = "activities/minutesFairlyActive"
	RESOURCE_ACTIVITIES_MINUTES_VERY_ACTIVE    = "activities/minutesVeryActive"
	RESOURCE_ACTIVITIES_ACTIVITY_CALORIES      = "activities/activityCalories"
)

type Activities struct {
	API Client
}

/****
Minutes Sedentary
****/
func (a *Activities) GetMinutesSedentary() ([]TimeSeriesItem, error) {
	return a.GetMinutesSedentaryForDateAndPeriod("today", "1d")
}

func (a *Activities) GetMinutesSedentaryForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_SEDENTARY, date, period)
}

/****
Minutes Lightly Active
****/
func (a *Activities) GetMinutesLightlyActive() ([]TimeSeriesItem, error) {
	return a.GetMinutesLightlyActiveForDateAndPeriod("today", "1d")
}

func (a *Activities) GetMinutesLightlyActiveForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_LIGHTLY_ACTIVE, date, period)
}

/****
Minutes Fairly Active
****/
func (a *Activities) GetMinutesFairlyActive() ([]TimeSeriesItem, error) {
	return a.GetMinutesFairlyActiveForDateAndPeriod("today", "1d")
}

func (a *Activities) GetMinutesFairlyActiveForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_FAIRLY_ACTIVE, date, period)
}

/****
Minutes Very Active
****/
func (a *Activities) GetMinutesVeryActive() ([]TimeSeriesItem, error) {
	return a.GetMinutesVeryActiveForDateAndPeriod("today", "1d")
}

func (a *Activities) GetMinutesVeryActiveForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_VERY_ACTIVE, date, period)
}

/****
Steps
****/
func (a *Activities) GetSteps() ([]TimeSeriesItem, error) {
	return a.GetStepsForDateAndPeriod("today", "1d")
}

func (a *Activities) GetStepsForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_STEPS, date, period)
}
