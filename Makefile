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


# Migrations
migrate-new:
	@migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

migrate-up:
	@migrate -database "$(DB_URL)" -path $(MIGRATIONS_DIR) up

migrate-down:
	@migrate -database "$(DB_URL)" -path $(MIGRATIONS_DIR) down 1

migrate-reset:
	@migrate -database "$(DB_URL)" -path $(MIGRATIONS_DIR) drop -f

