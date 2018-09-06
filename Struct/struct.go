package Struct

import "github.com/nlopes/slack"

type Message struct {
	Message   string
	Channel   string
	User      string
	Timestamp string
	API       *slack.Client
}
