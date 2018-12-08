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

func Kaamelott(option string, event *Struct.Message) bool {
	bytes, err := ioutil.ReadFile("Citation/kaamelott.txt")

	if err != nil {
		fmt.Println(err)
		return false
	}
	splited := strings.Split(string(bytes), "\n")
	citation := splited[rand.Int()%len(splited)]
	params := slack.PostMessageParameters{UnfurlMedia: true, UnfurlLinks: true, Markdown: true, IconURL: "https://img15.hostingpics.net/pics/4833663350.jpg", Username: "Perceval"}
	Utils.PostMsg(event, slack.MsgOptionText("> "+citation, false), slack.MsgOptionPostMessageParameters(params))
	return true
}
