package message

import (
	"fmt"
	"strings"
)

type Message struct {
	UserName string
	Text     string
	Channel  string
}

func (msg *Message) Validate() error {
	var errors []string

	if msg.Channel == "" {
		errors = append(errors, "channel must contain a channel string (#****)")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, " "))
	}
	return nil
}
