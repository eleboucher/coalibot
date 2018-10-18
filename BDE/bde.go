package BDE

import (
	"strings"

	"github.com/genesixx/coalibot/Struct"
)

func Bde(option string, event *Struct.Message) bool {
	var params = Struct.SlackParams
	params.IconURL = "https://bde.student42.fr/img/bde42-logo-1538664197.jpg"
	params.Username = "Undefined Bot"
	if option == "" {
		event.API.PostMessage(event.Channel, "Pour avoir accès au shop tapez `!bde shop`, pour avoir plus d'info sur la soirée Blood Horror Party tapez `!bde event`!", params)
		return true
	}
	switch strings.ToLower(strings.Split(option, " ")[0]) {
	case "shop":
		return Shop(option, event)
	case "event":
		return Event(option, event)
	}
	return true
}
