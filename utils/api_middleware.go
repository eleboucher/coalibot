package utils

import (
	"os"

	log "github.com/sirupsen/logrus"

	"gitlab.com/clafoutis/api42"
)

// FortyTwoMiddleware is a middleware that makes sure that the api client for the fortytwo api
// is properly created
func FortyTwoMiddleware(fn func(string, *Message) bool) func(string, *Message) bool {
	return func(cmd string, config *Message) bool {
		if config.FortyTwo == nil {
			client, err := api42.NewAPI(os.Getenv("INTRA_CLIENT_ID"), os.Getenv("INTRA_SECRET"))
			if err != nil {
				log.Error("Error with the api")
			}
			config.FortyTwo = client
		}
		return fn(cmd, config)
	}
}
