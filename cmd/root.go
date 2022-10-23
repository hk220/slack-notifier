package cmd

import (
	"github.com/hk220/slack-sender/converter"
	"github.com/hk220/slack-sender/reader"
	"github.com/hk220/slack-sender/sender"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use: "slack-sender",
		RunE: func(cmd *cobra.Command, args []string) error {
			s := sender.NewSlackSender(viper.GetString("token"))
			r := reader.NewStdinReader()
			c := converter.NewSlackConverter(viper.GetString("username"), viper.GetString("channel"))

			sc := NewSendCommand(s, r, c)

			return sc.Execute()
		},
	}
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

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "slack-sender.toml", "config file")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.PersistentFlags().String("token", "", "token")
	rootCmd.PersistentFlags().String("username", "", "username")
	rootCmd.PersistentFlags().String("channel", "", "channel")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("channel", rootCmd.PersistentFlags().Lookup("channel"))
}
