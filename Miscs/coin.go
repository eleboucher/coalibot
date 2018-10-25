package Miscs

import (
	"math/rand"

	"github.com/genesixx/coalibot/Struct"
)

func Coin(option string, event *Struct.Message) bool {
	if rand.Intn(2) == 0 {
		event.API.PostMessage(event.Channel, "<@"+event.User+">: Heads", Struct.SlackParams)
	} else {
		event.API.PostMessage(event.Channel, "<@"+event.User+">: Tails", Struct.SlackParams)
	}

	return true
}
