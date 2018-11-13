package Assos

import (
	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func Shop(option string, event *Struct.Message) bool {
	var params = Struct.SlackParams
	params.IconURL = "https://bde.student42.fr/img/bde42-logo-1538664197.jpg"
	params.Username = "Undefined Bot"
	event.API.PostMessage(event.Channel, slack.MsgOptionText("Shop du BDE Undefined ! https://bde.student42.fr/", false), slack.MsgOptionPostMessageParameters(params))

	return true
}
