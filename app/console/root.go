package console

import (
	"github.com/spf13/cobra"
	"orange_message_service/app/components/config"
	"orange_message_service/app/components/mlog"
	"orange_message_service/app/console/commands"
	"os"
)

var Commands = []*cobra.Command{
	commands.HelloCommand,
}

var RootCmd = &cobra.Command{
	SilenceErrors: true,
	Use:           "meicli",
	Short:         "run commands with meicli",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	config.Init()
	mlog.Init()

	for _, command := range Commands {
		RootCmd.AddCommand(command)
	}
}
