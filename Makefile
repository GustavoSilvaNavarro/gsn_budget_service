#!make
PORT = 8080
SERVICE_NAME = gsn_manager_service
CONTAINER_NAME = $(SERVICE_NAME)
DOCKER_COMPOSE_TAG = $(SERVICE_NAME)_1
TICKET_PREFIX := $(shell git branch --show-current | cut -d '_' -f 1)
DEFAULT_DB_URL := postgres://postgres:password@localhost:5432/budget_gsn?sslmode=disable

# App Commands
start:
	go run ./cmd/main.go

dev:
	air

unit:
	@echo "🏃‍♂️ Running Unit Tests..."
	go test -v ./tests/unit/...

unit-pretty:
	gotestsum --format short-verbose ./tests/unit/...

clean-cache:
	go clean -modcache

# Format code
format:
	@echo "Formatting code..."
	go fmt ./...
	@echo "✅ Code formatted!"

lint:
	@echo "Running linter..."
	golangci-lint run
	@echo "✅ Linting completed!"

# Migration Commands
# note: need to install golang-migrate
create-migration:
	migrate create -ext sql -dir migrations '$(m)'

# note to make this work need to install
# go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate-up:
	@echo "Migrating database at: $${DB_URL:-$(DEFAULT_DB_URL)}"
	migrate -path ./migrations -database "$${DB_URL:-$(DEFAULT_DB_URL)}" up

migrate-down:
	@echo "Downgrading database migration at: $${DB_URL:-$(DEFAULT_DB_URL)}"
	migrate -path ./migrations -database "$${DB_URL:-$(DEFAULT_DB_URL)}" down

create-queries:
	@echo "Creating types for my sql queries..."
	sqlc generate

# DB Commands
run-external-services:
	@DOCKER_BUILDKIT=1 docker compose -f ./docker-compose.inf.yml up -d db

# Docker commands
.PHONY: build-base
build-base:
	@COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker buildx build -f Dockerfile.base -t $(SERVICE_NAME)_base .

.PHONY: up
up: build-base
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker compose -f ./docker-compose.yaml -f ./docker-compose.inf.yml build --parallel
	docker compose -f ./docker-compose.yaml -f ./docker-compose.inf.yml up -d --force-recreate

.PHONY: down-rm
down-rm:
	docker compose -f ./docker-compose.yml -f ./docker-compose.inf.yml down --remove-orphans --rmi all --volumes
