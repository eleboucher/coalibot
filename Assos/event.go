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
		Title:     "A la conquête de Paname !",
		TitleLink: "https://www.facebook.com/events/1041473989381147",
		Text:      "le BDE vous organise une soirée au BAR3 🍻l'occasion idéale pour rencontrer les petits nouveaux qui viennent d'arriver en Avril. On vous prépare des animations et des surprises !",
		Footer:    "Powered by Coalibot",
		Color:     "#009a49",
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Date",
				Value: "le 11 avril de 19h à 02h.",
				Short: false,
			},
			slack.AttachmentField{
				Title: "Biere :beer:",
				Value: "Pinte 4 €",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Shot",
				Value: "3€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Jägerbomb",
				Value: "5€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Adresse : Bar3 (3 Rue de l'Ancienne Comédie, 75006 Paris)",
				Short: true,
			},
		},
	}
	Utils.PostMsg(event, slack.MsgOptionAttachments(attachment), slack.MsgOptionPostMessageParameters(params))
	// Utils.PostMsg(event, slack.MsgOptionText("Coming soon!", false), slack.MsgOptionPostMessageParameters(params))

	return true
}
