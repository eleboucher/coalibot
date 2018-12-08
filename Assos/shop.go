package Assos

import (
	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
	"github.com/nlopes/slack"
)

func Shop(option string, event *Struct.Message) bool {
	var params = Struct.SlackParams
	params.IconURL = "https://bde.student42.fr/img/bde42-logo-1538664197.jpg"
	params.Username = "Undefined Bot"
	Utils.PostMsg(event, slack.MsgOptionText("Shop du BDE Undefined ! https://bde.student42.fr/", false), slack.MsgOptionPostMessageParameters(params))

	return true
}
