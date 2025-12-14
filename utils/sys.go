package utils

import (
	"runtime"
)

// GetArch 获取 CPU 架构信息
func GetArch() string {
	return runtime.GOARCH
}

// GetOS 获取操作系统信息
func GetOS() string {
	return runtime.GOOS
}

// GetPlatform 获取平台信息 (操作系统+架构)
func GetPlatform() string {
	return runtime.GOOS + "/" + runtime.GOARCH
}

// IsARM 检查是否为 ARM 架构
func IsARM() bool {
	arch := runtime.GOARCH
	return arch == "arm" || arch == "arm64"
}

// IsX86 检查是否为 x86 架构
func IsX86() bool {
	arch := runtime.GOARCH
	return arch == "386" || arch == "amd64"
}

// Is64Bit 检查是否为 64 位架构
func Is64Bit() bool {
	arch := runtime.GOARCH
	return arch == "amd64" || arch == "arm64" || arch == "ppc64" || arch == "mips64" || arch == "s390x"
}