package converter

import "github.com/hk220/slack-notifier/message"

type Converter interface {
	Convert(msg string) *message.Message
}
