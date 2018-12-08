package Miscs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/genesixx/coalibot/Utils"

	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func RouletteStat(option string, event *Struct.Message) bool {
	user, err := event.API.GetUserInfo(event.User)
	if err != nil {
		fmt.Println(err)
		return false
	}
	file, err := os.OpenFile("rouletteStat.json", os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadFile("rouletteStat.json")
	c := make(map[string]int)

	// unmarschal JSON
	json.Unmarshal(byteValue, &c)

	var count = 0
	if c[user.Name] != 0 {
		count = c[user.Name]
	}
	countstr := strconv.Itoa(count)
	Utils.PostMsg(event, slack.MsgOptionText("<@"+event.User+">: "+countstr+" Bang!", false))
	return true
}
