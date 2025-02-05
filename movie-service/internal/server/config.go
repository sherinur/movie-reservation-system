package server

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	Port         string
	DbUri        string
	DbName       string
	JwtSecretKey string
	ExpHours     string
}

func GetConfig() *Config {
	once.Do(func() {
		config, err := ParseEnvConfig()
		if err != nil {
			// log.Errorf("Error of parsing environment variables: %s", err.Error())
			// log.Warn("Failed to load config. Using default values.")
			instance = GetDefaultConfig()
		} else {
			instance = config
		}
	})

	return instance
}

func GetDefaultConfig() *Config {
	return &Config{
		Port:   ":8080",
		DbUri:  "mongodb://localhost:27017",
		DbName: "movieDB",
	}
}

func ParseEnvConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	var (
		port        = os.Getenv("PORT")
		mongoUri    = os.Getenv("MONGO_URI")
		mongoDbName = os.Getenv("DB_NAME")
	)

	if port == "" || mongoUri == "" || mongoDbName == "" {
		return nil, ErrInvalidEnv
	}

	return &Config{
		Port:   ":" + port,
		DbUri:  mongoUri,
		DbName: mongoDbName,
		// JwtSecretKey: jwtSecretKey,
		// ExpHours:     expHours,
	}, nil
}
