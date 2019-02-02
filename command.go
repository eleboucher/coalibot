package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"github.com/genesixx/coalibot/Assos"
	"github.com/genesixx/coalibot/Bars"
	"github.com/genesixx/coalibot/Citation"
	"github.com/genesixx/coalibot/Database"
	"github.com/genesixx/coalibot/FortyTwo"
	"github.com/genesixx/coalibot/Miscs"
	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Users"
	"github.com/genesixx/coalibot/Utils"
	"github.com/nlopes/slack"
	 "github.com/sirupsen/logrus"

)

var commands = map[string]func(string, *Struct.Message) bool{
	"hello":        Miscs.Hello,
	"vdm":          Miscs.Vdm,
	"roulette":     Miscs.Roulette,
	"coin":         Miscs.Coin,
	"meteo":        Miscs.Meteo,
	"roll":         Miscs.Roll,
	"roulettestat": Miscs.RouletteStat,
	"score":        FortyTwo.Score,
	"alliance":     FortyTwo.Alliance,
	"prof":         FortyTwo.Prof,
	"logtime":      FortyTwo.Logtime,
	"who":          FortyTwo.Who,
	"where":        FortyTwo.Where,
	"oss":          Citation.Oss,
	"kaamelott":    Citation.Kaamelott,
	"mhenni":       Citation.Mhenni,
	"glados":       Citation.GLaDOS,
	"help":         Miscs.Help,
	"music":        Miscs.Music,
	"addmusic":     Miscs.AddMusic,
	"dtc":          Miscs.Dtc,
	"event":        FortyTwo.Event,
	"roulettetop":  Miscs.RouletteTop,
	"anroche":      Users.Anroche,
	"calba":        Users.Calba,
	"fciprian":     Users.Fciprian,
	"elebouch":     Users.Elebouch,
	"spritz":       Bars.Spritz,
	"cdt":          Bars.Cdt,
	"moty":         Bars.Moty,
	"gfaim":        Miscs.Gfaim,
	"apero":        Miscs.Apero,
	"skin":         Miscs.Skin,
	"shop":         Assos.Shop,
	"bde":          Assos.Bde,
	"asso":         Assos.Assos,
}

func handleCommand(event *Struct.Message, log *logrus.Logger) {
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
	if Utils.IndexOf(strings.ToLower(splited[0]), []string{"coalibot", "bc", "cb"}) > -1 && len(splited) > 1 {
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
		go Database.AddCommand(event, command, option)
	}
}

func reply(command string, event *Struct.Message) bool {
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
	Utils.PostMsg(event, slack.MsgOptionText(c[command].(string), false))
	return true
}
