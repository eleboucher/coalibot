package Miscs

import (
	"math/rand"

	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func Coin(option string, event *Struct.Message) bool {
	if rand.Intn(2) == 0 {
		event.API.PostMessage(event.Channel, slack.MsgOptionText("<@"+event.User+">: Face", false))
	} else {
		event.API.PostMessage(event.Channel, slack.MsgOptionText("<@"+event.User+">: Pile", false))
	}

	return true
}
