package main

import (
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

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1
			}
		}
	}
}

func main() {
	Env_load()
	api := slack.New(os.Getenv("SLACK_API_TOKEN"))
	os.Exit(run(api))
}
