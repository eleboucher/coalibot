package citation

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/eleboucher/coalibot/utils"
	"github.com/slack-go/slack"
)

func GLaDOS(option string, event *utils.Message) bool {
	bytes, err := ioutil.ReadFile("citation/glados.txt")

	if err != nil {
		fmt.Println(err)
		return false
	}
	splitedText := strings.Split(string(bytes), "\n")
	citation := getRandomQuote(splitedText)
	params := slack.PostMessageParameters{UnfurlMedia: true, UnfurlLinks: true, Markdown: true, IconURL: "https://vignette.wikia.nocookie.net/deathbattlefanon/images/e/e0/GLaDOS.png", Username: "GLaDOS"}
	utils.PostMsg(event, slack.MsgOptionText("> "+citation, false), slack.MsgOptionPostMessageParameters(params))
	return true
}
