package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/gr4y/fitbit-graphite/lib/fitbit"
	"github.com/gr4y/fitbit-graphite/lib/processor"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "fitbit-graphite"
	app.Usage = "Exports your FitBit Data into your very own graphite instance"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "ClientID,CID",
			Value: "229G69",
			Usage: "OAuth 2.0 Client ID",
		},
		cli.StringFlag{
			Name:  "ClientSecret,CS",
			Value: "cbe3e9792c1c495db76506b2204a834d",
			Usage: "OAuth 2.0 Client Secret",
		},
		cli.StringFlag{
			Name:  "CarbonPrefix,CP",
			Value: "fitness",
			Usage: "Prefix for Carbon",
		},
	}
	app.Action = func(c *cli.Context) {

		clientConfig := fitbit.ClientConfig{
			ClientID:     c.String("ClientID"),
			ClientSecret: c.String("ClientSecret"),
			Scopes:       []string{"activity", "heartrate", "location", "nutrition", "profile", "settings", "sleep", "social", "weight"},
		}

		callbackFunc := func(url string) string {
			fmt.Println("Open the following URL in your browser: ", url)
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

		processors := []processor.Processor{
			processor.ActivitiesProcessor{Activities: client.Activities},
			processor.BodyProcessor{Body: client.Body},
			processor.SleepProcessor{Sleep: client.Sleep},
		}

		var userId string
		profileClient := client.Profile
		profile, err := profileClient.GetProfile()
		if err != nil {
			userId = "-"
		} else {
			userId = profile.User.ID
		}

		var lines []string
		for _, proc := range processors {
			items, err := proc.FetchData("today", "1m")
			if err != nil {
				log.Fatal(err)
			}
			lines = append(lines, items...)
		}

		sendMetrics(lines, fmt.Sprintf("%s.%s", c.String("CarbonPrefix"), userId))

	}
	app.Run(os.Args)
}

func sendMetrics(lines []string, prefix string) {
	for _, line := range lines {
		fmt.Println(fmt.Sprintf("%s.%s", prefix, line))
	}
}
