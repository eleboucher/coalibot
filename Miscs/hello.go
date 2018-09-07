package Miscs

import (
	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func Hello(option string, event *Struct.Message) bool {
	event.API.PostMessage(event.Channel, "Hello <@"+event.User+"> ! powered by go", slack.PostMessageParameters{})
	return true
}
