package app

import (
	"os"
)

type DatabaseConfig struct {
	User     string
	Password string
	Address  string
}

type Config struct {
	RedisAddress string
	DatabaseInfo DatabaseConfig
}

// Required Configs:
// - REDIS_ADDR
// - DB_ADDRESS
// - DB_USER
// - DB_PASSWORD
// - WORKER_COUNT
func LoadConfig() Config {
	cfg := Config{
		RedisAddress: "localhost:6379",
		DatabaseInfo: DatabaseConfig{
			User:     "user",
			Password: "password",
			Address:  "localhost:3306"},
	}

	if redisAddr, exists := os.LookupEnv("REDIS_ADDR"); exists {
		cfg.RedisAddress = redisAddr
	}

	if dbAddress, exists := os.LookupEnv("DB_ADDRESS"); exists {
		cfg.DatabaseInfo.Address = dbAddress
	}

	if dbUser, exists := os.LookupEnv("DB_USER"); exists {
		cfg.DatabaseInfo.User = dbUser
	}

	if dbPassword, exists := os.LookupEnv("DB_PASSWORD"); exists {
		cfg.DatabaseInfo.Password = dbPassword
	}

	return cfg
}
