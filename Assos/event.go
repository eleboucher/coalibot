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
		Title:     "Soirée de la Saint Patrick !",
		TitleLink: "https://www.facebook.com/events/1041473989381147",
		Text:      "Le BDE Undefined vous propose un événement en E0 le 14 mars pour la St Patrick #bière #pizza! Des pizzas seront en libre service et GRATUIT 🍕. De la musique, de la bière et des cocktails !",
		Footer:    "Powered by Coalibot",
		Color:     "#009a49",
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Date",
				Value: "le 14 mars de 18h à 23h.",
				Short: false,
			},
			slack.AttachmentField{
				Title: "Buchette :beer:",
				Value: "Pinte 3 €/Demi 2 €",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Delirium Red :beer:",
				Value: "Pinte 4,50€/Demi 2,50€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Karmeliet :beer:",
				Value: "Pinte 4,50€/Demi 2,50€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Cuvée des Trolls :beer:",
				Value: "Pinte 4,50€/Demi 2,50€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Cocktails :cocktail:",
				Value: "Avec alcool, à base de vodka: 3€/Sans alcool: 1,50€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Les softs (Coca Cola, Jus de fruits) sont gratuits, comme toujours.",
				Short: true,
			},
		},
	}
	Utils.PostMsg(event, slack.MsgOptionAttachments(attachment), slack.MsgOptionPostMessageParameters(params))
	// Utils.PostMsg(event, slack.MsgOptionText("Coming soon!", false), slack.MsgOptionPostMessageParameters(params))

	return true
}
