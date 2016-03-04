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
	BASE_URL                                   = "https://api.fitbit.com/1/user/-"
	RESOURCE_ACTIVITIES_CALORIES               = "activities/calories"
	RESOURCE_ACTIVITIES_CALORIES_BMR           = "activities/caloriesBMR"
	RESOURCE_ACTIVITIES_STEPS                  = "activities/steps"
	RESOURCE_ACTIVITIES_DISTANCE               = "activities/distance"
	RESOURCE_ACTIVITIES_FLOOR                  = "activities/floors"
	RESOURCE_ACTIVITIES_ELEVATION              = "activities/elevation"
	RESOURCE_ACTIVITIES_MINUTES_SEDENTARY      = "activities/minutesSedentary"
	RESOURCE_ACTIVITIES_MINUTES_LIGHTLY_ACTIVE = "activities/minutesLightlyActive"
	RESOURCE_ACTIVITIES_MINUTES_FAIRLY_ACTIVE  = "activities/minutesFairlyActive"
	RESOURCE_ACTIVITIES_MINUTES_VERY_ACTIVE    = "activities/minutesVeryActive"
	RESOURCE_ACTIVITIES_ACTIVITY_CALORIES      = "activities/activityCalories"
	RESOURCE_BODY_WEIGHT                       = "body/weight"
	RESOURCE_BODY_BMI                          = "body/bmi"
	RESOURCE_BODY_FAT                          = "body/fat"
	TOKEN_FILE                                 = "token.json"
)

type Data map[string][]TimeSeriesItem
type TimeSeriesItem struct {
	DateTime string `json:"dateTime"`
	Value    string `json:"value"`
}

type CodeCallbackFunc func(url string) (code string)
type Client struct {
	Client http.Client
}

// Returns a new fitbit.Client Object.
func Connect(conf ClientConfig, codeCallback CodeCallbackFunc) (Client, error) {
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
				return Client{}, err
			}
		} else {
			return Client{}, errors.New("code can't be empty.")
		}
	}

	// Create a TokenSource with the provided token
	// This will automatically refesh the AccessToken when it, thanks to the golang.org/x/oauth2 package.
	tokenSrc := oAuthConfig.TokenSource(oauth2.NoContext, token)
	newToken, err := tokenSrc.Token()
	if err != nil {
		return Client{}, err
	}

	// Initialize new Client with TokenSource
	newClient := oauth2.NewClient(oauth2.NoContext, tokenSrc)

	// Saves the new token, just in case it was refreshed.
	tokenToJSON(newToken)

	return Client{
		Client: *newClient,
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
Minutes Sedentary
****/
func (c *Client) GetMinutesSedentary() ([]TimeSeriesItem, error) {
	return c.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_SEDENTARY, "today", "1d")
}

func (c *Client) GetMinutesSedentaryForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return c.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_SEDENTARY, date, period)
}

/****
Minutes Lightly Active
****/
func (c *Client) GetMinutesLightlyActive() ([]TimeSeriesItem, error) {
	return c.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_LIGHTLY_ACTIVE, "today", "1d")
}

func (c *Client) GetMinutesLightlyActiveForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return c.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_LIGHTLY_ACTIVE, date, period)
}

/****
Minutes Fairly Active
****/
func (c *Client) GetMinutesFairlyActive() ([]TimeSeriesItem, error) {
	return c.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_FAIRLY_ACTIVE, "today", "1d")
}

func (c *Client) GetMinutesFairlyActiveForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return c.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_FAIRLY_ACTIVE, date, period)
}

/****
Minutes Very Active
****/
func (c *Client) GetMinutesVeryActive() ([]TimeSeriesItem, error) {
	return c.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_VERY_ACTIVE, "today", "1d")
}

func (c *Client) GetMinutesVeryActiveForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return c.getTimeSeriesData(RESOURCE_ACTIVITIES_MINUTES_VERY_ACTIVE, date, period)
}

/****
Steps
****/
func (c *Client) GetSteps() ([]TimeSeriesItem, error) {
	return c.GetStepsForDateAndPeriod("today", "1d")
}

func (c *Client) GetStepsForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return c.getTimeSeriesData(RESOURCE_ACTIVITIES_STEPS, date, period)
}

/****
Weight
****/
func (c *Client) GetWeight() ([]TimeSeriesItem, error) {
	return c.GetWeightForDateAndPeriod("today", "1d")
}

func (c *Client) GetWeightForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return c.getTimeSeriesData(RESOURCE_BODY_WEIGHT, date, period)
}

/****
BMI
****/
func (c *Client) GetBMI() ([]TimeSeriesItem, error) {
	return c.GetBMIForDateAndPeriod("today", "1d")
}

func (c *Client) GetBMIForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return c.getTimeSeriesData(RESOURCE_BODY_BMI, date, period)
}

/****
Fat
****/
func (c *Client) GetFat() ([]TimeSeriesItem, error) {
	return c.GetFatForDateAndPeriod("today", "1d")
}

func (c *Client) GetFatForDateAndPeriod(date string, period string) ([]TimeSeriesItem, error) {
	return c.getTimeSeriesData(RESOURCE_BODY_FAT, date, period)
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
