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
	params := slack.PostMessageParameters{UnfurlMedia: true, UnfurlLinks: true, Markdown: true, IconURL: "https://is4-ssl.mzstatic.com/image/thumb/Purple118/v4/c6/56/f9/c656f9f4-8bbe-3881-b437-a2d8306dc417/AppIcon-0-1x_U007emarketing-0-85-220-7.png/246x0w.jpg", Username: "Mohammed Henni"}
	event.API.PostMessage(event.Channel, "> "+citation, params)
	return true
}
