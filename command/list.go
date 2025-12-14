package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func InitListCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "list",
		Short: "列出已安装的包",
		Run:   execListCommand,
	}
	command.Flags().StringVarP(&name, "name", "n", "", "包名称")
	command.Flags().StringVarP(&version, "version", "v", "", "包版本")
	command.MarkFlagRequired("name")
	return command
}

func execListCommand(cmd *cobra.Command, args []string) {
	fmt.Println("Listing installed packages...")
	// 实际实现中这里会查询并显示已安装的包列表
}
