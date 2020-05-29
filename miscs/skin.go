package miscs

import (
	"strings"

	"github.com/eleboucher/coalibot/utils"
	"github.com/nlopes/slack"
)

func Skin(option string, event *utils.Message) bool {
	switch strings.ToLower(strings.Split(option, " ")[0]) {
	case "":
		utils.PostMsg(event, slack.MsgOptionText("!skin [alliance | assembly | order | federation | 42 | ricard]", false))
		return true
	case "alliance":
		utils.PostMsg(event, slack.MsgOptionText("#1e2124,#2C3849,#33c47f,#ffffff,#1d3b2f,#ffffff,#33c47f,#c90828", false))
		return true
	case "assembly":
		utils.PostMsg(event, slack.MsgOptionText("#1E2124,#2C3849,#a061d1,#ffffff,#531582,#ffffff,#a061d1,#c90828", false))
		return true
	case "order":
		utils.PostMsg(event, slack.MsgOptionText("#1e2124,#2C3849,#FF6950,#000000,#4a231e,#ffffff,#FF6950,#c90828", false))
		return true
	case "federation":
		utils.PostMsg(event, slack.MsgOptionText("#1e2124,#2C3849,#4180DB,#ffffff,#254a7d,#ffffff,#4180DB,#c90828", false))
		return true
	case "ricard":
		utils.PostMsg(event, slack.MsgOptionText("#004684,#395882,#FFD300,#000000,#594e14,#ffffff,#FFD300,#c90828", false))
		return true
	case "42":
		utils.PostMsg(event, slack.MsgOptionText("#1e2124,#2C3849,#00BABC,#ffffff,#4A5664,#e3e3e3,#00ffc4,#c90828", false))
		return true
	}
	return false
}
