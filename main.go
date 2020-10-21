package main

import (
	"os"

	"github.com/eleboucher/coalibot/utils"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"gitlab.com/clafoutis/api42"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	api := slack.New(os.Getenv("SLACK_API_TOKEN"))
	rtm := api.NewRTM()
	reactions := InitReaction()
	// api.SetDebug(true)
	client, err := api42.NewAPI(os.Getenv("INTRA_CLIENT_ID"), os.Getenv("INTRA_SECRET"))
	if err != nil {
		log.Error("Error with the api")
	}

	go rtm.ManageConnection()
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			log.Info("Ready")
		case *slack.MessageEvent:
			var message = utils.Message{
				Message:         ev.Msg.Text,
				Channel:         ev.Msg.Channel,
				User:            ev.Msg.User,
				Timestamp:       ev.Msg.Timestamp,
				ThreadTimestamp: ev.Msg.ThreadTimestamp,
				API:             api,
				FortyTwo:        client,
			}

			go React(message, reactions)

			if message.User != "" {
				go handleCommand(&message)
			}
		case *slack.RTMError:
			log.Fatal(ev.Error())
		case *slack.InvalidAuthEvent:
			log.Fatal("Invalid credentials")
			return
		}
	}
}
