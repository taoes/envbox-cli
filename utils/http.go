package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// Get 发送HTTP GET请求并返回响应体内容
func Get(rawURL string, params map[string]string) (string, error) {
	// 解析URL
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}

	// 构建查询参数
	if len(params) > 0 {
		query := u.Query()
		for key, value := range params {
			query.Set(key, value)
		}
		u.RawQuery = query.Encode()
	}

	// 发送GET请求
	resp, err := http.Get(u.String())
	if err != nil {
		return "", fmt.Errorf("failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed with status: %s", resp.Status)
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

// ProgressReader 包装原始reader以跟踪进度
type ProgressReader struct {
	reader     io.Reader
	total      int64
	current    int64
	onProgress func(current, total int64)
}

// Read 实现io.Reader接口
func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	pr.current += int64(n)

	if pr.onProgress != nil {
		pr.onProgress(pr.current, pr.total)
	}

	return n, err
}

// Download 下载文件到指定目录，并提供进度回调
func Download(downloadUrl string, destDir string, onProgress func(current, total int64)) error {
	// 创建目标目录
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// 获取文件名
	fileURL, err := url.Parse(downloadUrl)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}

	filename := filepath.Base(fileURL.Path)
	if filename == "." || filename == "/" {
		return fmt.Errorf("cannot determine filename from URL")
	}

	destPath := filepath.Join(destDir, filename)

	// 发起HTTP请求
	resp, err := http.Get(downloadUrl)
	if err != nil {
		return fmt.Errorf("failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %s", resp.Status)
	}

	// 获取文件大小
	var total int64
	if resp.Header.Get("Content-Length") != "" {
		fmt.Sscanf(resp.Header.Get("Content-Length"), "%d", &total)
	}

	// 创建进度追踪器
	progressReader := &ProgressReader{
		reader:     resp.Body,
		total:      total,
		onProgress: onProgress,
	}

	// 创建目标文件
	out, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// 复制数据
	_, err = io.Copy(out, progressReader)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
