package Bars

import "github.com/genesixx/coalibot/Struct"

func Cdt(option string, event *Struct.Message) bool {
	if IsSpritzOpen() {
		event.API.PostMessage(event.Channel, "Go Spritz!", Struct.SlackParams)
		return true
	} else {
		event.API.PostMessage(event.Channel, "Le Spritz reste mieux!", Struct.SlackParams)
		return true
	}
}
