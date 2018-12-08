package Miscs

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
	"github.com/nlopes/slack"
)

func Meteo(option string, event *Struct.Message) bool {
	var lat = "48.90"
	var lon = "2.32"
	if option != "" && len(strings.Split(option, " ")) > 0 {
		if len(strings.Split(option, " ")) > 2 {
			Utils.PostMsg(event, slack.MsgOptionText("`bc meteo || bc meteo 48.9 2.32`", false))
			return false
		}
		lat = strings.Split(option, " ")[0]
		lon = strings.Split(option, " ")[1]

		if a, _ := strconv.ParseFloat(lat, 64); a > 90 || a < -90 {
			Utils.PostMsg(event, slack.MsgOptionText("`Latitude incorrecte`", false))
			return false
		}
		if b, _ := strconv.ParseFloat(lon, 64); b > 90 || b < -90 {
			Utils.PostMsg(event, slack.MsgOptionText("`Longitude incorrecte`", false))
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

	Utils.PostMsg(event, slack.MsgOptionText("```"+meteo+"```", false))
	return true
}
