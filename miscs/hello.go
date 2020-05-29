package miscs

import (
	"github.com/eleboucher/coalibot/utils"
	"github.com/slack-go/slack"
)

func Hello(option string, event *utils.Message) bool {
	utils.PostMsg(event, slack.MsgOptionText("Hello <@"+event.User+">!", false))
	return true
}
