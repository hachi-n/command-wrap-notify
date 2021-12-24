package slack

import (
	"encoding/json"
	"fmt"
	"github.com/hachi-n/command-wrap-notify/internal/http"
	"github.com/hachi-n/command-wrap-notify/internal/message"
)

type Slack struct {
	Mention  string
	Channel  string
	Endpoint string
	Username string
}

type slackMessage struct {
	channel     string
	username    string
	text        string
	attachments []attachment
}

type attachment struct {
	Color    string  `json:"color"`
	ImageUrl string  `json:"image_url"`
	Fields   []field `json:"fields"`
}
type field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func newSlackMessage(channel string, username string, message string) *slackMessage {
	return &slackMessage{channel: channel, username: username, text: message}
}

func (s *slackMessage) MarshalJSON() ([]byte, error) {
	type alias struct {
		Channel     string       `json:"channel"`
		Username    string       `json:"username"`
		Text        string       `json:"text"`
		Attachments []attachment `json:"attachments"`
	}

	//TODO
	// Don't write here.
	a := &alias{
		Channel:  s.channel,
		Username: s.username,
		Text:     "",
		Attachments: []attachment{
			{
				Color: "danger",
				Fields: []field{
					{
						Title: "command message",
						Value: s.text,
						Short: true,
					},
				},
			},
		},
	}
	return json.Marshal(a)
}

func NewSlack(mention string, channel string, endpoint string, username string) *Slack {
	return &Slack{
		Mention:  mention,
		Channel:  channel,
		Endpoint: endpoint,
		Username: username,
	}
}

func (s Slack) Notify(m message.Message) ([]byte, error) {
	fixedMessage := fmt.Sprintf("<%s>\n%s", s.Mention, string(m.PrettyJson()))

	param := newSlackMessage(s.Channel, s.Username, fixedMessage)
	b, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	return http.JsonPost(s.Endpoint, b)
}
