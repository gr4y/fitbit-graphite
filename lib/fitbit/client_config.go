package fitbit

type ClientConfig struct {
	ClientID     string
	ClientSecret string
	Debug        bool
	Scopes       []string
}
