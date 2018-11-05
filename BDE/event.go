package BDE

import (
	"github.com/genesixx/coalibot/Struct"
)

func Event(option string, event *Struct.Message) bool {
	params := Struct.SlackParams
	params.IconURL = "https://bde.student42.fr/img/bde42-logo-1538664197.jpg"
	params.Username = "Undefined Bot"
	// attachment := slack.Attachment{
	// 	Title:     "Blood Horror Party",
	// 	TitleLink: "https://bde.student42.fr",
	// 	Text:      "Le BDE Undefined vous prépare une soirée d'Halloween à 42 organisée avec le BDE Wolf (de l'IDRAC) ainsi que le BDE Wastis (de l'école W).",
	// 	Footer:    "Powered by Coalibot",
	// 	ThumbURL:  "https://www.helloasso.com/assets/img/photos/evenements/42%20blood%20horror%20party-min%201%20-71313069e2e74a1ebe82780b6b06663a.png",
	// 	Color:     "#000000",
	// 	Fields: []slack.AttachmentField{
	// 		slack.AttachmentField{
	// 			Title: "Date",
	// 			Value: "Vendredi 2 novembre 2018 de 21:00 à 05:00, Entrée jusqu'à 1h.",
	// 			Short: false,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Entrée",
	// 			Value: "10€ pour les students, 12€ pour les externes",
	// 			Short: false,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Vestiaire",
	// 			Value: "Vestiaire gratuit pour les personnes déguisées, 1€ sinon.",
	// 			Short: false,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Tarifs des boissons:",
	// 			Value: "1 Token = 1€",
	// 			Short: false,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Bière :beer:",
	// 			Value: "Demi 2 token/Pinte 3 token",
	// 			Short: true,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Alcool fort :whiskey2:",
	// 			Value: "3 token",
	// 			Short: true,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Shot :tumbler_glass:",
	// 			Value: "2 token",
	// 			Short: true,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Teq Paf :whiskey2:",
	// 			Value: "2 token",
	// 			Short: true,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "+1 Token pour l'Energy Drink",
	// 			Short: true,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Soft Gratuit",
	// 			Short: true,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Happy Hour jusqu'a 23h",
	// 			Value: "-1 token sur les Bières",
	// 			Short: true,
	// 		},
	// 	},
	// }
	// params.Attachments = []slack.Attachment{attachment}
	event.API.PostMessage(event.Channel, "Comming soon!", params)
	return true
}
