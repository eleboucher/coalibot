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
	for index, elem := range coalitions {
		if elem.id == 2 {
			break
		}
	}
	rank, _ := strconv.Itoa(index + 1)
	if index == 0 {
		event.API.PostMessage(event.Channel, "Felicitations Nous sommes premiers avec "+rank+" points d'avance. :the-alliance:", Struct.SlackParams)
	}
	else {
		diff,_  := strconv.Itoa(coalitions[0] - coalitions[index])
		event.API.PostMessage(event.Channel, "Nous sommes à la "+rank+" eme place avec " + diff + " points de retard. :the-alliance:", Struct.SlackParams)
	}
}
