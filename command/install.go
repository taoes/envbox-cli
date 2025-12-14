package command

import (
	"envbox/utils"
	"fmt"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

func InitInstallCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "install [name]",
		Short: "安装指定包",
		Run:   execInstallCommand,
	}
	command.Flags().StringVarP(&name, "name", "n", "", "包名称")
	command.Flags().StringVarP(&version, "version", "v", "", "包版本")
	command.MarkFlagRequired("name")
	command.MarkFlagRequired("version")
	return command
}

func execInstallCommand(cmd *cobra.Command, args []string) {
	// 从配置中获取注册中心URL
	//config := utils.ReadConfig()
	//registryUrl := config.RegistryUrl

	// 添加 '/api/package' 路径
	//url := registryUrl + fmt.Sprintf("/api/package/%s/%s", name, version)

	var url string = "https://test-jpfile1.oss-cn-shenzhen.aliyuncs.com/Bom/bom/2022/1/19/2022011911355693652034.PDF"
	fmt.Printf("正在获取安装包信息%s(%s),请稍后....\n", name, version)

	bar := progressbar.Default(100, "下载进度")
	utils.Download(url, utils.GetHomeDir(), func(current, total int64) {
		bar.Set64(100 * current / total)
		time.Sleep(50 * time.Millisecond)
	})
}
