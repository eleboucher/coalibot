package FortyTwo

import (
	"fmt"

	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
	"github.com/nlopes/slack"
)

func Prof(option string, event *Struct.Message) bool {
	user, not_valid := Utils.GetLogin(option, event)
	if not_valid {
		event.API.PostMessage(event.Channel, "invalid login", Struct.SlackParams)
		return false
	}
	data, err := event.FortyTwo.GetUser(user)
	if err != nil {
		event.API.PostMessage(event.Channel, "invalid login", Struct.SlackParams)
		return false
	}

	logtime := getLogtime(user, event.FortyTwo)
	stage := hasDoneIntership(data)
	blocs, _ := event.FortyTwo.GetBlocs(nil)
	coalitions, _ := event.FortyTwo.GetCoalitionsByUser(user, nil)
	color, slug := getCoasRepr(user, event.FortyTwo, blocs, coalitions)
	location := "Unavailable"
	if data.Location != "" {
		location = data.Location
	}
	user = getTitle(data)
	params := Struct.SlackParams
	attachment := slack.Attachment{
		AuthorName: fmt.Sprintf("%s <%s|%s - %s>", slug, "https://profile.intra.42.fr/users/"+user, data.Displayname, user),
		ThumbURL:   data.ImageURL,
		Color:      color,
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Cursus",
				Value: cursusLevels(data.CursusUsers, blocs, coalitions, event.FortyTwo),
				Short: false,
			},
			slack.AttachmentField{
				Title: "Weekly Logtime",
				Value: Utils.FmtDuration(logtime),
				Short: true,
			},
			slack.AttachmentField{
				Title: "Internship",
				Value: stage,
				Short: true,
			},
			slack.AttachmentField{
				Title: "Location",
				Value: location,
				Short: true,
			},
			slack.AttachmentField{
				Title: "Campus",
				Value: getMainCampus(data),
				Short: true,
			},
			slack.AttachmentField{
				Title: "Contact",
				Value: fmt.Sprintf("_%s_ • %s", getNumber(data), data.Email),
				Short: false,
			},
		},
		Footer: "Powered by Coalibot",
	}
	params.Attachments = []slack.Attachment{attachment}
	event.API.PostMessage(event.Channel, "", params)

	return true
}
