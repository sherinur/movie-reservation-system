package server

type config struct {
	Port   string
	DbUri  string
	DbName string
}

func NewConfig(port string) *config {
	return &config{
		Port:   port,
		DbUri:  "mongodb://localhost:27017",
		DbName: "userDB",
	}
}
