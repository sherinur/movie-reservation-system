package server

type config struct {
	Port   string
	DBuri  string
	DBname string
}

func NewConfig(port string) *config {
	return &config{
		Port:   ":" + port,
		DBuri:  "mongodb://localhost:27017",
		DBname: "reservationDB",
	}
}
