package model

import (
	"encoding/json"
	"fmt"
	"github.com/hachi-n/command-wrap-notify/internal/http"
)

type Slack struct {
	Mention  string
	Channel  string
	Endpoint string
	Username string
}

type SlackMessage struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

func NewSlackMessage(channel string, username string, message string) *SlackMessage {
	return &SlackMessage{Channel: channel, Username: username, Text: message}
}

func NewSlack(mention string, channel string, endpoint string, username string) *Slack {
	return &Slack{
		Mention:  mention,
		Channel:  channel,
		Endpoint: endpoint,
		Username: username,
	}
}

func (s Slack) Notify(message string) ([]byte, error) {
	fixedMessage := fmt.Sprintf("<%s>\n%s", s.Mention, message)

	m := NewSlackMessage(s.Channel, s.Username, fixedMessage)
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return http.JsonPost(s.Endpoint, b)
}
