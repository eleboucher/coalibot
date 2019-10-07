package fortyTwo

import (
	"fmt"
	"regexp"
	"time"

	"github.com/genesixx/coalibot/utils"
	"github.com/nlopes/slack"
	"gitlab.com/clafoutis/api42"
)

var guardians = []string{
	"dcirlig",
	"korlandi",
	"syzhang",
	"vtennero",
	"elebouch",
	"fbabin",
	"tbailly-",
	"mmerabet",
	"aledru",
	"dlavaury",
	"jauplat",
	"jraymond",
}

var branleCouilleRegex = regexp.MustCompile(`branles?.couilles?.(.+)`)

func branleCouille(user string, event *utils.Message) (string, error) {
	params := api42.NewParameter()
	y, m, d := time.Now().Date()
	rangeBegin := time.Date(y, m, d, 0, 0, 0, 0, time.Now().Location())
	rangeEnd := rangeBegin.AddDate(0, 0, -7)
	logtime := utils.IntraLogtime(user, rangeEnd, rangeBegin, event.FortyTwo)
	if logtime.Hours() >= 35 {
		return "*" + user + "* is not a branle couille", nil
	}
	data, err := event.FortyTwo.GetUserLocations(user, params)
	if err != nil {
		return "login invalide", err
	}

	return formatLocation(data, user), nil
}

func formatLocation(data []api42.Location42, user string) string {
	if len(data) == 0 || data[0].EndAt != nil {
		var diff = time.Now().Sub(*data[0].EndAt)

		return "*" + user + "* est hors-ligne depuis *" + utils.PrettyDurationPrinting(diff) + "*"
	}
	return "*" + user + "* est à la place *" + data[0].Host + "*"
}

func Where(option string, event *utils.Message) bool {
	params := api42.NewParameter()
	if branleCouilleRegex.MatchString(option) {
		user := branleCouilleRegex.FindStringSubmatch(option)[1]
		message, err := branleCouille(user, event)

		if err != nil {
			println(err)
			return false
		}
		utils.PostMsg(event, slack.MsgOptionText(message, false))
		return true
	}
	user, error := utils.GetLogin(option, event)
	if error == true {
		return false
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
		var str string
		for i := 0; i < len(guardians); i++ {
			data, err := event.FortyTwo.GetUserLocations(guardians[i], params)
			if err != nil {
				fmt.Println(err)
				str += "login invalide\n"
				return false
			}
			str += formatLocation(data, guardians[i]) + "\n"
		}
		utils.PostMsg(event, slack.MsgOptionText(str, false))

		return true
	}

	data, err := event.FortyTwo.GetUserLocations(user, params)
	if err != nil {
		utils.PostMsg(event, slack.MsgOptionText("login invalide", false))
		return false
	}

	utils.PostMsg(event, slack.MsgOptionText(formatLocation(data, user), false))
	return true
}
