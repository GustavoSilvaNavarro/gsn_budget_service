#!make
PORT = 8080
SERVICE_NAME = gsn_manager_service
CONTAINER_NAME = $(SERVICE_NAME)
DOCKER_COMPOSE_TAG = $(SERVICE_NAME)_1
TICKET_PREFIX := $(shell git branch --show-current | cut -d '_' -f 1)

# App Commands
start:
	go run ./src/main.go

dev:
	air

unit:
	@echo "🏃‍♂️ Running Unit Tests..."
	go test -v ./tests/unit/...

unit-pretty:
	gotestsum --format short-verbose ./tests/unit/...

clean-cache:
	go clean -modcache

# DB Commands
run-external-services:
	docker compose -f ./docker-compose.inf.yml up -d mongodb  mongo-express

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
	docker compose -f ./docker-compose.yaml -f ./docker-compose.inf.yml down --remove-orphans --rmi all --volumes
