package main

import (
	"fmt"
	"github.com/gr4y/fitbit-graphite/lib/fitbit"
	"github.com/gr4y/fitbit-graphite/lib/plugins"
	"log"
)

func main() {

	clientConfig := fitbit.ClientConfig{
		ClientID:     "229G69",
		ClientSecret: "cbe3e9792c1c495db76506b2204a834d",
		Scopes:       []string{"activity", "heartrate", "location", "nutrition", "profile", "settings", "sleep", "social", "weight"},
	}

	callbackFunc := func(url string) string {
		fmt.Println(url)
		var code string
		_, err := fmt.Scan(&code)
		if err != nil {
			return ""
		}
		return code
	}

	// Connect to FitBit
	client, err := fitbit.Connect(clientConfig, callbackFunc)
	if err != nil {
		log.Fatal(err)
	}

	plugins := []plugins.Plugin{
		plugins.StepsPlugin{Client: client},
		plugins.WeightPlugin{Client: client},
		//		plugins.CaloriesPlugin{Client: client},
	}

	for _, plugin := range plugins {
		plugin.FetchData("today", "1m")
	}

}
