package miscs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/eleboucher/coalibot/utils"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

var CorrectLinkRe = regexp.MustCompile(`(?:youtube\.com\/\S*(?:(?:\/e(?:mbed))?\/|watch\/?\?(?:\S*?&?v=))|youtu\.be\/)([a-zA-Z0-9_-]{6,11})`)

func AddMusic(option string, event *utils.Message) bool {
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
	var musics []utils.Music

	if err := json.Unmarshal(byteValue, &musics); err != nil {
		log.Error(err)
		return false
	}
	u, err := event.API.GetUserInfo(event.User)
	if err != nil {
		return false
	}
	login := u.Name
	link := strings.Split(option, " ")[0]
	if checkDuplicate(musics, link) && (CorrectLinkRe.MatchString(link) || strings.Contains(link, "soundcloud")) {
		newLink := utils.Music{Link: link, Login: login}
		musics = append(musics, newLink)
		toJson, _ := json.Marshal(musics)
		if _, err := file.Write(toJson); err != nil {
			log.Error(err)
			return false
		}
		if err := file.Sync(); err != nil {
			log.Error(err)
			return false
		}
		utils.PostMsg(event, slack.MsgOptionText("Musique ajoutée", false))
	} else {
		utils.PostMsg(event, slack.MsgOptionText("Lien incorrect ou déjà enregistré", false))
		return false
	}
	return true
}

func checkDuplicate(musics []utils.Music, link string) bool {
	for i := 0; i < len(musics); i++ {
		if musics[i].Link == link {
			return false
		}
	}
	return true
}
