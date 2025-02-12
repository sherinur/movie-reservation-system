BINARY_NAME=movie-service

MAIN_FILE=cmd/main.go

test:
	go test -v ./internal/handler/
	
build:
	go mod tidy
	gofumpt -l -w .
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) $(MAIN_FILE)

run: build
	./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	go mod tidy
	gofumpt -l -w .
	rm data -rf
	rm -f $(BINARY_NAME)

deploy:
	make -C ../. deploy

help:
	@echo "Makefile commands:"
	@echo "  make build   - Build the project"
	@echo "  make deploy  - Docker compose"
	@echo "  make run     - Build and run the project"
	@echo "  make clean   - Remove the compiled binary"
	@echo "  make help    - Show this help message"

commit:
	git add .
	git commit -m "Commit $$(date '+%Y-%m-%d %H:%M:%S')"
	git push

.PHONY: build run clean help