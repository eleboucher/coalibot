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
		Text:      "Le BDE undefined 42, en partenariat avec CPP et 42Green vous invitent ce jeudi 23 mai 2019 à venir partager avec nous une journée récréative, qui se terminera avec un barbecue offert à tous les étudiants (vegan compatible) !",
		Footer:    "Powered by Coalibot",
		Color:     "#009a49",
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Date",
				Value: "le jeudi 23 mai 2019 de 12:00 à 23:00",
				Short: false,
			},
			slack.AttachmentField{
				Title: "Programme",
				Value: "",
				Short: false,
			},
			slack.AttachmentField{
				Title: "Dés 10h, et jusque 14h environ, tournoi de pétanque.",
				Value: "Inscriptions ci-dessous : https://doodle.com/poll/eda7pqvb2khyk8se Puis pétanque libre à partir de 14h, jusqu'à la tombée de la nuit !",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Dés 14h, et jusque 23h, tournoi de cup pong.",
				Value: "Inscriptions ci-dessous : 14h-18h : https://doodle.com/poll/puftq63gkbsu6a8r 18h-22h : https://doodle.com/poll/97ynm78u4zv27f8w",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Dés 19h, barbecue gratuit, sans inscriptions.",
				Value: "",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Buvette",
				Value: "",
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
