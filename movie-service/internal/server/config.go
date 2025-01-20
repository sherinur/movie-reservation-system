package server

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

var ErrInvalidEnv = errors.New("missing environment variables")

type config struct {
	Port      string
	DbUri     string
	DbName    string
	SecretKey string
}

func NewConfig() *config {
	config, err := GetEnvConfig()
	if err != nil {
		return GetDefaultConfig()
	}

	return config
}

func GetDefaultConfig() *config {
	return &config{
		Port:      ":8080",
		DbUri:     "mongodb://localhost:27017",
		DbName:    "movieDB",
		SecretKey: "a5d52d1471164c78450ee0f6095cfN2f2c712e45525010b0e46e936cc61e6d205",
	}
}

func GetEnvConfig() (*config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	var (
		port         = os.Getenv("PORT")
		mongoUri     = os.Getenv("MONGO_URI")
		mongoDbName  = os.Getenv("DB_NAME")
		jwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	)

	if port == "" || mongoUri == "" || mongoDbName == "" || jwtSecretKey == "" {
		return nil, ErrInvalidEnv
	}

	return &config{
		Port:      ":" + port,
		DbUri:     mongoUri,
		DbName:    mongoDbName,
		SecretKey: jwtSecretKey,
	}, nil
}
