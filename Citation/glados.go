package Citation

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func GLaDOS(option string, event *Struct.Message) bool {
	bytes, err := ioutil.ReadFile("Citation/glados.txt")

	if err != nil {
		fmt.Println(err)
		return false
	}
	splited := strings.Split(string(bytes), "\n")
	citation := splited[rand.Int()%len(splited)]
	params := slack.PostMessageParameters{UnfurlMedia: true, UnfurlLinks: true, Markdown: true, IconURL: "https://vignette.wikia.nocookie.net/deathbattlefanon/images/e/e0/GLaDOS.png", Username: "GLaDOS"}
	event.API.PostMessage(event.Channel, "> "+citation, params)
	return true
}
