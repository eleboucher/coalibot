package main

import (
	"os"

	"github.com/genesixx/coalibot/Struct"
	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
	"gitlab.com/clafoutis/api42"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
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
			log.Fatal("Ready")
		case *slack.MessageEvent:
			var message = Struct.Message{Message: ev.Msg.Text, Channel: ev.Msg.Channel, User: ev.Msg.User, Timestamp: ev.Msg.Timestamp, API: api, FortyTwo: client}
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
