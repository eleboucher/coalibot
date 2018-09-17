package Miscs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/genesixx/coalibot/Struct"
)

func AddMusic(option string, event *Struct.Message) bool {
	if option == "" {
		return false
	}
	file, err := os.OpenFile("music.json", os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadFile("music.json")
	var musics []Struct.Music

	json.Unmarshal(byteValue, &musics)
	u, err := event.API.GetUserInfo(event.User)
	if err != nil {
		return false
	}
	login := u.Name
	link := strings.Split(option, " ")[0]
	isCorrect, _ := regexp.Compile(`(?:youtube\.com\/\S*(?:(?:\/e(?:mbed))?\/|watch\/?\?(?:\S*?&?v=))|youtu\.be\/)([a-zA-Z0-9_-]{6,11})`)
	if checkDuplicate(musics, link) && (isCorrect.MatchString(link) || strings.IndexAny(link, "soundcloud") != -1) {
		newLink := Struct.Music{Link: link, Login: login}
		musics = append(musics, newLink)
		toJson, _ := json.Marshal(musics)
		file.Write(toJson)
		event.API.PostMessage(event.Channel, "Musique ajoutée", Struct.SlackParams)
	} else {
		event.API.PostMessage(event.Channel, "Lien incorrect ou déjà enregistré", Struct.SlackParams)
		return false
	}
	return true
}

func checkDuplicate(musics []Struct.Music, link string) bool {
	for i := 0; i < len(musics); i++ {
		if musics[i].Link == link {
			return false
		}
	}
	return true
}
