package Assos

import (
	"github.com/genesixx/coalibot/Struct"
)

func Event(option string, event *Struct.Message) bool {
	params := Struct.SlackParams
	params.IconURL = "https://bde.student42.fr/img/bde42-logo-1538664197.jpg"
	params.Username = "Undefined Bot"
	// attachment := slack.Attachment{
	// 	Title:     "Galette Des Reines",
	// 	TitleLink: "https://www.facebook.com/events/2634975786518016",
	// 	Text:      "Le BDE vous invite à la galette des reines en cantina le 17 janvier de 16h à 23h. Il y aura de la galette, du cidre ainsi que de la bière.",
	// 	Footer:    "Powered by Coalibot",
	// 	Color:     "#FFD700",
	// 	Fields: []slack.AttachmentField{
	// 		slack.AttachmentField{
	// 			Title: "Date",
	// 			Value: " le 17 janvier de 16h à 23h.",
	// 			Short: false,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Pinte :beer:",
	// 			Value: "3 Euro",
	// 			Short: true,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Demi :wine_glass:",
	// 			Value: "2 Euro",
	// 			Short: true,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Part de galette",
	// 			Value: "1 Euro",
	// 			Short: true,
	// 		},
	// 	},
	// }
	// Utils.PostMsg(event, slack.MsgOptionAttachments(attachment), slack.MsgOptionPostMessageParameters(params))
	// Utils.PostMsg(event, slack.MsgOptionText("Coming soon!", false), slack.MsgOptionPostMessageParameters(params))

	return true
}
