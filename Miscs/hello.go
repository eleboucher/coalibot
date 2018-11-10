package Miscs

import (
	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func Hello(option string, event *Struct.Message) bool {
	event.API.PostMessage(event.Channel, slack.MsgOptionText("Hello <@"+event.User+">!", false))
	return true
}
