package miscs

import (
	"math/rand"

	"github.com/eleboucher/coalibot/utils"
	"github.com/nlopes/slack"
)

func Coin(option string, event *utils.Message) bool {
	if rand.Intn(2) == 0 {
		utils.PostMsg(event, slack.MsgOptionText("<@"+event.User+">: Face", false))
	} else {
		utils.PostMsg(event, slack.MsgOptionText("<@"+event.User+">: Pile", false))
	}

	return true
}
