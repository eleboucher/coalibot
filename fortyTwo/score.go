package fortyTwo

import (
	"sort"
	"strconv"

	"github.com/genesixx/coalibot/utils"

	"github.com/nlopes/slack"
)

func Score(option string, event *utils.Message) bool {
	login, notValid := utils.GetLogin(option, event)
	if notValid {
		utils.PostMsg(event, slack.MsgOptionText("invalid login", false))
		return false
	}
	data, err := event.FortyTwo.GetUser(login)
	if err != nil {
		utils.PostMsg(event, slack.MsgOptionText("invalid login", false))
		return false
	}
	coalitions, err := getCoalitionsByUser(data, event.FortyTwo)

	msgRef := slack.NewRefToMessage(event.Channel, event.Timestamp)
	go event.API.AddReaction(":the-alliance:", msgRef)
	if err != nil {
		return false
	}
	sort.Slice(coalitions, func(i, j int) bool {
		return coalitions[i].Score > coalitions[j].Score
	})
	var fields []slack.AttachmentField
	for i := 0; i < len(coalitions); i++ {
		var score string
		if i > 0 {
			score = strconv.Itoa(coalitions[i].Score) + " (-" + strconv.Itoa(coalitions[0].Score-coalitions[i].Score) + ")"
		} else {
			score = strconv.Itoa(coalitions[i].Score)
		}
		fields = append(fields, slack.AttachmentField{
			Title: strconv.Itoa(i+1) + "# " + coalitions[i].Name,
			Value: score,
			Short: true,
		})
	}
	attachment := slack.Attachment{
		Color:      coalitions[0].Color,
		AuthorLink: "https://profile.intra.42.fr/blocs/1/coalitions",
		Fields:     fields,
		Footer:     "Powered by Coalibot",
	}
	utils.PostMsg(event, slack.MsgOptionAttachments(attachment))
	return true
}
