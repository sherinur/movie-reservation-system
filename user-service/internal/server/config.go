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
			Log.Errorf("Error of parsing environment variables: %s", err.Error())
			Log.Warn("Failed to load config. Using default values.")
			instance = GetDefaultConfig()
		} else {
			instance = config
		}
	})

	return instance
}

func GetDefaultConfig() *Config {
	return &Config{
		Port:         ":8080",
		DbUri:        "mongodb://localhost:27017",
		DbName:       "userDB",
		JwtSecretKey: "a5d52d1471164c78450ee0f6095cfN2f2c712e45525010b0e46e936cc61e6d205",
		ExpHours:     "1440",
	}
}

func ParseEnvConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	var (
		port         = os.Getenv("PORT")
		mongoUri     = os.Getenv("MONGO_URI")
		mongoDbName  = os.Getenv("DB_NAME")
		jwtSecretKey = os.Getenv("JWT_SECRET_KEY")
		expHours     = os.Getenv("EXP_HOURS")
	)

	if port == "" || mongoUri == "" || mongoDbName == "" || jwtSecretKey == "" {
		return nil, ErrInvalidEnv
	}

	return &Config{
		Port:         ":" + port,
		DbUri:        mongoUri,
		DbName:       mongoDbName,
		JwtSecretKey: jwtSecretKey,
		ExpHours:     expHours,
	}, nil
}
