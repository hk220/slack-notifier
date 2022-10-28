package notifier

import "github.com/hk220/slack-notifier/message"

type Notifier interface {
	Notify(*message.Message) error
}
