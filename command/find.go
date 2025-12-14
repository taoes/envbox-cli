package command

import (
	"encoding/json"
	"envbox/model"
	"envbox/utils"
	"fmt"

	"github.com/spf13/cobra"
)

func InitFindCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "find [name]",
		Short: "搜索安装包",
		Run:   execCommand,
	}
	return command
}

func execCommand(cmd *cobra.Command, args []string) {
	// 从配置中获取注册中心URL
	config := utils.ReadConfig()
	registryUrl := config.RegistryUrl

	// 添加 '/api/package' 路径
	url := registryUrl + "/api/package"

	// 准备查询参数
	params := make(map[string]string)
	params["packageName"] = args[0]
	params["packageVersion"] = args[0]

	// 发送HTTP GET请求
	response, err := utils.Get(url, params)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}

	// 将响应字符串反序列化为CommonList<Package>结构体
	var packageList model.CommonList[model.Package]
	err = json.Unmarshal([]byte(response), &packageList)
	if err != nil {
		fmt.Printf("Error unmarshaling response: %v\n", err)
		return
	}

	// 打印结果
	if !packageList.Success {
		fmt.Printf("查询失败: %s\n", packageList.Message)
		return
	}
	fmt.Printf("找到 %d 个包:\n", len(packageList.Data))
	for _, pkg := range packageList.Data {
		fmt.Printf("- %s(%s): %s\n", pkg.Name, pkg.Version, pkg.Url)
	}

}
