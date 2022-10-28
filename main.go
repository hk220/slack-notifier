package main

import (
	"github.com/hk220/slack-notifier/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.Execute())
}
