package FortyTwo

import (
	"strings"
	"time"

	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
	"gitlab.com/clafoutis/api42"
)

func Where(option string, event *Struct.Message) bool {
	params := api42.NewParameter()
	params.AddFilter("active", true)
	if len(strings.Split(option, " ")) == 4 && (strings.IndexAny(option, "le branle couille") != -1 || strings.IndexAny(option, "la branle couille") != -1) {
		y, m, d := time.Now().Date()
		rangeBegin := time.Date(y, m, d, 0, 0, 0, 0, time.Now().Location())
		rangeEnd := rangeBegin.AddDate(0, 0, -7)
		user := strings.Split(option, " ")[3]
		logtime := Utils.IntraLogtime(user, rangeEnd, rangeBegin, event.FortyTwo)
		if logtime.Hours() >= 35 {
			event.API.PostMessage(event.Channel, "*"+user+"* is not a branle couille", Struct.SlackParams)
			return true
		}
		data, err := event.FortyTwo.GetUserLocations(user, params)
		if err != nil {
			event.API.PostMessage(event.Channel, "login invalide", Struct.SlackParams)
			return false
		}

		if len(data) == 0 || data[0].EndAt != nil {
			event.API.PostMessage(event.Channel, "*"+user+"* est hors-ligne", Struct.SlackParams)
		} else {
			event.API.PostMessage(event.Channel, "*"+user+"* est à la place "+data[0].Host+"*", Struct.SlackParams)
		}
		return true
	}
	var user string

	if option != "" && len(strings.Split(option, " ")) == 1 {
		user = strings.Split(option, " ")[0]
		if user[0] == '<' && user[len(user)-1] == '>' && user[1] == '@' {
			u, err := event.API.GetUserInfo(user[2 : len(user)-1])
			if err != nil {
				return false
			}
			user = u.Profile.Email[0:strings.IndexAny(u.Profile.Email, "@")]
		}
	} else {
		u, err := event.API.GetUserInfo(event.User)
		if err != nil {
			return false
		}
		user = u.Profile.Email[0:strings.IndexAny(u.Profile.Email, "@")]
	}

	if user[0] == '!' || user[0] == '?' {
		return false
	}
	if user == "dieu" {
		user = "elebouch"
	} else if user == "manager" {
		user = "vtennero"
	}
	if user == "guardians" || user == "gardiens" {

		var guardians = []string{
			"dcirlig",
			"vtennero",
			"elebouch",
			"fbabin",
			"tbailly-",
			"mmerabet",
			"aledru",
			"dlavaury",
		}
		for i := 0; i < len(guardians); i++ {
			data, err := event.FortyTwo.GetUserLocations(guardians[i], params)
			if err != nil {
				event.API.PostMessage(event.Channel, "login invalide", Struct.SlackParams)
				return false
			}
			if len(data) == 0 || data[0].EndAt != nil {
				event.API.PostMessage(event.Channel, "*"+guardians[i]+"* est hors-ligne", Struct.SlackParams)
			} else {
				event.API.PostMessage(event.Channel, "*"+guardians[i]+"* est à la place *"+data[0].Host+"*", Struct.SlackParams)
			}
		}
		return true
	}
	data, err := event.FortyTwo.GetUserLocations(user, params)
	if err != nil {
		event.API.PostMessage(event.Channel, "login invalide", Struct.SlackParams)
		return false
	}

	if len(data) == 0 || data[0].EndAt != nil {
		event.API.PostMessage(event.Channel, "*"+user+"* est hors-ligne", Struct.SlackParams)
	} else {
		event.API.PostMessage(event.Channel, "*"+user+"* est à la place *"+data[0].Host+"*", Struct.SlackParams)
	}
	return true
}
