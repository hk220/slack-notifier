package converter

import "github.com/hk220/slack-sender/message"

type Converter interface {
	Convert(msg string) *message.Message
}
