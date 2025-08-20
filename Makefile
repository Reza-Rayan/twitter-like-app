APP_NAME=twitter-like-app

# Default target
.PHONY: run
run:
	go run ./cmd/

# Build binary
.PHONY: build
build:
	go build -o bin/$(APP_NAME) ./cmd/

# Run built binary
.PHONY: start
start: build
	./bin/$(APP_NAME)

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf bin
