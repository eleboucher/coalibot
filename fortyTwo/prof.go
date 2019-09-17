package fortyTwo

import (
	"fmt"

	"github.com/genesixx/coalibot/utils"
	"github.com/nlopes/slack"
)

func Prof(option string, event *utils.Message) bool {
	user, notValid := utils.GetLogin(option, event)
	if notValid {
		utils.PostMsg(event, slack.MsgOptionText("invalid login", false))
		return false
	}
	data, err := event.FortyTwo.GetUser(user)
	if err != nil {
		utils.PostMsg(event, slack.MsgOptionText("invalid login", false))
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
				Value: utils.FmtDuration(logtime),
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
				Value: fmt.Sprintf("%s", data.Email),
				Short: false,
			},
		},
		Footer: "Powered by Coalibot",
	}
	utils.PostMsg(event, slack.MsgOptionAttachments(attachment))

	return true
}
