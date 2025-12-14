package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func InitStopCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "stop [name]",
		Short: "停止指定服务",
		Run:   execStopCommand,
	}
	return command
}

func execStopCommand(cmd *cobra.Command, args []string) {
	name := "default-service"
	if len(args) > 0 {
		name = args[0]
	}
	fmt.Printf("Stopping service %s...\n", name)
}