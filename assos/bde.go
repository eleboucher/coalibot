package assos

import (
	"strings"

	"github.com/genesixx/coalibot/utils"

	"github.com/nlopes/slack"
)

func Bde(option string, event *utils.Message) bool {
	var params = utils.SlackParams
	params.IconURL = "https://scontent-cdt1-1.xx.fbcdn.net/v/t1.0-9/72837094_1434269160066611_1550685959262044160_n.png?_nc_cat=100&_nc_oc=AQnnBrO0tgIxEsYONp9dIR-9bZ830RKV3jI-xNxw8dvcP0qVBrR7ttej4JeI9fLKk9s&_nc_ht=scontent-cdt1-1.xx&oh=841847d7c60b39ed54b28057e1ae5770&oe=5E170AA7"
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
