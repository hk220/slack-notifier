package cmd

import (
	"github.com/hk220/slack-sender/config"
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

func NewSendCommand(config *config.Config) *SendCommand {
	ss := &SendCommand{
		Sender:    sender.NewSlackSender(config.Token),
		Reader:    reader.NewStdinReader(),
		Converter: converter.NewSlackConverter(config.Username, config.Channel),
	}
	return ss
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
		config := &config.Config{
			Token:    viper.GetString("token"),
			Channel:  viper.GetString("channel"),
			Username: viper.GetString("username"),
		}

		SendCommand := NewSendCommand(config)

		if err := SendCommand.Execute(); err != nil {
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
