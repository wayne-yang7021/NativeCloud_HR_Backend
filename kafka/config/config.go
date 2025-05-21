package config

import (
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	}
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}
	MessageQueue struct {
		Type    string `yaml:"type"`
		Brokers string `yaml:"brokers"`
		Topic   string `yaml:"topic"`
		GroupID string `yaml:"group_id"`
	}
}

func LoadConfig() (*Config, error) {

	// 載入 .env 檔案
	err := godotenv.Load("/app/config/.env")
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile("/app/config/config.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// 覆蓋 YAML（可加上必要的 ENV fallback）
	if env := os.Getenv("KAFKA_BROKER"); env != "" {
		cfg.MessageQueue.Brokers = env
	}
	if topic := os.Getenv("KAFKA_TOPIC"); topic != "" {
		cfg.MessageQueue.Topic = topic
	}
	if group := os.Getenv("KAFKA_GROUP_ID"); group != "" {
		cfg.MessageQueue.GroupID = group
	}

	// 覆蓋 DB 用 .env
	if val := os.Getenv("DB_HOST"); val != "" {
		cfg.Database.Host = val
	}
	if val := os.Getenv("DB_PORT"); val != "" {
		// optional: parse string to int if needed
	}
	if val := os.Getenv("DB_USER"); val != "" {
		cfg.Database.User = val
	}
	if val := os.Getenv("DB_PASSWORD"); val != "" {
		cfg.Database.Password = val
	}
	if val := os.Getenv("DB_NAME"); val != "" {
		cfg.Database.Name = val
	}

	return &cfg, nil
}
