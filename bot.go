package main

import (
  "github.com/nlopes/slack"
  "Slack_Share_Timeline/timeline"

  "fmt"
  "strings"
)

const botIcon = ":innocent:"

var (
  commands = map[string]string {
    "help": "Display all of the commands.",
  }
)

type Bot struct {
  api *slack.Client
  rtm *slack.RTM
}

func NewBot(token string) *Bot {
  bot := new(Bot)
  bot.api = slack.New(token)
  bot.rtm = bot.api.NewRTM()
  return bot
}

func (b *Bot) handleResponse(user, text, channel string) {
  var cmd string
  var attachment slack.Attachment

  commandArray := strings.Fields(text)
  cmd = commandArray[1]
  if len(commandArray) >= 4 {
    switch cmd {
    case "create":
      // create timeline  object => {OwnerID => Owner channel id, ClientID => }
      roomArray := commandArray[2:]
      attachment = timeline.Create(roomArray)
    case "update":
      // Update timeline object
    case "destroy":
      // Destroy timeline
    }
  } else {
  }

  params := slack.PostMessageParameters{
    Attachments: []slack.Attachment{attachment},
    Username: botName,
    IconEmoji: botIcon,
  }

  _, _, err := b.api.PostMessage(channel, "", params)
  if err != nil {
    b.rtm.SendMessage(b.rtm.NewOutgoingMessage(fmt.Sprintf("Sorry %s is error.... %s", cmd, err), channel))
    return
  }
}
