package miscs

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nlopes/slack"

	"github.com/eleboucher/coalibot/utils"
)

func Vdm(option string, event *utils.Message) bool {
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
		utils.PostMsg(event, slack.MsgOptionText("max 10 requests", false), slack.MsgOptionPostMessageParameters(params))
		return true
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	elem := doc.Find(".article-contents").ChildrenFiltered("a")
	vdms := ""
	for i := 0; i < nb; i++ {
		vdm := strings.Join(strings.Fields(elem.Eq(i).Text()), " ")
		if vdm != "" && strings.HasPrefix(vdm, "Aujourd'hui,") {
			vdms += fmt.Sprintf(">%s\n\n", vdm)
		}
	}
	if vdms != "" {
		if nb > 1 {
			utils.PostMsg(event, slack.MsgOptionText(vdms, false), slack.MsgOptionPostMessageParameters(params), slack.MsgOptionTS(event.Timestamp))

		} else {
			utils.PostMsg(event, slack.MsgOptionText(vdms, false), slack.MsgOptionPostMessageParameters(params))
		}
	}
	return true
}
