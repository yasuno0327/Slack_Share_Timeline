package main

import (
	"Slack_Share_Timeline/timeline"
	"github.com/nlopes/slack"

	"fmt"
	"strings"
)

const botIcon = ":unitus:"

var (
	commands = map[string]string{
		"help": "Display all of the commands.",
		"create": "Create the timeline. Example => @timeline create #owner_timeline #client_timelines [#clinet_timelines...]",
		"update": "Update the timeline. ※ Unimplemented",
		"delete": "Delete the timeline. ※ Unimplemented",
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

func (b *Bot) handleDefaultMessage(user, text, channel string) {
	//var attachment slack.Attachment
	attachment, owner := timeline.HandleMessageResponse(user, text, channel)
	fmt.Println(attachment, owner)

	if len(attachment.Text) == 0 || len(owner) == 0 {
		return
	}

	params := slack.PostMessageParameters{
		Attachments: []slack.Attachment{attachment},
		Username: botName,
		IconEmoji: botIcon,
	}

	_, _, err := b.api.PostMessage(owner, "", params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (b *Bot) handleResponse(user, text, channel string) {
	var cmd string
	var attachment slack.Attachment

	commandArray := strings.Fields(text)
	cmd = commandArray[1]
	if len(commandArray) >= 1 {
		switch cmd {
		case "create":
			// create timeline  object => {OwnerID => Owner channel id, ClientID => }
			roomArray := commandArray[2:]
			attachment = timeline.Create(roomArray)
		case "help":
			attachment = b.help()
		case "update":
			// Update timeline object
		case "destroy":
			// Destroy timeline
		}
	}

	params := slack.PostMessageParameters{
		Attachments: []slack.Attachment{attachment},
		Username:    botName,
		IconEmoji:   botIcon,
	}

	_, _, err := b.api.PostMessage(channel, "", params)
	if err != nil {
		b.rtm.SendMessage(b.rtm.NewOutgoingMessage(fmt.Sprintf("Sorry %s is error.... %s", cmd, err), channel))
		return
	}
}

func (b *Bot) help() (attachment slack.Attachment) {
	fields := make([]slack.AttachmentField, 0)

	for name, command := range commands {
		fields = append(fields, slack.AttachmentField{
			Title: "@"+botName+" " + name,
			Value: command,
		})
	}

	attachment = slack.Attachment{
		Pretext: botName + " Command List",
		Color: "#B733FF",
		Fields: fields,
	}
	return attachment
}
