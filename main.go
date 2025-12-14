package main

import (
	"envbox/command"
	"envbox/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	utils.ReadConfig()
	rootCmd := buildRootCmd()

	findCommand := command.InitFindCommand()
	installCommand := command.InitInstallCommand()
	listCommand := command.InitListCommand()
	startCommand := command.InitStartCommand()
	stopCommand := command.InitStopCommand()
	uninstallCommand := command.InitUninstallCommand()

	rootCmd.AddCommand(findCommand)
	rootCmd.AddCommand(installCommand)
	rootCmd.AddCommand(listCommand)
	rootCmd.AddCommand(startCommand)
	rootCmd.AddCommand(stopCommand)
	rootCmd.AddCommand(uninstallCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func buildRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{Use: "envBox", Short: "🔥 envBox CLI 主程序,致力于开发环境的管理、运行 & 分发"}

	return rootCmd
}
