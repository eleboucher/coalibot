package FortyTwo

import (
	"fmt"
	"strings"
	"time"

	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
	"github.com/nlopes/slack"
	"gitlab.com/clafoutis/api42"
)

func Prof(option string, event *Struct.Message) bool {
	var user string

	if len(option) > 0 {
		fmt.Println(strings.Split(option, " ")[0])
		user = strings.Split(option, " ")[0]
	} else {
		u, err := event.API.GetUserInfo(event.User)
		if err != nil {
			return false
		}
		user = u.Profile.Email[0:strings.IndexAny(u.Profile.Email, "@")]
	}
	data, err := event.FortyTwo.GetUser(user)
	if err != nil {
		event.API.PostMessage(event.Channel, "invalid login", Struct.SlackParams)
		return false
	}

	coaldata, _ := event.FortyTwo.GetCoalitionsByUser(user, nil)
	var lvlPiscine string
	var level float64
	if data.PoolYear == "2013" || data.PoolYear == "2014" {
		lvlPiscine = fmt.Sprintf("%02.2f", 0.0)
		level = data.CursusUsers[0].Level
	} else if len(data.CursusUsers) == 1 {
		lvlPiscine = fmt.Sprintf("%02.2f", data.CursusUsers[0].Level)
		level = 0.0
	} else {
		lvlPiscine = fmt.Sprintf("%02.2f", data.CursusUsers[1].Level)
		level = data.CursusUsers[0].Level
	}
	y, m, d := time.Now().Date()
	rangeBegin := time.Date(y, m, d, 0, 0, 0, 0, time.Now().Location())
	rangeEnd := rangeBegin.AddDate(0, 0, -7)
	logtime := Utils.IntraLogtime(user, rangeEnd, rangeBegin, event.FortyTwo)
	fmt.Println(data.Projects)
	stage := hasDoneIntership(data)
	color := "#D40000"
	slug := ""
	if len(coaldata) > 0 {
		color = coaldata[0].Color
		slug = ":" + coaldata[0].Slug + ":"
	}
	location := "Hors ligne"
	if data.Location != "" {
		location = data.Location
	}
	params := Struct.SlackParams
	attachment := slack.Attachment{
		Title:     data.Displayname + " - " + user + " " + slug,
		TitleLink: "https://profile.intra.42.fr/users/" + user,
		ThumbURL:  data.ImageURL,
		Color:     color,
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Niveau",
				Value: fmt.Sprintf("%02.2f", level),
				Short: true,
			},
			slack.AttachmentField{
				Title: "Niveau piscine",
				Value: lvlPiscine + " " + data.PoolMonth + " " + data.PoolYear,
				Short: true,
			},
			slack.AttachmentField{
				Title: "Temps de log cette semaine",
				Value: Utils.FmtDuration(logtime),
				Short: true,
			},
			slack.AttachmentField{
				Title: "Stage",
				Value: stage,
				Short: true,
			},
			slack.AttachmentField{
				Title: "Location",
				Value: location,
				Short: true,
			},
		},
		Footer: "Powered by Coalibot",
	}
	params.Attachments = []slack.Attachment{attachment}
	event.API.PostMessage(event.Channel, "", params)

	return true
}

func hasDoneIntership(user *api42.User42) string {
	var stage = ":negative_squared_cross_mark:"
	var indexInternProject = -1
	var indexContractProject = -1
	for k, v := range user.Projects {
		if v.Project.ID == 118 {
			fmt.Println(v)
			indexInternProject = k
		} else if v.Project.ID == 119 {
			indexContractProject = k
		}
	}
	if indexInternProject != -1 && indexContractProject != -1 &&
		user.Projects[indexContractProject].Status == "finished" &&
		*user.Projects[indexInternProject].FinalMark > 0 {

		switch user.Projects[indexInternProject].Status {
		case "finished":
			stage = ":white_check_mark:"
		case "in_progress":
			stage = "clock1"
		}
	}
	return stage
}
