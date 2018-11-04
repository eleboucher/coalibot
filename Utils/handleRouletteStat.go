package Utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/genesixx/coalibot/Struct"
)

func HandleRouletteStat(event *Struct.Message) {
	user, err := event.API.GetUserInfo(event.User)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Open our jsonFile
	file, err := os.OpenFile("rouletteStat.json", os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadFile("rouletteStat.json")
	// a map container to decode the JSON structure into
	c := make(map[string]int)

	// unmarschal JSON
	json.Unmarshal(byteValue, &c) // these lines to see the difference
	if c[user.Name] != 0 {
		c[user.Name] = c[user.Name] + 1
	} else {
		c[user.Name] = 1
	}
	toJson, _ := json.Marshal(c)
	file.Write(toJson)
	file.Sync()
}
