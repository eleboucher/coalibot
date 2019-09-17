package fortyTwo

import (
	"strings"
	"time"

	"github.com/genesixx/coalibot/utils"
	"github.com/nlopes/slack"

	"gitlab.com/clafoutis/api42"
)

func Who(option string, event *utils.Message) bool {
	if option == "" || len(strings.Split(option, " ")) > 1 || len(strings.Split(option, " ")) == 0 {
		utils.PostMsg(event, slack.MsgOptionText("Prend une place en parametre", false))
		return false
	}
	var place = strings.Split(option, " ")[0]
	if place[0] == '!' || place[0] == '?' {
		return false
	}
	params := api42.NewParameter()
	params.AddFilter("host", place)
	data, err := event.FortyTwo.GetLocations(params)
	if err != nil {
		return false
	}

	if len(data) == 0 {
		utils.PostMsg(event, slack.MsgOptionText("Place *"+place+"* vide.", false))
	} else if data[0].EndAt == nil {
		utils.PostMsg(event, slack.MsgOptionText("*"+data[0].User.Login+"* est à la place *"+place+"*", false))

	} else {
		var diff = time.Now().Sub(*data[0].EndAt)
		utils.PostMsg(event, slack.MsgOptionText("Place *"+place+"* vide, ancien utilisateur *"+data[0].User.Login+"* deconnecté depuis *"+utils.PrettyDurationPrinting(diff)+"*", false))
	}
	return true
}
