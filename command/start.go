package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func InitStartCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "start [name]",
		Short: "启动指定服务",
		Run:   execStartCommand,
	}
	command.Flags().StringVarP(&name, "name", "n", "", "包名称")
	command.Flags().StringVarP(&version, "version", "v", "", "包版本")
	command.MarkFlagRequired("name")
	return command
}

func execStartCommand(cmd *cobra.Command, args []string) {
	name := "default-service"
	if len(args) > 0 {
		name = args[0]
	}
	fmt.Printf("Starting service %s...\n", name)
}
