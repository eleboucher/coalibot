package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/hako/durafmt"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

func IndexOf(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}

func FmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02dh%02dm", h, m)
}

func PrettyDurationPrinting(d time.Duration) string {
	d = d.Round(time.Minute)
	if d.Hours() > 168 {
		return durafmt.ParseShort(d).String()
	}
	return durafmt.Parse(d).String()
}

func GetLogin(option string, event *Message) (string, bool) {
	var user string
	if option != "" && len(strings.Split(option, " ")) == 1 {
		user = strings.Split(option, " ")[0]
		if user[0] == '<' && user[len(user)-1] == '>' && user[1] == '@' {
			u, err := event.API.GetUserInfo(user[2 : len(user)-1])
			if err != nil {
				return "", true
			}
			user = u.Profile.Email[0:strings.IndexAny(u.Profile.Email, "@")]
		}
	} else {
		u, err := event.API.GetUserInfo(event.User)
		if err != nil {
			return "", true
		}
		user = u.Profile.Email[0:strings.IndexAny(u.Profile.Email, "@")]
	}
	return user, false
}

func Choice(option []string) string {
	return option[rand.Int()%len(option)]
}

func PostMsg(event *Message, options ...slack.MsgOption) {
	channel := event.Channel

	if event.ThreadTimestamp != "" {
		options = append(options, slack.MsgOptionTS(event.Timestamp))
	}
	if _, _, err := event.API.PostMessage(channel, options...); err != nil {
		log.Error(err)
	}
}
