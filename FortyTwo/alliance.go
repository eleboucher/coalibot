package FortyTwo

import (
	"sort"
	"strconv"

	"github.com/genesixx/coalibot/Struct"
)

func Alliance(option string, event *Struct.Message) bool {
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
		event.API.PostMessage(event.Channel, "Felicitations Nous sommes premiers avec "+diff+" points d'avance. :the-alliance:", Struct.SlackParams)
	} else {
		diff := strconv.Itoa(coalitions[0].Score - coalitions[i].Score)
		event.API.PostMessage(event.Channel, "Nous sommes à la "+rank+" eme place avec "+diff+" points de retard. :the-alliance:", Struct.SlackParams)
	}
	return true
}
