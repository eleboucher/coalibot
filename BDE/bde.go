package BDE

import (
	"strings"

	"github.com/genesixx/coalibot/Struct"
)

func Bde(option string, event *Struct.Message) bool {
	if option == "" {
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
