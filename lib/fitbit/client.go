package fitbit

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	DEBUG      = true
	BASE_URL   = "https://api.fitbit.com/1/user/-"
	TOKEN_FILE = "token.json"
)

type TimeSeriesData map[string][]TimeSeriesItem
type TimeSeriesItem struct {
	DateTime CustomDate `json:"dateTime"`
	Value    string     `json:"value"`
}

type CustomDate struct {
	time.Time
}

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}

	timeObj, err := time.Parse("2006-01-02", string(b))
	if err != nil {
		return err
	}
	*cd = CustomDate{Time: timeObj}
	return nil
}

type CodeCallbackFunc func(url string) (code string)
type API struct {
	Profile    Profile
	Activities Activities
	Body       Body
	Sleep      Sleep
}

type Client struct {
	Client http.Client
}

// Returns a new fitbit.API Object.
func Connect(conf ClientConfig, codeCallback CodeCallbackFunc) (API, error) {
	oAuthConfig := &oauth2.Config{
		ClientID: conf.ClientID, ClientSecret: conf.ClientSecret, Scopes: conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://fitbit.com/oauth2/authorize",
			TokenURL: "https://api.fitbit.com/oauth2/token",
		},
	}

	// Read Token from JSON
	token, err := tokenFromJSON()

	// If Token is nil, or error is not nil, we will fetch a new AccessToken
	if err != nil || token == nil {
		state := random_string(128)
		url := oAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOnline)
		code := codeCallback(url)
		if code != "" {
			token, err = oAuthConfig.Exchange(oauth2.NoContext, code)
			if err != nil {
				return API{}, err
			}
		} else {
			return API{}, errors.New("code can't be empty.")
		}
	}

	// Create a TokenSource with the provided token
	// This will automatically refesh the AccessToken when it, thanks to the golang.org/x/oauth2 package.
	tokenSrc := oAuthConfig.TokenSource(oauth2.NoContext, token)
	newToken, err := tokenSrc.Token()
	if err != nil {
		return API{}, err
	}

	// Initialize new Client with TokenSource
	newHttpClient := oauth2.NewClient(oauth2.NoContext, tokenSrc)

	// Saves the new token, just in case it was refreshed.
	tokenToJSON(newToken)

	// Intialize new API Client, wrapping an simple http.Client
	apiClient := Client{Client: *newHttpClient}

	return API{
		Profile:    Profile{API: apiClient},
		Activities: Activities{API: apiClient},
		Body:       Body{API: apiClient},
		Sleep:      Sleep{API: apiClient},
	}, nil
}

func tokenToJSON(token *oauth2.Token) error {
	if d, err := json.Marshal(token); err != nil {
		return err
	} else {
		return ioutil.WriteFile(TOKEN_FILE, d, 0644)
	}
}

func tokenFromJSON() (*oauth2.Token, error) {
	var token oauth2.Token
	bytes, err := ioutil.ReadFile(TOKEN_FILE)
	if err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(bytes, &token); err != nil {
			return nil, err
		}
		return &token, nil
	}
}

func random_string(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

/****
Internal Functions
****/
func (c *Client) getTimeSeriesData(resourcePath string, date string, period string) (TimeSeriesData, error) {
	var timeSeriesData TimeSeriesData
	url := c.buildUrl(resourcePath, date, period)
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if DEBUG {
		fmt.Println("#### DEBUG Output Begin ####")
		fmt.Println(fmt.Sprintf("Request URL: %s\r\n", url))
		fmt.Println("\r\n### HTTP Request ###")
		fmt.Println(fmt.Sprintf("Request: %+v", resp.Request))
		fmt.Println("\r\n### HTTP Response")
		fmt.Println(fmt.Sprintf("Response Object: %+v", resp))
		fmt.Println(fmt.Sprintf("Response Body: %s", string(bytes)))
		fmt.Println("#### DEBUG Output End ####")
	}

	if err := json.Unmarshal(bytes, &timeSeriesData); err != nil {
		return nil, err
	}
	return timeSeriesData, nil
}

func (c *Client) buildUrl(resourcePath string, date string, period string) string {
	base_url := fmt.Sprintf("%s/%s", BASE_URL, resourcePath)
	if date == "" && period == "" {
		return fmt.Sprintf("%s.json", base_url)
	} else {
		return fmt.Sprintf("%s/date/%s/%s.json", base_url, date, period)
	}
}
