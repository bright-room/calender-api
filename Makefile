default: .go/install

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
	migrate create -ext sql -dir /migrations -seq $(@F)

db/up:
	migrate \
      -path "database/migrations" \
      -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" \
      up

db/down:
	migrate \
	  -path "database/migrations" \
	  -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" \
	  down
