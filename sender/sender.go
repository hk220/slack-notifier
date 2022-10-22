package sender

import "github.com/hk220/slack-sender/message"

type Sender interface {
	Send(*message.Message) error
}
