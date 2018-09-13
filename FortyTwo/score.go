package FortyTwo

import (
	"sort"
	"strconv"

	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
	"gitlab.com/clafoutis/api42"
)

func Score(option string, event *Struct.Message) bool {
	var coalitions []api42.Coalition42
	_, err := event.FortyTwo.Get("/v2/blocs/1/coalitions", nil, &coalitions)
	if err != nil {
		return false
	}
	sort.Slice(coalitions, func(i, j int) bool {
		return coalitions[i].Score > coalitions[j].Score
	})
	var fields []slack.AttachmentField
	for i := 0; i < len(coalitions); i++ {
		score := strconv.Itoa(coalitions[i].Score)
		fields = append(fields, slack.AttachmentField{
			Title: coalitions[i].Name,
			Value: score,
			Short: true,
		})
	}
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Color:      coalitions[0].Color,
		AuthorLink: "https://profile.intra.42.fr/blocs/1/coalitions",
		Fields:     fields,
		Footer:     "Powered by Coalibot",
	}
	params.Attachments = []slack.Attachment{attachment}
	event.API.PostMessage(event.Channel, "", params)
	return true
}
