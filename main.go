package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

var (
	botId string
	botName string
)

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	bot := NewBot(os.Getenv("SLACK_API_TOKEN"))
	go bot.rtm.ManageConnection()

	for {
		select {
		case msg := <-bot.rtm.IncomingEvents:
			switch ev := msg.Data.(type) {

			case *slack.ConnectedEvent:
				botId = os.Getenv("SLACK_BOT_ID")
				botName = "timeline"

			case *slack.MessageEvent:
				user := ev.User
				text := ev.Text
				channel := ev.Channel
				fmt.Println(text)
				fmt.Println(botId)
				if ev.Type == "message" && strings.HasPrefix(text, "<@"+botId+">") {
					bot.handleResponse(user, text, channel)
				}
				rtm.SendMessage(rtm.NewOutgoingMessage(ev.Msg.Text, os.Getenv("SLACK_SAMPLE_TIMELINE")))

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
