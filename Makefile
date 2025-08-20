APP_NAME=twitter-like-app
MIGRATE=migrate
DB_FILE=./twitter.db
MIGRATIONS=./migrations

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
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS) $(NAME)

migrate-up:
	$(MIGRATE) -path $(MIGRATIONS) -database sqlite://$(DB_FILE) up

migrate-down:
	$(MIGRATE) -path $(MIGRATIONS) -database sqlite3://$(DB_FILE) down -all

migrate-reset:
	@migrate -database "$(DB_URL)" -path $(MIGRATIONS_DIR) drop -f

