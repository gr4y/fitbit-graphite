package fitbit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	RESOURCE_PROFILE = "profile"
)

type Profile struct {
	API Client
}

type UserProfile struct {
	User struct {
		ID string `json:"encodedId"`
	} `json:"user"`
}

func (p *Profile) GetProfile() (UserProfile, error) {
	var profile UserProfile = UserProfile{}
	var err error

	profileUrl := p.API.buildUrl(RESOURCE_PROFILE, "", "")
	resp, err := p.API.Client.Get(profileUrl)
	bytes, err := ioutil.ReadAll(resp.Body)

	if DEBUG {
		fmt.Println("#### DEBUG Output Begin ####")
		fmt.Println(fmt.Sprintf("Request URL: %s\r\n", profileUrl))
		fmt.Println("\r\n### HTTP Request ###")
		fmt.Println(fmt.Sprintf("Request: %+v", resp.Request))
		fmt.Println("\r\n### HTTP Response")
		fmt.Println(fmt.Sprintf("Response Object: %+v", resp))
		fmt.Println(fmt.Sprintf("Response Body: %s", string(bytes)))
		fmt.Println("#### DEBUG Output End ####")
	}
	err = json.Unmarshal(bytes, &profile)

	return profile, err
}
