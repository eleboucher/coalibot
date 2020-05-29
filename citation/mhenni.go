package citation

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/eleboucher/coalibot/utils"

	"github.com/slack-go/slack"
)

func Mhenni(option string, event *utils.Message) bool {
	bytes, err := ioutil.ReadFile("citation/mhenni.txt")

	if err != nil {
		fmt.Println(err)
		return false
	}
	splitedText := strings.Split(string(bytes), "\n")
	citation := getRandomQuote(splitedText)

	params := slack.PostMessageParameters{UnfurlMedia: true, UnfurlLinks: true, Markdown: true, IconURL: "https://risibank.fr/cache/stickers/d656/65606-full.png", Username: "Mohammed Henni"}
	utils.PostMsg(event, slack.MsgOptionText("> "+citation, false), slack.MsgOptionPostMessageParameters(params))
	return true
}
