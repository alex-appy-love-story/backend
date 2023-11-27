package app

import (
	"os"
)

type DatabaseConfig struct {
	User         string
	Password     string
	Address      string
	DatabaseName string
}

type Config struct {
	RedisAddress   string
	DatabaseConfig DatabaseConfig
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
		DatabaseConfig: DatabaseConfig{
			User:         "user",
			Password:     "password",
			Address:      "localhost:3306",
			DatabaseName: "db",
		},
	}

	if redisAddr, exists := os.LookupEnv("REDIS_ADDR"); exists {
		cfg.RedisAddress = redisAddr
	}

	if dbAddress, exists := os.LookupEnv("DB_ADDRESS"); exists {
		cfg.DatabaseConfig.Address = dbAddress
	}

	if dbUser, exists := os.LookupEnv("DB_USER"); exists {
		cfg.DatabaseConfig.User = dbUser
	}

	if dbPassword, exists := os.LookupEnv("DB_PASSWORD"); exists {
		cfg.DatabaseConfig.Password = dbPassword
	}

	if dbName, exists := os.LookupEnv("DB_NAME"); exists {
		cfg.DatabaseConfig.DatabaseName = dbName
	}

	return cfg
}
