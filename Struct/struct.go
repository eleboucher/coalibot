package Struct

import (
	"regexp"

	"github.com/nlopes/slack"
	"gitlab.com/clafoutis/api42"
)

type Message struct {
	Message   string
	Channel   string
	User      string
	Timestamp string
	API       *slack.Client
	FortyTwo  *api42.Client42
	React     []React
}

type React struct {
	Name      string   `json:"name"`
	Reaction  string   `json:"reaction"`
	Reactions []string `json:"reactions"`
	Match     string   `json:"match"`
	Compiled  *regexp.Regexp
}
type Music struct {
	Login string `json:"login"`
	Link  string `json:"link"`
}

var SlackParams = slack.PostMessageParameters{UnfurlMedia: true, UnfurlLinks: true, Markdown: true}
