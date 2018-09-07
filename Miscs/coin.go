package Miscs

import (
	"math/rand"

	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func Coin(option string, event *Struct.Message) bool {
	if rand.Intn(1) == 0 {
		event.API.PostMessage(event.Channel, "<@"+event.User+">: Heads", slack.PostMessageParameters{})
	} else {
		event.API.PostMessage(event.Channel, "<@"+event.User+">: Tails", slack.PostMessageParameters{})
	}

	return true
}
