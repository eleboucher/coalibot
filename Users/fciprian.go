package Users

import (
	"fmt"
	"time"

	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"

	"github.com/nlopes/slack"
)

func Fciprian(option string, event *Struct.Message) bool {
	loc, _ := time.LoadLocation("Europe/Paris")
	stage, _ := time.ParseInLocation("2006-01-02 15:04", "2018-11-27 18:30", loc)
	ts := -time.Since(stage)

	ts += +Decisecond / 2
	d := (ts / Day)
	ts = ts % Day
	h := ts / time.Hour
	ts = ts % time.Hour
	m := ts / time.Minute
	ts = ts % time.Minute
	s := ts / time.Second
	Utils.PostMsg(event, slack.MsgOptionText(fmt.Sprintf("Délivrance dans %02d jours, %02d heures %02d minutes %02d secondes", d, h, m, s), false))
	return true
}
