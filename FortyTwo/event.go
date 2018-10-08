package FortyTwo

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
	"gitlab.com/clafoutis/api42"
)

func Event(option string, event *Struct.Message) bool {
	var beginAt, endAt string
	if option == "" {
		y, m, d := time.Now().Date()
		rangeBegin := time.Date(y, m, d, 0, 0, 0, 0, time.Now().Location())
		endAt = rangeBegin.AddDate(0, 0, 1).Format("2006-01-02")
		beginAt = rangeBegin.Format("2006-01-02")
	} else if len(strings.Split(option, " ")) == 1 {
		rangeBegin, _ := time.Parse("02/01/2006", strings.Split(option, " ")[0])
		endAt = rangeBegin.AddDate(0, 0, 1).Format("2006-01-02")
		beginAt = rangeBegin.Format("2006-01-02")
	} else {
		return false
	}
	params := api42.NewParameter()
	params.AddRange("begin_at", beginAt, endAt)
	data, err := event.FortyTwo.GetEventsByCampus("1", params)
	if err != nil {
		return false
	}
	sort.Slice(data, func(i, j int) bool { return data[i].BeginAt.Before(*data[j].BeginAt) })
	if len(data) == 0 {
		event.API.PostMessage(event.Channel, "Pas d'event ce jour!", Struct.SlackParams)
		return true
	}
	for i := 0; i < len(data); i++ {
		var desc = data[i].Description
		if len(data[i].Description) > 150 {
			desc = data[i].Description[:150]
		}
		params := Struct.SlackParams
		attachments := slack.Attachment{
			Title:     data[i].Name,
			TitleLink: "https://profile.intra.42.fr/events/" + strconv.Itoa(data[i].ID),
			Text:      desc + "...",
			Footer:    strconv.Itoa(data[i].NbrSubscribers) + "/" + strconv.Itoa(data[i].MaxPeople) + " Participants",
			Ts:        json.Number(strconv.FormatInt(int64(data[i].BeginAt.Unix()), 10)),
			Color:     "#01babc",
		}
		params.Attachments = []slack.Attachment{attachments}
		event.API.PostMessage(event.Channel, "", params)
	}
	return true
}
