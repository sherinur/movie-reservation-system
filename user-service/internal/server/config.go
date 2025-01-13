package server

type config struct {
	Port   string
	DbUri  string
	DbName string
}

func NewConfig(port string) *config {
	// portValue := port
	// if port == "8080" {
	// 	portValue = os.Getenv("PORT")
	// }

	return &config{
		Port:   ":" + port,
		DbUri:  "mongodb://localhost:27017",
		DbName: "userDB",
	}
}
