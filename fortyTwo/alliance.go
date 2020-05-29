package fortyTwo

import (
	"sort"
	"strconv"

	"github.com/eleboucher/coalibot/utils"

	"github.com/nlopes/slack"
)

func Alliance(option string, event *utils.Message) bool {
	msgRef := slack.NewRefToMessage(event.Channel, event.Timestamp)
	go event.API.AddReaction(":the-alliance:", msgRef)
	if event.Channel == "G833JLNS0" {
		coalitions, err := event.FortyTwo.GetCoalitionsByBloc(1, nil)
		if err != nil {
			return false
		}
		sort.Slice(coalitions, func(i, j int) bool {
			return coalitions[i].Score > coalitions[j].Score
		})
		var i int
		for i = 0; i < len(coalitions); i++ {
			if coalitions[i].ID == 2 {
				break
			}
		}
		rank := strconv.Itoa(i + 1)

		if i == 0 {
			diff := strconv.Itoa(coalitions[0].Score - coalitions[1].Score)
			utils.PostMsg(event, slack.MsgOptionText("Felicitations Nous sommes premiers avec "+diff+" points d'avance. :the-alliance:", false))
		} else {
			diff := strconv.Itoa(coalitions[0].Score - coalitions[i].Score)
			utils.PostMsg(event, slack.MsgOptionText("Nous sommes à la "+rank+" eme place avec "+diff+" points de retard. :the-alliance:", false))
		}
	} else {
		utils.PostMsg(event, slack.MsgOptionText(":the-alliance:", false))

	}
	return true
}
