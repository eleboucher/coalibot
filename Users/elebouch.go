package Users

import (
	"fmt"
	"time"

	"github.com/genesixx/coalibot/Struct"
)

func Elebouch(option string, event *Struct.Message) bool {
	loc, _ := time.LoadLocation("Europe/Paris")
	stage, _ := time.ParseInLocation("2006-01-02 15:04", "2018-11-05 09:30", loc)
	ts := -time.Since(stage)

	ts += +Decisecond / 2
	d := (ts / Day)
	ts = ts % Day
	h := ts / time.Hour
	ts = ts % time.Hour
	m := ts / time.Minute
	ts = ts % time.Minute
	s := ts / time.Second
	event.API.PostMessage(event.Channel, fmt.Sprintf("Debut du stage dans %02d jours, %02d heures %02d minutes %02d secondes", d, h, m, s), Struct.SlackParams)
	return true
}
