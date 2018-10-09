package Citation

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func Mhenni(option string, event *Struct.Message) bool {
	bytes, err := ioutil.ReadFile("Citation/mhenni.txt")

	if err != nil {
		fmt.Println(err)
		return false
	}
	splited := strings.Split(string(bytes), "\n")
	citation := splited[rand.Int()%len(splited)]
	params := slack.PostMessageParameters{UnfurlMedia: true, UnfurlLinks: true, Markdown: true, IconURL: "https://risibank.fr/cache/stickers/d656/65606-full.png", Username: "Mohammed Henni"}
	event.API.PostMessage(event.Channel, "> "+citation, params)
	return true
}
