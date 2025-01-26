package server

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Port   string
	DBuri  string
	DBname string
}

func NewConfig() *config {
	conf, err := CreateEnvConfig()
	if err != nil {
		return DefaultConfig()
	}

	return conf
}

func DefaultConfig() *config {
	return &config{
		Port:   ":8080",
		DBuri:  "mongodb://localhost:27017",
		DBname: "reservationDB",
	}
}

func CreateEnvConfig() (*config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	conf := &config{
		Port:   ":" + os.Getenv("PORT"),
		DBuri:  os.Getenv("DB_URI"),
		DBname: os.Getenv("DB_NAME"),
	}

	if conf.Port == "" || conf.DBuri == "" || conf.DBname == "" {
		return nil, errors.New("env variables not set")
	}

	return conf, nil
}
