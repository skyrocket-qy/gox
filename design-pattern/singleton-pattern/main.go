package main

import (
	"os"
	"sync"
)

type AppConfig struct {
	DatabaseURL string
	APIKey      string
	// ... 其他設定欄位
}

var (
	instance *AppConfig
	once     sync.Once
)

func GetConfig() *AppConfig {
	once.Do(func() {
		instance = &AppConfig{
			DatabaseURL: os.Getenv("DATABASE_URL"),
			APIKey:      os.Getenv("API_KEY"),
			// ... 加載其他設定
		}
	})

	return instance
}
