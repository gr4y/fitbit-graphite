package processor

type Processor interface {
	FetchData(start_date string, period string)
	// WriteData()
}
