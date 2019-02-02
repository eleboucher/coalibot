package main

import (
	"net"
	"os"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/genesixx/coalibot/Struct"
	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
	"github.com/sirupsen/logrus"
	"gitlab.com/clafoutis/api42"
)

func main() {
	err := godotenv.Load()
	log := logrus.New()
	conn, err := net.Dial("tcp", os.Getenv("LOGSTASH_URL"))
	if err == nil {
		hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "Coalibot"}))
		log.Hooks.Add(hook)
	}
	api := slack.New(os.Getenv("SLACK_API_TOKEN"))
	rtm := api.NewRTM()
	reactions := InitReaction()
	// api.SetDebug(true)
	client, err := api42.NewAPI(os.Getenv("INTRA_CLIENT_ID"), os.Getenv("INTRA_SECRET"))
	if err != nil {
		log.Fatal("Error with the api")
		return
	}
	go rtm.ManageConnection()
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			log.Info("Ready")
		case *slack.MessageEvent:
			var message = Struct.Message{Message: ev.Msg.Text, Channel: ev.Msg.Channel, User: ev.Msg.User, Timestamp: ev.Msg.Timestamp, ThreadTimestamp: ev.Msg.ThreadTimestamp, API: api, FortyTwo: client}
			log.WithFields(logrus.Fields{"Channel": message.Channel, "User": message.User, "Text": message.Message}).Info()
			go React(message, reactions)

			if message.User != "" {
				go handleCommand(&message, log)
			}
		case *slack.RTMError:
			log.Fatal(ev.Error())
		case *slack.InvalidAuthEvent:
			log.Fatal("Invalid credentials")
			return
		}
	}
}
