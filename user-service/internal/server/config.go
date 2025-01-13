package server

type config struct {
	Port string
}

func NewConfig(port string) *config {
	return &config{
		Port: port,
	}
}
