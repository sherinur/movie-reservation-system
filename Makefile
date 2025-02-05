USER_SERVICE_BIN=user-service
USER_SERVICE_MAIN=user-service/cmd/main.go

MOVIE_SERVICE_BIN=movie-service
MOVIE_SERVICE_MAIN=movie-service/cmd/main.go

RESERVATION_SERVICE_BIN=reservation-service
RESERVATION_SERVICE_MAIN=reservation-service/cmd/main.go

build:
	@echo "Building services..."
	make -C user-service/ build
	make -C movie-service/ build
	make -C reservation-service/ build
	@echo "Services built successfully."

deploy:
	@echo "Deploying the project..."
	make build
	docker compose -f ./docker-compose.yml  up --build -d
	# make -C web/ run