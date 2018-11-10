package Citation

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/genesixx/coalibot/Struct"
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
	params := slack.PostMessageParameters{IconURL: "https://img15.hostingpics.net/pics/4833663350.jpg", Username: "Perceval"}
	event.API.PostMessage(event.Channel, slack.MsgOptionText("> "+citation, false), slack.MsgOptionPostMessageParameters(params))
	return true
}
