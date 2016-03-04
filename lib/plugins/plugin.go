package plugins

type Plugin interface {
	FetchData(start_date string, period string)
	WriteData()
}
