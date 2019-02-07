package Twitch

import (
	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"

	"github.com/nlopes/slack"
)

func Emotes(option string, event *Struct.Message) bool {
	params := slack.FileUploadParameters{
		File: "emotes/" + option + ".png",
	}
	Utils.PostMsg(event, slack.MsgOptionPostMessageParameters(params))
	return true
}
