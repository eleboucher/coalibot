package miscs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/eleboucher/coalibot/utils"
	"github.com/slack-go/slack"
)

func Music(option string, event *utils.Message) bool {
	file, err := os.OpenFile("music.json", os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadFile("music.json")
	var musics []utils.Music

	json.Unmarshal(byteValue, &musics)
	if len(musics) == 0 {
		return false
	}
	music := musics[rand.Int()%len(musics)]
	if music.Login == "pk" {
		music.Login = "p/k"
	}
	utils.PostMsg(event, slack.MsgOptionText(music.Login+" "+music.Link, false))
	return true
}
