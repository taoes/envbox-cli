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
	return command
}

func execListCommand(cmd *cobra.Command, args []string) {
	fmt.Println("Listing installed packages...")
	// 实际实现中这里会查询并显示已安装的包列表
}