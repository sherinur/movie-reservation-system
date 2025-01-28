package server

import (
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
		return nil, ErrEnvVar
	}

	var (
		port   = os.Getenv("PORT")
		dbUri  = os.Getenv("MONGO_URI")
		dbName = os.Getenv("DB_NAME")
	)

	if port == "" || dbUri == "" || dbName == "" {
		return nil, ErrEnvVar
	}

	conf := &config{
		Port:   ":" + port,
		DBuri:  dbUri,
		DBname: dbName,
	}

	return conf, nil
}
