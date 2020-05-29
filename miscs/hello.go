package miscs

import (
	"github.com/eleboucher/coalibot/utils"
	"github.com/nlopes/slack"
)

func Hello(option string, event *utils.Message) bool {
	utils.PostMsg(event, slack.MsgOptionText("Hello <@"+event.User+">!", false))
	return true
}
