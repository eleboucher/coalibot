package Users

import (
	"fmt"
	"time"

	"github.com/genesixx/coalibot/Struct"
)

const (
	Decisecond = 100 * time.Millisecond
	Day        = 24 * time.Hour
)

func Anroche(option string, event *Struct.Message) bool {
	stage, _ := time.Parse("2006-01-02 15:04", "2018-11-02 18:00")
	ts := -time.Since(stage)

	ts += +Decisecond / 2
	d := (ts / Day)
	ts = ts % Day
	h := ts / time.Hour
	ts = ts % time.Hour
	m := ts / time.Minute
	ts = ts % time.Minute
	s := ts / time.Second
	event.API.PostMessage(event.Channel, fmt.Sprintf("Fin du stage dans %02d days, %02d hours %02d minutes %02d seconds", d, h, m, s), Struct.SlackParams)
	return true
}
