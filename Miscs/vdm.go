package Miscs

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nlopes/slack"

	"github.com/genesixx/coalibot/Struct"
)

func Vdm(option string, event *Struct.Message) bool {
	res, err := http.Get("https://www.viedemerde.fr/aleatoire")
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
	vdm := strings.Join(strings.Fields(doc.Find(".panel-content").ChildrenFiltered(".block").Has("a").First().Text()), " ")
	fmt.Println(vdm)
	params := slack.PostMessageParameters{UnfurlMedia: true, UnfurlLinks: true, Markdown: true, IconURL: "http://golem13.fr/wp-content/uploads/2012/10/vdm.gif", Username: "Vie De Merde"}
	event.API.PostMessage(event.Channel, "> "+vdm, params)

	return true
}
