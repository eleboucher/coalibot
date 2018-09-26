package Bars

import (
	"fmt"
	"time"

	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func Spritz(option string, event *Struct.Message) bool {
	now := time.Now()
	open := "Fermé !"
	color := "danger"
	fmt.Println(int(now.Weekday()), now.Hour())
	if int(now.Weekday()) > 0 && int(now.Weekday()) < 6 && now.Hour() >= 11 && now.Hour() <= 22 {
		open = "Ouvert !"
		color = "good"
	}
	params := Struct.SlackParams
	attachment := slack.Attachment{
		Color: color,
		Title: "Spritz",
		Text:  open + "\n*Horaire:* 11-22h tous les jours sauf weekend et jours fériés!",
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Ricard",
				Value: "2.70e",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Kronenbourg",
				Value: "Demi: 2.70e Pinte: 5.00e",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Carlsberg/Grimbergen",
				Value: "Demi: 3.70e Pinte: 7.00e",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Liqueur/Shot",
				Value: "2.00e",
				Short: true,
			},
		},
		Footer: "Powered by Coalibot",
	}
	params.Attachments = []slack.Attachment{attachment}
	event.API.PostMessage(event.Channel, "", params)
	return true
}
