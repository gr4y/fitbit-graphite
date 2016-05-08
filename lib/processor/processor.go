package processor

import (
	"fmt"
	"github.com/gr4y/fitbit-graphite/lib/fitbit"
	"strings"
)

type Processor interface {
	FetchData(start_date string, period string) ([]string, error)
	// WriteData()
}

func convertTimeSeriesData(collectedData []fitbit.TimeSeriesData) []string {
	var lines []string
	for _, dataItem := range collectedData {
		for key, items := range dataItem {
			for _, datapoint := range items {
				lines = append(lines, fmt.Sprintf("%s %s %d", strings.Replace(key, "-", ".", 1), datapoint.Value, datapoint.DateTime.Unix()))
			}
		}
	}
	return lines
}
