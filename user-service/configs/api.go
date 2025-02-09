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
	Port             string
	DbUri            string
	DbName           string
	JwtAccessSecret  string
	JwtRefreshSecret string
	JwtExpiration    int
	GoEnv            string
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
		port             = os.Getenv("PORT")
		mongoUri         = os.Getenv("MONGO_URI")
		mongoDbName      = os.Getenv("DB_NAME")
		jwtAccessSecret  = os.Getenv("JWT_ACCESS_SECRET")
		jwtRefreshSecret = os.Getenv("JWT_REFRESH_SECRET")
		jwtExpirationStr = os.Getenv("JWT_EXPIRATION")
		goEnv            = os.Getenv("GO_ENV")
	)

	if port == "" || mongoUri == "" || mongoDbName == "" || jwtAccessSecret == "" || jwtRefreshSecret == "" || goEnv == "" || jwtExpirationStr == "" {
		return nil, ErrInvalidEnv
	}

	jwtExpiration, err := strconv.Atoi(jwtExpirationStr)
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:             ":" + port,
		DbUri:            mongoUri,
		DbName:           mongoDbName,
		JwtAccessSecret:  jwtAccessSecret,
		JwtRefreshSecret: jwtRefreshSecret,
		JwtExpiration:    jwtExpiration,
		GoEnv:            goEnv,
	}, nil
}
