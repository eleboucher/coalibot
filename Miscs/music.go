package Miscs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/genesixx/coalibot/Struct"
)

func Music(option string, event *Struct.Message) bool {
	file, err := os.OpenFile("music.json", os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadFile("music.json")
	var musics []Struct.Music

	json.Unmarshal(byteValue, &musics)
	if len(musics) == 0 {
		return false
	}
	music := musics[rand.Int()%len(musics)]
	if music.Login == "pk" {
		music.Login = "p/k"
	}
	event.API.PostMessage(event.Channel, music.Login+" "+music.Link, Struct.SlackParams)
	return true
}
