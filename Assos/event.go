package Assos

import (
	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
	"github.com/nlopes/slack"
)

func Event(option string, event *Struct.Message) bool {
	params := Struct.SlackParams
	params.IconURL = "https://bde.student42.fr/img/bde42-logo-1538664197.jpg"
	params.Username = "Undefined Bot"
	attachment := slack.Attachment{
		Title:     "Apéro de Noël",
		TitleLink: "https://www.facebook.com/events/591779434601258/",
		Text:      "Pour finir cette année en mode chill, le BDE vous propose un apéro en e0 le jeudi 20 décembre de 16h à 23h.",
		Footer:    "Powered by Coalibot",
		Color:     "#d42426",
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Date",
				Value: "Jeudi 20 decembre 2018 de 16:00 à 23:00.",
				Short: false,
			},
			slack.AttachmentField{
				Title: "Bière :beer:",
				Value: "Pinte 3 Euro",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Vin Chaud :wine_glass:",
				Value: "1  Euro",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Beaucoup de choses à grignoter [GRATUIT]",
				Short: false,
			},
		},
	}
	Utils.PostMsg(event, slack.MsgOptionAttachments(attachment), slack.MsgOptionPostMessageParameters(params))
	return true
}
