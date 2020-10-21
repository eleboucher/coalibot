package miscs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/eleboucher/coalibot/utils"
	log "github.com/sirupsen/logrus"

	"github.com/slack-go/slack"
)

func RouletteStat(option string, event *utils.Message) bool {
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
	if err := json.Unmarshal(byteValue, &c); err != nil {
		log.Error(err)
		return false
	}

	var count = 0
	if c[user.Name] != 0 {
		count = c[user.Name]
	}
	countstr := strconv.Itoa(count)
	utils.PostMsg(event, slack.MsgOptionText("<@"+event.User+">: "+countstr+" Bang!", false))
	return true
}
