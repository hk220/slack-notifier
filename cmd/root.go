package cmd

import (
	"os"

	"github.com/hk220/slack-notifier/converter"
	"github.com/hk220/slack-notifier/notifier"
	"github.com/hk220/slack-notifier/reader"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use: "slack-notifier",
		RunE: func(cmd *cobra.Command, args []string) error {
			n := notifier.NewSlackNotifier(viper.GetString("token"))
			r := reader.NewStdinReader()
			c := converter.NewSlackConverter(viper.GetString("username"), viper.GetString("channel"))

			sc := NewNotifyCommand(n, r, c)

			return sc.Execute()
		},
	}
)

type SendCommand struct {
	Notifier  notifier.Notifier
	Reader    reader.Reader
	Converter converter.Converter
}

func NewNotifyCommand(n notifier.Notifier, r reader.Reader, c converter.Converter) *SendCommand {
	return &SendCommand{
		Notifier:  n,
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
	if err := ss.Notifier.Notify(msg); err != nil {
		return err
	}
	return nil
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "slack-notifier.toml", "config file")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.PersistentFlags().String("token", "", "token")
	rootCmd.PersistentFlags().String("username", "", "username")
	rootCmd.PersistentFlags().String("channel", "", "channel")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("channel", rootCmd.PersistentFlags().Lookup("channel"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".slack-notifier")
	}

	err := viper.ReadInConfig()
	cobra.CheckErr(err)
}
