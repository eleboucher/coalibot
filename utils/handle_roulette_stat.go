package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

func HandleRouletteStat(event *Message) {
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

	// unmarshal JSON
	if err := json.Unmarshal(byteValue, &c); err != nil {
		log.Error(err)
		return
	}
	if c[user.Name] != 0 {
		c[user.Name] = c[user.Name] + 1
	} else {
		c[user.Name] = 1
	}
	toJson, _ := json.Marshal(c)
	if _, err := file.Write(toJson); err != nil {
		log.Error(err)
		return
	}
	if err := file.Sync(); err != nil {
		log.Error(err)
		return
	}
}
