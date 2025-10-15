package config

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

type ServerConfig struct {
	Port string
}

type SecurityConfig struct {
	JWTSecret string
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Security SecurityConfig
}

func Load() Config {
	cfg := Config{
		Database: DatabaseConfig{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
		},
		Server: ServerConfig{
			Port: valueOrDefault(os.Getenv("APP_PORT"), "8080"),
		},
		Security: SecurityConfig{
			JWTSecret: os.Getenv("JWT_SECRET"),
		},
	}

	return cfg
}

func (d DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.User, d.Password, d.Host, d.Port, d.Name)
}

func valueOrDefault(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
