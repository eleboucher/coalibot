package assos

import (
	"strings"

	"github.com/eleboucher/coalibot/utils"

	"github.com/nlopes/slack"
)

func Bde(option string, event *utils.Message) bool {
	var params = utils.SlackParams
	params.IconURL = "https://i.imgur.com/zLpfNV7.png"
	params.Username = "Unicode Bot"
	if option == "" {
		utils.PostMsg(event, slack.MsgOptionText("Pour avoir accès au shop tapez `!bde shop`, pour avoir plus d'info sur les events tapez `!bde event`!", false), slack.MsgOptionPostMessageParameters(params))
		return true
	}
	switch strings.ToLower(strings.Split(option, " ")[0]) {
	case "shop":
		return Shop(option, event)
	case "event":
		return Event(option, event)
	}
	return true
}
