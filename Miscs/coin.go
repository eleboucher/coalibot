package Miscs

import (
	"math/rand"

	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
	"github.com/nlopes/slack"
)

func Coin(option string, event *Struct.Message) bool {
	if rand.Intn(2) == 0 {
		Utils.PostMsg(event, slack.MsgOptionText("<@"+event.User+">: Face", false))
	} else {
		Utils.PostMsg(event, slack.MsgOptionText("<@"+event.User+">: Pile", false))
	}

	return true
}
