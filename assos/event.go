package assos

import (
	"github.com/genesixx/coalibot/utils"
	"github.com/nlopes/slack"
)

func Event(option string, event *utils.Message) bool {
	params := utils.SlackParams
	params.IconURL = "https://scontent-cdt1-1.xx.fbcdn.net/v/t1.0-9/72837094_1434269160066611_1550685959262044160_n.png?_nc_cat=100&_nc_oc=AQnnBrO0tgIxEsYONp9dIR-9bZ830RKV3jI-xNxw8dvcP0qVBrR7ttej4JeI9fLKk9s&_nc_ht=scontent-cdt1-1.xx&oh=841847d7c60b39ed54b28057e1ae5770&oe=5E170AA7"
	params.Username = "Unicode Bot"
	attachment := slack.Attachment{
		Title:     "Noche de los Muertos",
		TitleLink: "https://www.facebook.com/events/505707719983028",
		Text:      "Le 08/11 le BDE UNICODE vous organise une soirée dans 42. Theme: Dias de los Muertos/ Neon 👻🇲🇽. (1 token == 1 euro)",
		Footer:    "Powered by Coalibot",
		Color:     "#ff32bb",
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Date",
				Value: "vendredi 8 Novembre de 21:00 à 04:00",
				Short: false,
			},
			slack.AttachmentField{
				Title: "Tarif",
				Value: "Prévente: 10€ (adhérent), 12€ (non adhérent).\nSur place: 15€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Bière blonde :beer:",
				Value: "3 tokens la pinte, 2 tokens le demi",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Bière rouge :beer:",
				Value: "4 tokens la pinte, 3 tokens le demi",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Cocktails :cocktail:",
				Value: "5 tokens",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Alcools forts + Softs (Whisky coca, vodka orange, etc...) :whisky:",
				Value: "12 €",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Vestiaire gratuit pour les personnes déguisées et pour les adherents BDE, sinon 1 Un peso (token)",
				Short: true,
			},
		},
	}
	utils.PostMsg(event, slack.MsgOptionAttachments(attachment), slack.MsgOptionPostMessageParameters(params))
	// utils.PostMsg(event, slack.MsgOptionText("Coming soon!", false), slack.MsgOptionPostMessageParameters(params))

	return true
}
