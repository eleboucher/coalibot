package Miscs

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
	"github.com/nlopes/slack"
)

func Roll(option string, event *Struct.Message) bool {
	splited := strings.Split(option, " ")
	length, err := strconv.Atoi(splited[0])

	if err != nil {
		return false
	}
	params := Struct.SlackParams
	params.ThreadTimestamp = event.Timestamp
	if len(splited) >= 2 && strings.IndexAny(option, "[") != -1 &&
		strings.IndexAny(option, "]") != -1 &&
		strings.IndexAny(option, "[") < strings.IndexAny(option, "]") {
		ranthings := strings.Split(
			strings.TrimSpace(
				option[strings.IndexAny(option, "[")+1:strings.IndexAny(option, "]")]),
			",")
		if len(ranthings) < 2 || len(ranthings) > 1000000 || length > 100 || length <= 0 {
			return false
		}
		var str string
		for i := 0; i < length; i++ {
			str += ranthings[rand.Intn(len(ranthings))]
			if i < length-1 {
				str += " "
			}
		}
		Utils.PostMsg(event, slack.MsgOptionText(str, false))
		return true
	} else if matched := regexp.MustCompile(`^\d+-\d+$`); len(splited) == 2 && matched.MatchString(splited[1]) == true {
		fmt.Println(strings.Split(splited[1], "-")[0])
		min, err := strconv.Atoi(strings.Split(splited[1], "-")[0])
		max, err1 := strconv.Atoi(strings.Split(splited[1], "-")[1])
		if err != nil || err1 != nil || max < min {
			return false
		}
		if length > 100 || max > 1000000 || length <= 0 || max <= 0 || min < 0 {
			Utils.PostMsg(event, slack.MsgOptionText("taille max == 100 et tailledude max == 1000000", false))
			return false
		}
		var str string
		for i := 0; i < length; i++ {
			str += strconv.Itoa(rand.Intn(max-min+1) + min)
			if i < length-1 {
				str += " "
			}
		}
		Utils.PostMsg(event, slack.MsgOptionText(str, false))
		return true
	} else if max, err := strconv.Atoi(splited[1]); err == nil {

		if length > 100 || max > 1000000 || length <= 0 || max <= 0 {
			Utils.PostMsg(event, slack.MsgOptionText("taille max == 100 et tailledude max == 1000000", false))
			return false
		}
		var str string
		for i := 0; i < length; i++ {
			str += strconv.Itoa(rand.Intn(max + 1))
			if i < length-1 {
				str += " "
			}
		}
		Utils.PostMsg(event, slack.MsgOptionText(str, false))
		return true
	}
	return false
}
