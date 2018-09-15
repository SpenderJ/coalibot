package Struct

import (
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
}

var SlackParams = slack.PostMessageParameters{UnfurlMedia: true, UnfurlLinks: true, Markdown: true}