package miscs

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/eleboucher/coalibot/utils"
	"github.com/nlopes/slack"
)

var helper = "bc roll nbResultat [option1, option2, ...]\nbc roll nbResultat min-max\nbc roll nbResultat max```"

// Returns an int >= min, < max
func randomInt(min int, max int, maxincluded bool) int {
	if maxincluded {
		return min + rand.Intn(max-min+1)
	}
	return min + rand.Intn(max-min)
}

func Roll(option string, event *utils.Message) bool {
	splited := strings.Split(option, " ")
	length, err := strconv.Atoi(splited[0])

	if err != nil {
		utils.PostMsg(event, slack.MsgOptionText(helper, false), slack.MsgOptionTS(event.Timestamp))
		return false
	}

	if len(splited) == 1 {
		str := strconv.Itoa(randomInt(1, length, true))

		utils.PostMsg(event, slack.MsgOptionText(str, false), slack.MsgOptionTS(event.Timestamp))
		return true

	} else if len(splited) >= 2 && strings.IndexAny(option, "[") != -1 &&
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
		utils.PostMsg(event, slack.MsgOptionText(str, false), slack.MsgOptionTS(event.Timestamp))
		return true
	} else if matched := regexp.MustCompile(`^\d+-\d+$`); len(splited) == 2 && matched.MatchString(splited[1]) == true {
		fmt.Println(strings.Split(splited[1], "-")[0])
		min, err := strconv.Atoi(strings.Split(splited[1], "-")[0])
		max, err1 := strconv.Atoi(strings.Split(splited[1], "-")[1])
		if err != nil || err1 != nil || max < min {
			return false
		}
		if length > 100 || max > 1000000 || length <= 0 || max <= 0 || min < 0 {
			utils.PostMsg(event, slack.MsgOptionText("taille max == 100 et tailledude max == 1000000", false), slack.MsgOptionTS(event.Timestamp))
			return false
		}
		var str string
		for i := 0; i < length; i++ {
			str += strconv.Itoa(randomInt(min, max, true))
			if i < length-1 {
				str += " "
			}
		}
		utils.PostMsg(event, slack.MsgOptionText(str, false), slack.MsgOptionTS(event.Timestamp))
		return true
	} else if max, err := strconv.Atoi(splited[1]); err == nil {

		if length > 100 || max > 1000000 || length <= 0 || max <= 0 {
			utils.PostMsg(event, slack.MsgOptionText("taille max == 100 et tailledude max == 1000000", false), slack.MsgOptionTS(event.Timestamp))
			return false
		}
		var str string
		for i := 0; i < length; i++ {
			str += strconv.Itoa(randomInt(1, max, true))
			if i < length-1 {
				str += " "
			}
		}
		utils.PostMsg(event, slack.MsgOptionText(str, false), slack.MsgOptionTS(event.Timestamp))
		return true
	}
	utils.PostMsg(event, slack.MsgOptionText(helper, false), slack.MsgOptionTS(event.Timestamp))
	return false
}
