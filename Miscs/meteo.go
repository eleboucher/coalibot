package Miscs

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/genesixx/coalibot/Struct"
)

func Meteo(option string, event *Struct.Message) bool {
	var lat = "48.90"
	var lon = "2.32"
	if option != "" && len(strings.Split(option, " ")) > 0 {
		if len(strings.Split(option, " ")) > 2 {
			event.API.PostMessage(event.Channel, "`bc meteo || bc meteo 48.9 2.32`", Struct.SlackParams)
			return false
		}
		lat = strings.Split(option, " ")[0]
		lon = strings.Split(option, " ")[1]

		if a, _ := strconv.ParseFloat(lat, 64); a > 90 || a < -90 {
			event.API.PostMessage(event.Channel, "`Latitude incorrecte`", Struct.SlackParams)
			return false
		}
		if b, _ := strconv.ParseFloat(lon, 64); b > 90 || b < -90 {
			event.API.PostMessage(event.Channel, "`Longitude incorrecte`", Struct.SlackParams)
			return false
		}
	}
	res, err := http.Get("http://fr.wttr.in/" + lat + "," + lon + "?T0")
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
	meteo := doc.Find("pre").Text()
	event.API.PostMessage(event.Channel, "```"+meteo+"```", Struct.SlackParams)
	return true
}
