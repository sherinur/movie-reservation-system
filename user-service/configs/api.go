package configs

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	Port                 string
	DbUri                string
	DbName               string
	JwtAccessSecret      string
	JwtRefreshSecret     string
	JwtAccessExpiration  int
	JwtRefreshExpiration int
	GoEnv                string
}

func GetConfig() *Config {
	once.Do(func() {
		config, err := ParseEnvConfig()
		if err != nil {
			log.Fatalf("Error of parsing environment variables: %s", err.Error())

			// log.Warn("Failed to load config. Using default values.")
			// instance = GetDefaultConfig()
		} else {
			instance = config
		}
	})

	return instance
}

// func GetDefaultConfig() *Config {
// 	return &Config{
// 		Port:         ":8080",
// 		DbUri:        "mongodb://localhost:27017",
// 		DbName:       "userDB",
// 		JwtSecretKey: "a5d52d1471164c78450ee0f6095cfN2f2c712e45525010b0e46e936cc61e6d205",
// 		ExpHours:     "1440",
// 		GoEnv:        "dev",
// 	}
// }

func ParseEnvConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	var (
		port                    = os.Getenv("PORT")
		mongoUri                = os.Getenv("MONGO_URI")
		mongoDbName             = os.Getenv("DB_NAME")
		jwtAccessSecret         = os.Getenv("JWT_ACCESS_SECRET")
		jwtRefreshSecret        = os.Getenv("JWT_REFRESH_SECRET")
		jwtAccessExpirationStr  = os.Getenv("JWT_ACCESS_EXPIRATION")
		jwtRefreshExpirationStr = os.Getenv("JWT_REFRESH_EXPIRATION")
		goEnv                   = os.Getenv("GO_ENV")
	)

	requiredEnvVars := []string{port, mongoUri, mongoDbName, jwtAccessSecret, jwtRefreshSecret, goEnv, jwtRefreshExpirationStr, jwtAccessExpirationStr}

	for _, envVar := range requiredEnvVars {
		if envVar == "" {
			return nil, ErrInvalidEnv
		}
	}

	jwtAccessExpiration, err := strconv.Atoi(jwtAccessExpirationStr)
	if err != nil {
		return nil, err
	}

	jwtRefreshExpiration, err := strconv.Atoi(jwtRefreshExpirationStr)
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:                 ":" + port,
		DbUri:                mongoUri,
		DbName:               mongoDbName,
		JwtAccessSecret:      jwtAccessSecret,
		JwtRefreshSecret:     jwtRefreshSecret,
		JwtAccessExpiration:  jwtAccessExpiration,
		JwtRefreshExpiration: jwtRefreshExpiration,
		GoEnv:                goEnv,
	}, nil
}
