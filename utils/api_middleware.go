package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"

	"gitlab.com/clafoutis/api42"
)

// FortyTwoMiddleware is a middleware that makes sure that the api client for the fortytwo api
// is properly created
func FortyTwoMiddleware(fn func(string, *Message) bool) func(string, *Message) bool {
	return func(cmd string, config *Message) bool {
		if config.FortyTwo == nil {
			client, err := api42.NewAPI(os.Getenv("INTRA_CLIENT_ID"), os.Getenv("INTRA_SECRET"))
			if err != nil {
				log.Errorf("Error with the api, %s", err)
				PostMsg(
					config,
					slack.MsgOptionText(err.Error(), false),
					slack.MsgOptionTS(config.Timestamp),
				)
			}
			config.FortyTwo = client
		}
		return fn(cmd, config)
	}
}
