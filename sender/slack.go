package sender

import (
	"context"

	"github.com/hk220/slack-sender/message"
	"github.com/slack-go/slack"
)

type SlackSender struct {
	Token string
}

func NewSlackSender(token string) *SlackSender {
	return &SlackSender{Token: token}
}

func (s *SlackSender) Send(msg *message.Message) error {
	if err := msg.Validate(); err != nil {
		return err
	}

	client := slack.New(s.Token)

	var options []slack.MsgOption

	if msg.UserName != "" {
		options = append(options, slack.MsgOptionUsername(msg.UserName))
	}

	if msg.Text != "" {
		options = append(options, slack.MsgOptionText(msg.Text, false))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if _, _, err := client.PostMessageContext(ctx, msg.Channel, options...); err != nil {
		return err
	}

	return nil
}
