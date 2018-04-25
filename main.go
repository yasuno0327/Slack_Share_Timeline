package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

var (
	botId   string
	botName string
)

func run(api *slack.Client) int {
	bot := NewBot(os.Getenv("SLACK_API_TOKEN"))
	go bot.rtm.ManageConnection()

	for {
		select {
		case msg := <-bot.rtm.IncomingEvents:
			switch ev := msg.Data.(type) {

			case *slack.ConnectedEvent:
				botId = ev.Info.User.ID
				botName = ev.Info.User.Name

			case *slack.MessageEvent:
				user := ev.User
				text := ev.Text
				channel := ev.Channel
				if ev.Type == "message" && strings.HasPrefix(text, "<@"+botId+">") {
					bot.handleResponse(user, text, channel)
				} else {
					bot.handleDefaultMessage(user, text, channel)
				}

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1
			}
		}
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	api := slack.New(os.Getenv("SLACK_API_TOKEN"))
	os.Exit(run(api))
}
