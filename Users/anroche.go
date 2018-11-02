package Users

import (
	"time"

	"github.com/genesixx/coalibot/Struct"
)

const (
	Decisecond = 100 * time.Millisecond
	Day        = 24 * time.Hour
)

func Anroche(option string, event *Struct.Message) bool {
	event.API.PostMessage(event.Channel, "https://www.youtube.com/watch?v=L0MK7qz13bU", Struct.SlackParams)
	return true
}
