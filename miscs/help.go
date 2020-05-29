package miscs

import (
	"github.com/eleboucher/coalibot/utils"
	"github.com/nlopes/slack"
)

func Help(option string, event *utils.Message) bool {
	attachment := slack.Attachment{
		Title:     "Coalibot Helper",
		TitleLink: "https://github.com/eleboucher/coalibot",
		Footer:    "Powered by Coalibot",
		Color:     "good",
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "alliance",
				Value: "Stat of the Alliance",
				Short: true,
			},
			slack.AttachmentField{
				Title: "addmusic",
				Value: "Add a music to the playlist",
				Short: true,
			},
			slack.AttachmentField{
				Title: "brew",
				Value: "How to install brew",
				Short: true,
			},
			slack.AttachmentField{
				Title: "event",
				Value: "today's event for 42 Paris",
				Short: true,
			},
			slack.AttachmentField{
				Title: "halp",
				Value: "Instruction for iscsi problems",
				Short: true,
			},
			slack.AttachmentField{
				Title: "intra-slack",
				Value: "Instruction to connect slack to the intranet",
				Short: true,
			},
			slack.AttachmentField{
				Title: "logtime",
				Value: "Give the logtime of the given user, get more help `bc logtime --help`",
				Short: true,
			},
			slack.AttachmentField{
				Title: "meteo",
				Value: "Give the weather",
				Short: true,
			},
			slack.AttachmentField{
				Title: "music",
				Value: "Give a random music",
				Short: true,
			},
			slack.AttachmentField{
				Title: "prof",
				Value: "Give info of the student",
				Short: true,
			},
			slack.AttachmentField{
				Title: "score",
				Value: "Give 42 paris coalitions score",
				Short: true,
			},
			slack.AttachmentField{
				Title: "skin",
				Value: "Give slack skins",
				Short: true,
			},
			slack.AttachmentField{
				Title: "source",
				Value: "Give the repo link of Coalibot",
				Short: true,
			},
			slack.AttachmentField{
				Title: "stat",
				Value: "Give stats of the bot",
				Short: true,
			},
			slack.AttachmentField{
				Title: "where",
				Value: "Give the position of the given student",
				Short: true,
			},
			slack.AttachmentField{
				Title: "who",
				Value: "Give the information of who is there",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Bonus",
				Short: false,
			},
			slack.AttachmentField{
				Title: "coin",
				Value: "Heads or tails",
				Short: true,
			},
			slack.AttachmentField{
				Title: "kaamelott",
				Value: "Random Kaamelott quote (FR)",
				Short: true,
			},
			slack.AttachmentField{
				Title: "oss",
				Value: "Random OSS 117 quote (FR)",
				Short: true,
			},
			slack.AttachmentField{
				Title: "mhenni",
				Value: "Random Mohammed Henni quote (FR)",
				Short: true,
			},
			slack.AttachmentField{
				Title: "glados",
				Value: "Random GLaDOS quote (FR)",
				Short: true,
			},
			slack.AttachmentField{
				Title: "roll",
				Value: "random",
				Short: true,
			},
			slack.AttachmentField{
				Title: "roulette",
				Value: "Russian Roulette",
				Short: true,
			},
		},
	}
	utils.PostMsg(event, slack.MsgOptionAttachments(attachment), slack.MsgOptionTS(event.Timestamp))
	return true
}
