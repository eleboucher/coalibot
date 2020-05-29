package assos

import (
	"github.com/eleboucher/coalibot/utils"
	"github.com/slack-go/slack"
)

func Shop(option string, event *utils.Message) bool {
	var params = utils.SlackParams
	params.IconURL = "https://bde.student42.fr/img/bde42-logo-1538664197.jpg"
	params.Username = "Unicode Bot"
	utils.PostMsg(event, slack.MsgOptionText("Shop du BDE Unicode ! https://bde.student42.fr/", false), slack.MsgOptionPostMessageParameters(params))

	return true
}
