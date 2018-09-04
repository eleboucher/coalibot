package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

type Message struct {
	Message   string
	Channel   string
	User      string
	Timestamp string
	API       *slack.Client
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	api := slack.New(os.Getenv("SLACK_API_TOKEN"))
	rtm := api.NewRTM()
	api.SetDebug(true)

	go rtm.ManageConnection()
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			fmt.Println("Ready")
		case *slack.MessageEvent:
			var message = Message{Message: ev.Msg.Text, Channel: ev.Msg.Channel, User: ev.Msg.User, Timestamp: ev.Msg.Timestamp, API: api}
			go handleCommand(&message)
		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())
		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return
		}
	}
}
