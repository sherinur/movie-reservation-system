module user-service

go 1.22.2

require go.mongodb.org/mongo-driver v1.17.2

require (
	github.com/joho/godotenv v1.5.1
	github.com/sherinur/movie-reservation-system/pkg/db v0.0.0-00010101000000-000000000000
	github.com/sherinur/movie-reservation-system/pkg/logging v0.0.0-00010101000000-000000000000
)

require (
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
)

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)

replace github.com/sherinur/movie-reservation-system/pkg/logging => ../pkg/logging

replace github.com/sherinur/movie-reservation-system/pkg/db => ../pkg/db
