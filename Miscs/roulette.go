package Miscs

import (
	"math/rand"
	"strconv"

	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
)

var tab = make(map[string][]int)

func Roulette(option string, event *Struct.Message) bool {
	if tab[event.Channel] == nil || len(tab[event.Channel]) == 0 {
		for i := 0; i < 6; i++ {
			tab[event.Channel] = append(tab[event.Channel], 0)
		}
		tab[event.Channel][rand.Intn(6)] = 1
		event.API.PostMessage(event.Channel, "On recharge le revolver!", Struct.SlackParams)
	}
	var count = 6 - len(tab[event.Channel]) + 1
	if tab[event.Channel][0] == 1 {
		tab[event.Channel] = nil
		event.API.PostMessage(event.Channel, "<@"+event.User+">: Bang ( "+strconv.Itoa(count)+" / 6 )", Struct.SlackParams)
		Utils.HandleRouletteStat(event)
	} else {
		tab[event.Channel] = tab[event.Channel][1:]
		event.API.PostMessage(event.Channel, "<@"+event.User+">: Click ( "+strconv.Itoa(count)+" / 6 )", Struct.SlackParams)
	}
	return true
}
