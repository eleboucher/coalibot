package Miscs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/genesixx/coalibot/Struct"
)

type s_roulette struct {
	user string
	bang int
}

func RouletteTop(option string, event *Struct.Message) bool {
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
	var roulette []s_roulette
	for k, v := range c {
		roulette = append(roulette, s_roulette{k, v})
	}
	sort.Slice(roulette, func(i, j int) bool {
		return roulette[i].bang > roulette[j].bang
	})
	var ret = "Score Roulette:\n"
	for i := 0; i < 5; i++ {
		if i < len(roulette) {
			ret += fmt.Sprintf("*%s*: %d Bangs\n", roulette[i].user, roulette[i].bang)
		}
	}
	event.API.PostMessage(event.Channel, ret, Struct.SlackParams)
	return true
}
