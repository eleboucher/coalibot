package Miscs

import (
	"strings"

	"github.com/genesixx/coalibot/Struct"
)

func Skin(option string, event *Struct.Message) bool {
	switch strings.ToLower(option) {
	case "alliance":
		event.API.PostMessage(event.Channel, "#1e2124,#2C3849,#33c47f,#ffffff,#1d3b2f,#ffffff,#33c47f,#c90828", Struct.SlackParams)
		return true
	case "assembly":
		event.API.PostMessage(event.Channel, "#1E2124,#2C3849,#a061d1,#ffffff,#531582,#ffffff,#a061d1,#c90828", Struct.SlackParams)
		return true
	case "order":
		event.API.PostMessage(event.Channel, "#1e2124,#2C3849,#FF6950,#000000,#4a231e,#ffffff,#FF6950,#c90828", Struct.SlackParams)
		return true
	case "federation":
		event.API.PostMessage(event.Channel, "#1e2124,#2C3849,#4180DB,#ffffff,#254a7d,#ffffff,#4180DB,#c90828", Struct.SlackParams)
		return true
	case "ricard":
		event.API.PostMessage(event.Channel, "#004684,#395882,#FFD300,#000000,#594e14,#ffffff,#FFD300,#c90828", Struct.SlackParams)
		return true
	}
	return false
}
