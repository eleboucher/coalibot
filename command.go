package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/eleboucher/coalibot/assos"
	"github.com/eleboucher/coalibot/bars"
	"github.com/eleboucher/coalibot/citation"
	"github.com/eleboucher/coalibot/database"
	"github.com/eleboucher/coalibot/fortyTwo"
	"github.com/eleboucher/coalibot/miscs"
	"github.com/eleboucher/coalibot/twitch"
	"github.com/eleboucher/coalibot/users"
	"github.com/eleboucher/coalibot/utils"
	"github.com/nlopes/slack"
	"github.com/sirupsen/logrus"
)

var commands = map[string]func(string, *utils.Message) bool{
	"hello":        miscs.Hello,
	"vdm":          miscs.Vdm,
	"roulette":     miscs.Roulette,
	"coin":         miscs.Coin,
	"meteo":        miscs.Meteo,
	"weather":      miscs.Meteo,
	"roll":         miscs.Roll,
	"roulettestat": miscs.RouletteStat,
	"score":        fortyTwo.Score,
	"alliance":     fortyTwo.Alliance,
	"prof":         fortyTwo.Prof,
	"logtime":      fortyTwo.Logtime,
	"who":          fortyTwo.Who,
	"where":        fortyTwo.Where,
	"oss":          citation.Oss,
	"kaamelott":    citation.Kaamelott,
	"mhenni":       citation.Mhenni,
	"glados":       citation.GLaDOS,
	"help":         miscs.Help,
	"music":        miscs.Music,
	"addmusic":     miscs.AddMusic,
	"dtc":          miscs.Dtc,
	"event":        fortyTwo.Event,
	"roulettetop":  miscs.RouletteTop,
	"anroche":      users.Anroche,
	"elebouch":     users.Elebouch,
	"spritz":       bars.Spritz,
	"cdt":          bars.Cdt,
	"moty":         bars.Moty,
	"gfaim":        miscs.Gfaim,
	"apero":        miscs.Apero,
	"skin":         miscs.Skin,
	"shop":         assos.Shop,
	"bde":          assos.Bde,
	"asso":         assos.Assos,
	"emote":        twitch.Emotes,
}

func handleCommand(event *utils.Message, log *logrus.Logger) {
	var isCommand = false
	var option = ""
	var command = ""

	sort.Strings(BlackList)
	i := sort.Search(len(BlackList),
		func(i int) bool { return BlackList[i] >= event.Channel })
	if (i < len(BlackList) && BlackList[i] == event.Channel) && !(strings.Index(strings.ToLower(event.Message), "bde") != -1 && event.Channel == "C04GT8U3Y") {
		return
	}

	event.Message = strings.Join(strings.Fields(event.Message), " ")

	splited := strings.Split(event.Message, " ")
	if event.Message == "" {
		return
	}
	if utils.IndexOf(strings.ToLower(splited[0]), []string{"coalibot", "bc", "cb"}) > -1 && len(splited) > 1 {
		command = strings.ToLower(splited[1])
		option = strings.Join(splited[2:], " ")
		isCommand = reply(command, event)
		if !isCommand && commands[command] != nil {
			isCommand = commands[strings.ToLower(command)](option, event)
		}
	} else if splited[0][0] == '!' && len(splited[0]) > 1 {
		command = strings.ToLower(splited[0][1:])
		option = strings.Join(splited[1:], " ")
		isCommand = reply(command, event)
		if !isCommand && commands[command] != nil {
			isCommand = commands[strings.ToLower(command)](option, event)
		}
	}
	if isCommand {
		log.WithFields(logrus.Fields{"Channel": event.Channel, "User": event.User, "command": event.Message}).Info()
		go database.AddCommand(event, command, option)
	}
}

func reply(command string, event *utils.Message) bool {
	// Open our jsonFile
	jsonFile, err := os.Open("reply.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// a map container to decode the JSON structure into
	c := make(map[string]interface{})

	// unmarschal JSON
	e := json.Unmarshal(byteValue, &c)
	if e != nil || c[command] == nil {
		return false
	}

	// output result to STDOUT
	fmt.Printf("reply %s\n", c[command].(string))
	utils.PostMsg(event, slack.MsgOptionText(c[command].(string), false))
	return true
}
