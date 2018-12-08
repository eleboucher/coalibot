package Miscs

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nlopes/slack"

	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
)

func Vdm(option string, event *Struct.Message) bool {
	nb := 1
	params := slack.PostMessageParameters{UnfurlMedia: true, UnfurlLinks: true, Markdown: true, IconURL: "http://golem13.fr/wp-content/uploads/2012/10/vdm.gif", Username: "Vie De Merde"}

	res, err := http.Get("https://www.viedemerde.fr/aleatoire")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	if nb, err = strconv.Atoi(option); err != nil {
		nb = 1
	}
	if nb > 10 {
		Utils.PostMsg(event, slack.MsgOptionText("only 10 request max is allowed", false), slack.MsgOptionPostMessageParameters(params))
		return true
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	elem := doc.Find(".panel-content").ChildrenFiltered(".block").Has("a")
	vdms := ""
	for i := 0; i < nb; i++ {
		vdm := strings.Join(strings.Fields(elem.Eq(i).Text()), " ")
		if vdm != "" {
			vdms += fmt.Sprintf(">%s\n\n", vdm)
		}
	}
	if vdms != "" {
		Utils.PostMsg(event, slack.MsgOptionText(vdms, false), slack.MsgOptionPostMessageParameters(params))
	}
	return true
}
