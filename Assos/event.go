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
		Title: "Soirée de la Saint Valentin",
		// TitleLink: "https://www.facebook.com/events/2634975786518016",
		Text:   "Le BDE vous invite à la Saint-Valentin le 14 février de 18h à 23h.Au programme vente de barbe à papa, de pommes d'amour faites maison, de bières ainsi que de cocktails avec et sans alcool. Comme d'habitude les softs seront gratuits !",
		Footer: "Powered by Coalibot",
		Color:  "#FF69B4",
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Date",
				Value: "le 14 fevrier de 18h à 23h.",
				Short: false,
			},
			slack.AttachmentField{
				Title: "Pinte :beer: ou Cocktail avec-alcool :cocktail:",
				Value: "3 Euro",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Pinte :beer: ou Cocktail sans-alcool :cocktail:",
				Value: "2 Euro",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Barbe a papa :beard::unicorn-pink:",
				Value: "2 Euro",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Pomme d'amour :heart:",
				Value: "2 Euro",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Pas d'alcool pour les piscineux",
				Short: true,
			},
		},
	}
	Utils.PostMsg(event, slack.MsgOptionAttachments(attachment), slack.MsgOptionPostMessageParameters(params))
	// Utils.PostMsg(event, slack.MsgOptionText("Coming soon!", false), slack.MsgOptionPostMessageParameters(params))

	return true
}
