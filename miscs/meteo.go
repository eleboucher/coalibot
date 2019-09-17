package miscs

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/genesixx/coalibot/utils"
	"github.com/nlopes/slack"
)

func Meteo(option string, event *utils.Message) bool {
	res, err := http.Get("http://en.wttr.in/" + strings.ReplaceAll(option, " ", "+") + "?T0")
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
