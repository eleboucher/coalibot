package miscs

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/eleboucher/coalibot/utils"
	"github.com/nlopes/slack"
)

func Meteo(option string, event *utils.Message) bool {
	var lat = "48.90"
	var lon = "2.32"
	var options string
	if option != "" {
		options = strings.ReplaceAll(option, " ", "+")
	} else {
		options = lat + "," + lon
	}
	res, err := http.Get("http://en.wttr.in/" + options + "?T0")
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

	utils.PostMsg(event, slack.MsgOptionText("```"+meteo+"```", false))
	return true
}
