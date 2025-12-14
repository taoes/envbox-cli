package command

import (
	"github.com/spf13/cobra"
)

func InitInstallCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "install [name]",
		Short: "安装指定包",
		Run:   execInstallCommand,
	}
	return command
}

func execInstallCommand(cmd *cobra.Command, args []string) {

}
