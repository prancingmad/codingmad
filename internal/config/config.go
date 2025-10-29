package config

import (
	"log"
	"os"
)

type Config struct {
	Port		string
	DBPath		string
	OllamaURL	string
}

func Load() *Config {
	cfg := &Config{
		Port:		getEnv("PORT", "8080"),
		DBPath:		getEnv("DB_PATH", "./data/codingmad.db"),
		OllamaURL:	getEnv("OLLAMA_URL", "http://localhost:11434"),
	}

	log.Println("Configuration loaded successfully")
	return cfg
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}