package fitbit

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	BASE_URL             = "https://api.fitbit.com/1/user/-"
	RESOURCE_BODY_WEIGHT = "body/weight"
	RESOURCE_BODY_BMI    = "body/bmi"
	RESOURCE_BODY_FAT    = "body/fat"
	TOKEN_FILE           = "token.json"
)

type Data map[string][]TimeSeriesItem
type TimeSeriesItem struct {
	DateTime string `json:"dateTime"`
	Value    string `json:"value"`
}

type CodeCallbackFunc func(url string) (code string)
type API struct {
	Activities Activities
	Body       Body
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
		Activities: Activities{API: apiClient},
		Body:       Body{API: apiClient},
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
func (c *Client) getTimeSeriesData(resourcePath string, date string, period string) ([]TimeSeriesItem, error) {
	var timeSeriesData Data
	url := c.buildUrl(resourcePath, date, period)
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println(fmt.Sprintf("URL %s: \r\n Response: %s", url, string(bytes)))
	if err := json.Unmarshal(bytes, &timeSeriesData); err != nil {
		return nil, err
	}
	key := strings.Replace(resourcePath, "/", "-", 1)
	return timeSeriesData[key], nil
}

func (c *Client) buildUrl(resourcePath string, date string, period string) string {
	return fmt.Sprintf("%s/%s/date/%s/%s.json", BASE_URL, resourcePath, date, period)
}
