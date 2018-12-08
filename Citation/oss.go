package Citation

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
	"github.com/nlopes/slack"
)

func Oss(option string, event *Struct.Message) bool {
	bytes, err := ioutil.ReadFile("Citation/oss.txt")

	if err != nil {
		fmt.Println(err)
		return false
	}
	splited := strings.Split(string(bytes), "\n")
	citation := splited[rand.Int()%len(splited)]
	params := slack.PostMessageParameters{UnfurlMedia: true, UnfurlLinks: true, Markdown: true, IconURL: "https://static-cdn.jtvnw.net/emoticons/v1/518312/3.0", Username: "Hubert Bonisseur de La Bath"}
	Utils.PostMsg(event, slack.MsgOptionText("> "+citation, false), slack.MsgOptionPostMessageParameters(params))
	return true
}
