package FortyTwo

import (
	"fmt"
	"strings"

	"github.com/genesixx/coalibot/Struct"
)

func Prof(option string, event *Struct.Message) bool {
	var user string

	if len(option) > 0 {
		user = strings.Split(option, " ")[0]
	} else {
		u, err := event.API.GetUserInfo(event.User)
		if err != nil {
			fmt.Println(err)
			return false
		}
		user = u.Profile.Email[0:strings.IndexAny(u.Profile.Email, "@")]
	}
	data, err := event.FortyTwo.GetUser(user)
	if err != nil {
		return false
	}

	coaldata, _ := event.FortyTwo.GetCoalitionUser(user)
	var lvlPiscine string
	if data.PoolYear == "2013" || data.PoolYear == "2014" {
		lvlPiscine = fmt.Sprintf("%2.6d", 0)
	} else if len(data.CursusUsers) == 1 {
		lvlPiscine = fmt.Sprintf("%2.6f", data.CursusUsers[0].Level)
	} else {
		lvlPiscine = fmt.Sprintf("%2.6f", data.CursusUsers[1].Level)
	}
	fmt.Println(coaldata, lvlPiscine)
	return true
}
