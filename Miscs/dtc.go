package Miscs

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/nlopes/slack"

	"github.com/genesixx/coalibot/Struct"
)

func Dtc(option string, event *Struct.Message) bool {
	res, err := http.Get("https://danstonchat.com/random0.html")
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
	var dtc string
	length := doc.Find(".item").ChildrenFiltered(".item-content").Has("a").Length()
	doc.Find(".item").ChildrenFiltered(".item-content").Has("a").Eq(rand.Int()%length - 1).Each(func(i int, s *goquery.Selection) {
		sentence := s.Text()
		dtc += sentence + "\n"
	})
	params := slack.PostMessageParameters{UnfurlMedia: true, UnfurlLinks: true, Markdown: true, ThreadTimestamp: event.Timestamp, IconURL: "https://danstonchat.com/icache/size/300c300/themes/danstonchat2016/images/logo-og.png", Username: "Dans Ton Chat"}
	event.API.PostMessage(event.Channel, slack.MsgOptionText("```"+dtc+"```", false), slack.MsgOptionPostMessageParameters(params))

	return true
}
