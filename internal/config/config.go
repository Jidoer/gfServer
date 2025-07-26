package config

import "time"

// Config 代理配置
type Config struct {
	Transport string        // 传输方式: "tcp" / "websocket" / "nat" / "custom"
	Listen    string        // 监听地址
	Target    string        // 目标服务器地址
	Timeout   time.Duration // 连接超时
}

// GetDefaultConfig 返回默认配置
func GetDefaultConfig() *Config {
	return &Config{
		Transport: "websocket",
		Listen:    "127.0.0.1:8080",
		Target:    "127.0.0.1:9000",
		Timeout:   10 * time.Second,
	}
}
