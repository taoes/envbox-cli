package utils

import (
	"encoding/json"
	"envbox/model"
	"os"
	"path/filepath"
)

// ReadConfig 从文件中读取配置，如果文件不存在则初始化配置文件
func ReadConfig() model.Config {
	configPath := GetConfigPath()

	// 检查配置文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// 如果配置文件不存在，则初始化配置
		InitConfig()
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		// 如果读取失败，返回默认配置
		return model.Config{
			Version:     "1.0.0",
			Verbose:     false,
			DataDir:     filepath.Join(GetHomeDir(), ".envbox", "data"),
			LogsDir:     filepath.Join(GetHomeDir(), ".envbox", "logs"),
			RegistryUrl: "https://registry.envbox.local",
		}
	}

	// 解析配置文件
	var config model.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		// 如果解析失败，返回默认配置
		return model.Config{
			Version:     "1.0.0",
			Verbose:     false,
			DataDir:     filepath.Join(GetHomeDir(), ".envbox", "data"),
			LogsDir:     filepath.Join(GetHomeDir(), ".envbox", "logs"),
			RegistryUrl: "https://registry.envbox.local",
		}
	}

	return config
}

// InitConfig 初始化配置文件
func InitConfig() {
	configPath := GetConfigPath()

	// 确保配置目录存在
	configDir := filepath.Dir(configPath)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0755)
	}

	// 创建默认配置
	config := model.Config{
		Version:     "1.0.0",
		Verbose:     false,
		DataDir:     filepath.Join(GetHomeDir(), ".envbox", "data"),
		LogsDir:     filepath.Join(GetHomeDir(), ".envbox", "logs"),
		RegistryUrl: "https://registry.envbox.local",
	}

	// 将配置写入文件
	data, _ := json.MarshalIndent(config, "", "  ")
	os.WriteFile(configPath, data, 0644)
}
