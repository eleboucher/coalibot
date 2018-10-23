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

func Calba(option string, event *Struct.Message) bool {
	loc, _ := time.LoadLocation("Europe/Paris")
	stage, _ := time.ParseInLocation("2006-01-02 15:04", "2018-12-03 17:15", loc)
	ts := -time.Since(stage)

	ts += +Decisecond / 2
	d := (ts / Day)
	ts = ts % Day
	h := ts / time.Hour
	ts = ts % time.Hour
	m := ts / time.Minute
	ts = ts % time.Minute
	s := ts / time.Second
	event.API.PostMessage(event.Channel, fmt.Sprintf("Libération dans %02d jours, %02d heures %02d minutes %02d secondes, d, h, m, s), Struct.SlackParams)
	return true
}
