package Miscs

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/genesixx/coalibot/Struct"
)

const (
	Decisecond = 100 * time.Millisecond
	Day        = 24 * time.Hour
)

func Gfaim(option string, event *Struct.Message) bool {
	now := time.Now()
	switch {
	case now.Hour() >= 12 && now.Hour() < 13:
		event.API.PostMessage(event.Channel, "C'est l'heure du dejeuner!", Struct.SlackParams)
	case now.Hour() < 12:
		d := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location()).Sub(now)
		event.API.PostMessage(event.Channel, "Dejeuner dans "+FormatCountdown(d), Struct.SlackParams)
	case now.Hour() >= 16 && now.Hour() < 17:
		event.API.PostMessage(event.Channel, "C'est l'heure du gouter!", Struct.SlackParams)
	case now.Hour() >= 13 && now.Hour() < 16:
		d := time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, now.Location()).Sub(now)
		event.API.PostMessage(event.Channel, "Gouter dans "+FormatCountdown(d), Struct.SlackParams)
	case now.Hour() >= 19 && now.Hour() < 20:
		event.API.PostMessage(event.Channel, "C'est l'heure du diner!", Struct.SlackParams)
	case now.Hour() >= 17 && now.Hour() < 19:
		d := time.Date(now.Year(), now.Month(), now.Day(), 19, 0, 0, 0, now.Location()).Sub(now)
		event.API.PostMessage(event.Channel, "Diner dans "+FormatCountdown(d), Struct.SlackParams)
	case now.Hour() > 20:
		event.API.PostMessage(event.Channel, "C' est plus l'heure de manger", Struct.SlackParams)
	}
	return true
}

func Apero(option string, event *Struct.Message) bool {
	res, err := http.Get("http://estcequecestbientotlapero.fr/")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	apero := strings.Join(strings.Fields(doc.Find("h2").First().Text()), " ")
	event.API.PostMessage(event.Channel, apero, Struct.SlackParams)
	return true
}

func FormatCountdown(ts time.Duration) string {

	ts += +Decisecond / 2
	ts = ts % Day
	h := ts / time.Hour
	ts = ts % time.Hour
	m := ts / time.Minute
	ts = ts % time.Minute
	s := ts / time.Second
	return fmt.Sprintf("%02dh %02dm %02ds", h, m, s)
}
