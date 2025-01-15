package server

type config struct {
	Port      string
	DbUri     string
	DbName    string
	SecretKey string
}

func NewConfig(port string) *config {
	// portValue := port
	// if port == "8080" {
	// 	portValue = os.Getenv("PORT")
	// }

	return &config{
		Port:      ":" + port,
		DbUri:     "mongodb://localhost:27017",
		DbName:    "userDB",
		SecretKey: "a5d52d1471164c78450ee0f6095cf2f2c712e45525010b0e46e936cc61e6d205",
	}
}
