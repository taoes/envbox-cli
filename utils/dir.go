package utils

import (
	"os"
	"path/filepath"
)

// GetConfigPath 获取配置文件路径
func GetConfigPath() string {
	homeDir := GetHomeDir()
	return filepath.Join(homeDir, ".config", "envbox", "config.json")
}

// GetHomeDir 获取用户主目录
func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		// 如果无法获取用户主目录，则使用当前工作目录
		homeDir, _ = os.Getwd()
	}
	return homeDir
}
