package miscs

import (
	"github.com/eleboucher/coalibot/utils"
	"github.com/slack-go/slack"
)

func Help(option string, event *utils.Message) bool {
	attachment := slack.Attachment{
		Title:     "Coalibot Helper",
		TitleLink: "https://github.com/eleboucher/coalibot",
		Footer:    "Powered by Coalibot",
		Color:     "good",
		Fields: []slack.AttachmentField{
			{
				Title: "alliance",
				Value: "Stat of the Alliance",
				Short: true,
			},
			{
				Title: "addmusic",
				Value: "Add a music to the playlist",
				Short: true,
			},
			{
				Title: "brew",
				Value: "How to install brew",
				Short: true,
			},
			{
				Title: "event",
				Value: "today's event for 42 Paris",
				Short: true,
			},
			{
				Title: "halp",
				Value: "Instruction for iscsi problems",
				Short: true,
			},
			{
				Title: "intra-slack",
				Value: "Instruction to connect slack to the intranet",
				Short: true,
			},
			{
				Title: "logtime",
				Value: "Give the logtime of the given user, get more help `bc logtime --help`",
				Short: true,
			},
			{
				Title: "meteo",
				Value: "Give the weather",
				Short: true,
			},
			{
				Title: "music",
				Value: "Give a random music",
				Short: true,
			},
			{
				Title: "prof",
				Value: "Give info of the student",
				Short: true,
			},
			{
				Title: "score",
				Value: "Give 42 paris coalitions score",
				Short: true,
			},
			{
				Title: "skin",
				Value: "Give slack skins",
				Short: true,
			},
			{
				Title: "source",
				Value: "Give the repo link of Coalibot",
				Short: true,
			},
			{
				Title: "stat",
				Value: "Give stats of the bot",
				Short: true,
			},
			{
				Title: "where",
				Value: "Give the position of the given student",
				Short: true,
			},
			{
				Title: "who",
				Value: "Give the information of who is there",
				Short: true,
			},
			{
				Title: "Bonus",
				Short: false,
			},
			{
				Title: "coin",
				Value: "Heads or tails",
				Short: true,
			},
			{
				Title: "kaamelott",
				Value: "Random Kaamelott quote (FR)",
				Short: true,
			},
			{
				Title: "oss",
				Value: "Random OSS 117 quote (FR)",
				Short: true,
			},
			{
				Title: "mhenni",
				Value: "Random Mohammed Henni quote (FR)",
				Short: true,
			},
			{
				Title: "glados",
				Value: "Random GLaDOS quote (FR)",
				Short: true,
			},
			{
				Title: "roll",
				Value: "random",
				Short: true,
			},
			{
				Title: "roulette",
				Value: "Russian Roulette",
				Short: true,
			},
		},
	}
	utils.PostMsg(event, slack.MsgOptionAttachments(attachment), slack.MsgOptionTS(event.Timestamp))
	return true
}
