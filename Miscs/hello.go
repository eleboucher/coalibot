package Miscs

import (
	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
	"github.com/nlopes/slack"
)

func Hello(option string, event *Struct.Message) bool {
	Utils.PostMsg(event, slack.MsgOptionText("Hello <@"+event.User+">!", false))
	return true
}
