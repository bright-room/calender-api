DOCKER_RUN ?= docker compose --env-file .env --env-file .env.local

default: .go/install

up:
	$(DOCKER_RUN) --profile default up -d

down:
	$(DOCKER_RUN) --profile default down -v

restart: down up

rmi:
	$(DOCKER_RUN) --profile all down --rmi all --remove-orphans

pre-push: fmt lint test

fmt:
	cd ./backend && \
	golangci-lint fmt ./...

lint:
	cd ./backend && \
	golangci-lint run ./...

test:
	cd ./backend && \
	gotestsum --junitfile report.xml --format testname -- -cover -coverprofile=coverage.out -short ./...

.go/install:
	cd ./backend && \
	go install ./cmd/...

go/mod/tidy:
	cd ./backend && \
	go mod tidy

db/generate/%:
	$(DOCKER_RUN) --profile migrate run --remove-orphans migrate-generator $(@F)

db/up:
	$(DOCKER_RUN) --profile migrate run --remove-orphans migrate-up up

db/down:
	$(DOCKER_RUN) --profile migrate run --remove-orphans migrate-up down
