package bars

import (
	"time"

	"github.com/genesixx/coalibot/utils"
	"github.com/nlopes/slack"
)

func Moty(option string, event *utils.Message) bool {
	open := "Fermé !"
	color := "danger"
	if IsMotyOpen() {
		open = "Ouvert !"
		color = "good"
	}
	attachment := slack.Attachment{
		Color: color,
		Title: "Moty",
		Text:  open,
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Horaire",
				Value: "Lundi-Vendredi: 7h30 - 2h, Weekend: 10h - 2h",
			},
			slack.AttachmentField{
				Title: "Ricard :sunny:",
				Value: "3.50€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Gaule :beer:",
				Value: "Demi: 3.10€ Pinte: 4.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Gaule d'Abbaye/Grolsch Blanche :beer:",
				Value: "Demi: 4.00€ Pinte: 5.50€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Vin :wine_glass:",
				Value: "3.50€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Liqueur :tumbler_glass:",
				Value: "3.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Cocktail " + utils.Choice([]string{":cocktail2:", ":cocktail:"}),
				Value: "6.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Alcool fort :whiskey2:",
				Value: "5.50€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Jus de fruit frais :kiwifruit:",
				Value: "3.80€",
				Short: true,
			},
		},
		Footer: "Powered by Coalibot",
	}
	utils.PostMsg(event, slack.MsgOptionAttachments(attachment))
	return true
}

func IsMotyOpen() bool {
	now := time.Now()
	if int(now.Weekday()) > 0 && int(now.Weekday()) < 6 && (now.Hour() == 7 && now.Minute() >= 30 || now.Hour() >= 8) && (now.Hour() <= 24 || now.Hour() < 2) ||
		(int(now.Weekday()) == 6 || int(now.Weekday()) == 0) && (now.Hour() >= 10) && (now.Hour() <= 24 || now.Hour() < 2) {
		return true
	}
	return false
}
