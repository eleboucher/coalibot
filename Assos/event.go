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
		Title:     "Quarante d'Œuf Easter Party",
		TitleLink: "https://www.facebook.com/events/1011459799052089/",
		Text:      "Le BDE Undefined vous propose un événement en E0 le 24 Avril pour Pâques 🥚🐇. Des fontaines à chocolat seront en libre service et GRATUITE 🍫.",
		Footer:    "Powered by Coalibot",
		Color:     "#009a49",
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Date",
				Value: "le 24 avril 2019 de 18:00 à 23:00",
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
