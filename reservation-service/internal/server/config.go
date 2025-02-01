package server

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Port      string
	DBuri     string
	DBname    string
	SecretKey string
	ExpHours  string
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
		Port:      ":8080",
		DBuri:     "mongodb://localhost:27017",
		DBname:    "reservationDB",
		SecretKey: "a5d52d1471164c78450ee0f6095cfN2f2c712e45525010b0e46e936cc61e6d205",
		ExpHours:  "1440",
	}
}

func CreateEnvConfig() (*config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, ErrEnvVar
	}

	var (
		port      = os.Getenv("PORT")
		dburi     = os.Getenv("MONGO_URI")
		dbname    = os.Getenv("DB_NAME")
		secretkey = os.Getenv("JWT_SECRET_KEY")
		exphours  = os.Getenv("EXP_HOURS")
	)

	if port == "" || dburi == "" || dbname == "" || secretkey == "" || exphours == "" {
		return nil, ErrEnvVar
	}

	conf := &config{
		Port:      ":" + port,
		DBuri:     dburi,
		DBname:    dbname,
		SecretKey: secretkey,
		ExpHours:  exphours,
	}

	return conf, nil
}
