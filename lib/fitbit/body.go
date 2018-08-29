package fitbit

const (
	RESOURCE_BODY_WEIGHT = "body/weight"
	RESOURCE_BODY_BMI    = "body/bmi"
	RESOURCE_BODY_FAT    = "body/fat"
	RESOURCE_HEART       = "body/heart"
)

type Body struct {
	API Client
}

/****
Weight
****/
func (b *Body) GetWeight() (TimeSeriesData, error) {
	return b.GetWeightForDateAndPeriod("today", "1d")
}

func (b *Body) GetWeightForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return b.API.getTimeSeriesData(RESOURCE_BODY_WEIGHT, date, period)
}

/****
BMI
****/
func (b *Body) GetBMI() (TimeSeriesData, error) {
	return b.GetBMIForDateAndPeriod("today", "1d")
}

func (b *Body) GetBMIForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return b.API.getTimeSeriesData(RESOURCE_BODY_BMI, date, period)
}

/****
Fat
****/
func (b *Body) GetFat() (TimeSeriesData, error) {
	return b.GetFatForDateAndPeriod("today", "1d")
}

func (b *Body) GetFatForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return b.API.getTimeSeriesData(RESOURCE_BODY_FAT, date, period)
}

/****
Fat
****/
func (b *Body) GetHeart() (TimeSeriesData, error) {
	return b.GetHeartForDateAndPeriod("today", "1d")
}

func (b *Body) GetHeartForDateAndPeriod(date string, period string) (TimeSeriesData, error) {
	return b.API.getTimeSeriesData(RESOURCE_HEART, date, period)
}
