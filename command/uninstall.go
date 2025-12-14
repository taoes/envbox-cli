package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func InitUninstallCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "uninstall [name]",
		Short: "卸载指定包",
		Run:   execUninstallCommand,
	}
	command.Flags().StringVarP(&name, "name", "n", "", "包名称")
	command.Flags().StringVarP(&version, "version", "v", "", "包版本")
	command.MarkFlagRequired("name")
	return command
}

func execUninstallCommand(cmd *cobra.Command, args []string) {
	name := "default-package"
	if len(args) > 0 {
		name = args[0]
	}
	fmt.Printf("Uninstalling %s...\n", name)
}
