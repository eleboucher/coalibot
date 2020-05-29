package miscs

import (
	"math/rand"
	"strconv"

	"github.com/eleboucher/coalibot/utils"
	"github.com/nlopes/slack"
)

var tab = make(map[string][]int)

func Roulette(option string, event *utils.Message) bool {
	if tab[event.Channel] == nil || len(tab[event.Channel]) == 0 {
		for i := 0; i < 6; i++ {
			tab[event.Channel] = append(tab[event.Channel], 0)
		}
		tab[event.Channel][rand.Intn(6)] = 1
		utils.PostMsg(event, slack.MsgOptionText("On recharge le revolver!", false))
	}
	var count = 6 - len(tab[event.Channel]) + 1
	if tab[event.Channel][0] == 1 {
		tab[event.Channel] = nil
		utils.PostMsg(event, slack.MsgOptionText("<@"+event.User+">: Bang ( "+strconv.Itoa(count)+" / 6 )", false))
		utils.HandleRouletteStat(event)
	} else {
		tab[event.Channel] = tab[event.Channel][1:]
		utils.PostMsg(event, slack.MsgOptionText("<@"+event.User+">: Click ( "+strconv.Itoa(count)+" / 6 )", false))
	}
	return true
}
