package converter

import "github.com/hk220/slack-sender/message"

type SlackConverter struct {
	Username string
	Channel  string
}

func NewSlackConverter(username string, channel string) *SlackConverter {
	return &SlackConverter{Username: username, Channel: channel}
}

func (sm *SlackConverter) Convert(msg string) *message.Message {
	return &message.Message{
		Text:     msg,
		UserName: sm.Username,
		Channel:  sm.Channel,
	}
}
