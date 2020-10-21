package fortyTwo

import (
	"fmt"

	"github.com/eleboucher/coalibot/utils"
	"github.com/slack-go/slack"
)

func Prof(option string, event *utils.Message) bool {
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

	logtime := getLogtime(login, event.FortyTwo)
	stage := hasDoneIntership(data)
	blocs, _ := event.FortyTwo.GetBlocs(nil)
	coalitions, _ := event.FortyTwo.GetCoalitionsByUser(login, nil)
	color, slug := getCoasRepr(login, event.FortyTwo, blocs, coalitions)
	location := "Unavailable"
	if data.Location != "" {
		location = data.Location
	}
	user := getTitle(data)
	attachment := slack.Attachment{
		AuthorName: fmt.Sprintf("%s <%s|%s - %s>", slug, "https://profile.intra.42.fr/users/"+login, data.Displayname, user),
		ThumbURL:   fmt.Sprintf("https://cdn.intra.42.fr/users/medium_%s.jpg", login),
		Color:      color,
		Fields: []slack.AttachmentField{
			{
				Title: "Cursus",
				Value: cursusLevels(data.CursusUsers, blocs, coalitions, event.FortyTwo),
				Short: false,
			},
			{
				Title: "Weekly Logtime",
				Value: utils.FmtDuration(logtime),
				Short: true,
			},
			{
				Title: "Internship",
				Value: stage,
				Short: true,
			},
			{
				Title: "Location",
				Value: location,
				Short: true,
			},
			{
				Title: "Campus",
				Value: getMainCampus(data),
				Short: true,
			},
			{
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
