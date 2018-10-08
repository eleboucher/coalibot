package Miscs

import (
	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func Help(option string, event *Struct.Message) bool {
	params := Struct.SlackParams
	params.ThreadTimestamp = event.Timestamp
	attachment := slack.Attachment{
		Title:     "Coalibot Helper",
		TitleLink: "https://github.com/genesixx/coalibot",
		Footer:    "Powered by Coalibot",
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "alliance",
				Value: "Stats de l'Alliance",
				Short: true,
			},
			slack.AttachmentField{
				Title: "addmusic",
				Value: "ajoute une musique a la playlist",
				Short: true,
			},
			slack.AttachmentField{
				Title: "brew",
				Value: "Commande pour installer brew",
				Short: true,
			},
			slack.AttachmentField{
				Title: "event",
				Value: "Event de ce jour",
				Short: true,
			},
			slack.AttachmentField{
				Title: "halp",
				Value: "Instructions pour les problèmes liés à iscsi",
				Short: true,
			},
			slack.AttachmentField{
				Title: "logtime",
				Value: "Pour plus d'info `bc logtime --help`",
				Short: true,
			},
			slack.AttachmentField{
				Title: "meteo",
				Value: "Donne le temps qu'il fait à 42",
				Short: true,
			},
			slack.AttachmentField{
				Title: "music",
				Value: "Donne une musique aléatoire",
				Short: true,
			},
			slack.AttachmentField{
				Title: "prof",
				Value: " Donne les infos de l'étudiant",
				Short: true,
			},
			slack.AttachmentField{
				Title: "score",
				Value: "Donne le score des coalitions",
				Short: true,
			},
			slack.AttachmentField{
				Title: "source",
				Value: "Donne le repo de Coalibot",
				Short: true,
			},
			slack.AttachmentField{
				Title: "stat",
				Value: "Stat du bot en direct",
				Short: true,
			},
			slack.AttachmentField{
				Title: "where",
				Value: "Donne la position de l'étudiant",
				Short: true,
			},
			slack.AttachmentField{
				Title: "who",
				Value: "Donne qui est à cette place",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Bonus",
				Short: false,
			},
			slack.AttachmentField{
				Title: "coin",
				Value: "pile ou face",
				Short: true,
			},
			slack.AttachmentField{
				Title: "kaamelott",
				Value: "Citation aléatoire de Kaamelott",
				Short: true,
			},
			slack.AttachmentField{
				Title: "oss",
				Value: "Citation aléatoire de OSS 117",
				Short: true,
			},
			slack.AttachmentField{
				Title: "mhenni",
				Value: "Citation aléatoire de Mohammed Henni",
				Short: true,
			},
			slack.AttachmentField{
				Title: "roll",
				Value: "random",
				Short: true,
			},
			slack.AttachmentField{
				Title: "roulette",
				Value: "Roulette Russe",
				Short: true,
			},
		},
	}
	params.Attachments = []slack.Attachment{attachment}
	event.API.PostMessage(event.Channel, "", params)
	return true
}
