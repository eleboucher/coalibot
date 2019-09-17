package assos

import (
	"github.com/genesixx/coalibot/utils"
	"github.com/nlopes/slack"
)

func Shop(option string, event *utils.Message) bool {
	var params = utils.SlackParams
	params.IconURL = "https://bde.student42.fr/img/bde42-logo-1538664197.jpg"
	params.Username = "Undefined Bot"
	utils.PostMsg(event, slack.MsgOptionText("Shop du BDE Undefined ! https://bde.student42.fr/", false), slack.MsgOptionPostMessageParameters(params))

	return true
}
