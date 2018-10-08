package Miscs

import (
	"github.com/genesixx/coalibot/Struct"
)

func Hello(option string, event *Struct.Message) bool {
	event.API.PostMessage(event.Channel, "Hello <@"+event.User+">!", Struct.SlackParams)
	return true
}
