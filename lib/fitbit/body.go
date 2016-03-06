package fitbit

type Body struct {
	API Client
}

/****
Weight
****/
func (b *Body) GetWeight() ([]TimeSeriesItem, error) {
	return b.GetWeightForDateAndPeriod("today", "1d")
}

func (b *Body) GetWeightForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return b.API.getTimeSeriesData(RESOURCE_BODY_WEIGHT, date, period)
}

/****
BMI
****/
func (b *Body) GetBMI() ([]TimeSeriesItem, error) {
	return b.GetBMIForDateAndPeriod("today", "1d")
}

func (b *Body) GetBMIForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return b.API.getTimeSeriesData(RESOURCE_BODY_BMI, date, period)
}

/****
Fat
****/
func (b *Body) GetFat() ([]TimeSeriesItem, error) {
	return b.GetFatForDateAndPeriod("today", "1d")
}

func (b *Body) GetFatForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return b.API.getTimeSeriesData(RESOURCE_BODY_FAT, date, period)
}
