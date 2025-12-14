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
	command.Flags().StringVarP(&name, "name", "n", "", "包名称")
	command.Flags().StringVarP(&version, "version", "v", "", "包版本")
	command.MarkFlagRequired("name")
	return command
}

func execStopCommand(cmd *cobra.Command, args []string) {
	name := "default-service"
	if len(args) > 0 {
		name = args[0]
	}
	fmt.Printf("Stopping service %s...\n", name)
}
