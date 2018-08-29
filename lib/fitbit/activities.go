package fitbit

const (
	RESOURCE_ACTIVITIES_CALORIES               = "activities/calories"
	RESOURCE_ACTIVITIES_CALORIES_BMR           = "activities/caloriesBMR"
	RESOURCE_ACTIVITIES_STEPS                  = "activities/steps"
	RESOURCE_ACTIVITIES_DISTANCE               = "activities/distance"
	RESOURCE_ACTIVITIES_FLOORS                 = "activities/floors"
	RESOURCE_ACTIVITIES_ELEVATION              = "activities/elevation"
	RESOURCE_ACTIVITIES_MINUTES_SEDENTARY      = "activities/minutesSedentary"
	RESOURCE_ACTIVITIES_MINUTES_LIGHTLY_ACTIVE = "activities/minutesLightlyActive"
	RESOURCE_ACTIVITIES_MINUTES_FAIRLY_ACTIVE  = "activities/minutesFairlyActive"
	RESOURCE_ACTIVITIES_MINUTES_VERY_ACTIVE    = "activities/minutesVeryActive"
	RESOURCE_ACTIVITIES_ACTIVITY_CALORIES      = "activities/activityCalories"
	RESOURCE_ACTIVITIES_HEART		   = "activities/heart"
)

type Activities struct {
	API Client
}

/****
Heart
****/
func (b *Body) GetHeartRate() (TimeSeriesData, error) {
	return a.GetHeartForDateAndPeriod("today", "1d")
}

func (b *Body) GetHeartRateForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_HEART, date, period)
}

/*** Calories ***/
func (a *Activities) GetCalories() (TimeSeriesData, error) {
	return a.GetCaloriesForDateAndPeriod("today", "1d")
}

func (a *Activities) GetCaloriesForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_CALORIES, date, period)
}

/*** CaloriesBMR ***/
func (a *Activities) GetCaloriesBMR() (TimeSeriesData, error) {
	return a.GetCaloriesBMRForDateAndPeriod("today", "1d")
}

func (a *Activities) GetCaloriesBMRForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_CALORIES_BMR, date, period)
}

/*** ActivityCalories ***/
func (a *Activities) GetActivityCalories() (TimeSeriesData, error) {
	return a.GetActivityCaloriesForDateAndPeriod("today", "1d")
}

func (a *Activities) GetActivityCaloriesForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_ACTIVITY_CALORIES, date, period)
}

/****
Minutes Sedentary
****/
func (a *Activities) GetMinutesSedentary() (TimeSeriesData, error) {
	return a.GetMinutesSedentaryForDateAndPeriod("today", "1d")
}

func (a *Activities) GetMinutesSedentaryForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_SEDENTARY, date, period)
}

/****
Minutes Lightly Active
****/
func (a *Activities) GetMinutesLightlyActive() (TimeSeriesData, error) {
	return a.GetMinutesLightlyActiveForDateAndPeriod("today", "1d")
}

func (a *Activities) GetMinutesLightlyActiveForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_LIGHTLY_ACTIVE, date, period)
}

/****
Minutes Fairly Active
****/
func (a *Activities) GetMinutesFairlyActive() (TimeSeriesData, error) {
	return a.GetMinutesFairlyActiveForDateAndPeriod("today", "1d")
}

func (a *Activities) GetMinutesFairlyActiveForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_FAIRLY_ACTIVE, date, period)
}

/****
Minutes Very Active
****/
func (a *Activities) GetMinutesVeryActive() (TimeSeriesData, error) {
	return a.GetMinutesVeryActiveForDateAndPeriod("today", "1d")
}

func (a *Activities) GetMinutesVeryActiveForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_VERY_ACTIVE, date, period)
}

/****
Steps
****/
func (a *Activities) GetSteps() (TimeSeriesData, error) {
	return a.GetStepsForDateAndPeriod("today", "1d")
}

func (a *Activities) GetStepsForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_STEPS, date, period)
}

/****
Distance
****/
func (a *Activities) GetDistance() (TimeSeriesData, error) {
	return a.GetDistanceForDateAndPeriod("today", "1d")
}

func (a *Activities) GetDistanceForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_DISTANCE, date, period)
}

/****
Floors
****/
func (a *Activities) GetFloors() (TimeSeriesData, error) {
	return a.GetFloorsForDateAndPeriod("today", "1d")
}

func (a *Activities) GetFloorsForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_FLOORS, date, period)
}

/****
Elevation
****/
func (a *Activities) GetElevation() (TimeSeriesData, error) {
	return a.GetElevationForDateAndPeriod("today", "1d")
}

func (a *Activities) GetElevationForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return a.API.getTimeSeriesData(RESOURCE_ACTIVITIES_ELEVATION, date, period)
}
