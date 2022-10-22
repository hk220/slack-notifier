package message

import "testing"

func TestValidValidate(t *testing.T) {
	msg := &Message{
		Channel:  "#random",
		Text:     "にゃんぱす〜",
		UserName: "れんげ",
	}

	if err := msg.Validate(); err != nil {
		t.Fatal(err)
	}
}
