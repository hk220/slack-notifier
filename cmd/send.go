package cmd

import (
	"github.com/hk220/slack-sender/converter"
	"github.com/hk220/slack-sender/reader"
	"github.com/hk220/slack-sender/sender"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type SendCommand struct {
	Sender    sender.Sender
	Reader    reader.Reader
	Converter converter.Converter
}

func NewSendCommand(s sender.Sender, r reader.Reader, c converter.Converter) *SendCommand {
	return &SendCommand{
		Sender:    s,
		Reader:    r,
		Converter: c,
	}
}

func (ss *SendCommand) Execute() error {
	text, err := ss.Reader.Read()
	if err != nil {
		return err
	}

	msg := ss.Converter.Convert(text)
	if err := ss.Sender.Send(msg); err != nil {
		return err
	}
	return nil
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send message to slack from stdin",
	RunE: func(cmd *cobra.Command, args []string) error {

		s := sender.NewSlackSender(viper.GetString("token"))
		r := reader.NewStdinReader()
		c := converter.NewSlackConverter(viper.GetString("username"), viper.GetString("channel"))

		sc := NewSendCommand(s, r, c)

		if err := sc.Execute(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	sendCmd.Flags().String("token", "", "token")
	sendCmd.Flags().String("username", "", "username")
	sendCmd.Flags().String("channel", "", "channel")
	viper.BindPFlag("token", sendCmd.Flags().Lookup("token"))
	viper.BindPFlag("username", sendCmd.Flags().Lookup("username"))
	viper.BindPFlag("channel", sendCmd.Flags().Lookup("channel"))
}
