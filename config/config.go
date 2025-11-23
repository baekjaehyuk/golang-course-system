package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Url  string
	Key  string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("env 파일을 불러오지 못했습니다: %v", err)
	}

	return &Config{
		Port: os.Getenv("PORT"),
		Url:  os.Getenv("SUPABASE_URL"),
		Key:  os.Getenv("SUPABASE_ANON_KEY"),
	}
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
