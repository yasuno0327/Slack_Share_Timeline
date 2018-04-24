package main

import (
	"Slack_Share_Timeline/client"
	"Slack_Share_Timeline/owner"
	"Slack_Share_Timeline/timeline"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				fmt.Println(ev)
				rtm.SendMessage(rtm.NewOutgoingMessage(ev.Msg.Text, os.Getenv("SLACK_SAMPLE_TIMELINE")))

			case *slack.BotAddedEvent
				// Create Timeline

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1
			}
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	api := slack.New(os.Getenv("SLACK_API_TOKEN"))
	os.Exit(run(api))
}
