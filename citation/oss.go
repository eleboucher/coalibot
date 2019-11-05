package citation

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/genesixx/coalibot/utils"
	"github.com/nlopes/slack"
)

func Oss(option string, event *utils.Message) bool {
	bytes, err := ioutil.ReadFile("citation/oss.txt")

	if err != nil {
		fmt.Println(err)
		return false
	}
	splitedText := strings.Split(string(bytes), "\n")
	citation := getRandomQuote(splitedText)

	params := slack.PostMessageParameters{UnfurlMedia: true, UnfurlLinks: true, Markdown: true, IconURL: "https://static-cdn.jtvnw.net/emoticons/v1/518312/3.0", Username: "Hubert Bonisseur de La Bath"}
	utils.PostMsg(event, slack.MsgOptionText("> "+citation, false), slack.MsgOptionPostMessageParameters(params))
	return true
}
