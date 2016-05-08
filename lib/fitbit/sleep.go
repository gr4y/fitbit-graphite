package fitbit

const (
	RESOURCE_SLEEP_START_TIME             = "sleep/startTime"
	RESOURCE_SLEEP_TIME_IN_BED            = "sleep/timeInBed"
	RESOURCE_SLEEP_MINUTES_ASLEEP         = "sleep/minutesAsleep"
	RESOURCE_SLEEP_AWAKENINGS_COUNT       = "sleep/awakeningsCount"
	RESOURCE_SLEEP_MINUTES_AWAKE          = "sleep/minutesAwake"
	RESOURCE_SLEEP_MINUTES_TO_FALL_ASLEEP = "sleep/minutesToFallAsleep"
	RESOURCE_MINUTES_AFTER_WAKEUP         = "sleep/minutesAfterWakeup"
	RESOURCE_SLEEP_EFFICIENCY             = "sleep/efficiency"
)

type Sleep struct {
	API Client
}

/*** Sleep Start Time ***/
func (sl *Sleep) GetStartTime() (TimeSeriesData, error) {
	return sl.GetStartTimeForDateAndPeriod("today", "1d")
}

func (sl *Sleep) GetStartTimeForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return sl.API.getTimeSeriesData(RESOURCE_SLEEP_START_TIME, date, period)
}

/*** Time In Bed ***/
func (sl *Sleep) GetTimeInBed() (TimeSeriesData, error) {
	return sl.GetTimeInBedForDateAndPeriod("today", "1d")
}

func (sl *Sleep) GetTimeInBedForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return sl.API.getTimeSeriesData(RESOURCE_SLEEP_TIME_IN_BED, date, period)
}

/*** Minutes Asleep ***/
func (sl *Sleep) GetMinutesAsleep() (TimeSeriesData, error) {
	return sl.GetMinutesAsleepForDateAndPeriod("today", "1d")
}

func (sl *Sleep) GetMinutesAsleepForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return sl.API.getTimeSeriesData(RESOURCE_SLEEP_MINUTES_ASLEEP, date, period)
}

/*** Awakenings Count ***/
func (sl *Sleep) GetAwakeningsCount() (TimeSeriesData, error) {
	return sl.GetAwakeningsCountForDateAndPeriod("today", "1d")
}

func (sl *Sleep) GetAwakeningsCountForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return sl.API.getTimeSeriesData(RESOURCE_SLEEP_AWAKENINGS_COUNT, date, period)
}

/*** Minutes Awake ***/
func (sl *Sleep) GetMinutesAwake() (TimeSeriesData, error) {
	return sl.GetMinutesAwakeForDateAndPeriod("today", "1d")
}

func (sl *Sleep) GetMinutesAwakeForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return sl.API.getTimeSeriesData(RESOURCE_SLEEP_MINUTES_AWAKE, date, period)
}

/*** Minutes To Fall Asleep ***/
func (sl *Sleep) GetMinutesToFallAsleep() (TimeSeriesData, error) {
	return sl.GetMinutesToFallAsleepForDateAndPeriod("today", "1d")
}

func (sl *Sleep) GetMinutesToFallAsleepForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return sl.API.getTimeSeriesData(RESOURCE_SLEEP_MINUTES_TO_FALL_ASLEEP, date, period)
}

/*** Minutes After Wakeup ***/
func (sl *Sleep) GetMinutesAfterWakeup() (TimeSeriesData, error) {
	return sl.GetMinutesAfterWakeupForDateAndPeriod("today", "1d")
}

func (sl *Sleep) GetMinutesAfterWakeupForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return sl.API.getTimeSeriesData(RESOURCE_MINUTES_AFTER_WAKEUP, date, period)
}

/*** Sleep Efficiency ***/
func (sl *Sleep) GetSleepEfficiency() (TimeSeriesData, error) {
	return sl.GetSleepEfficiencyForDateAndPeriod("today", "1d")
}

func (sl *Sleep) GetSleepEfficiencyForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return sl.API.getTimeSeriesData(RESOURCE_MINUTES_AFTER_WAKEUP, date, period)
}
