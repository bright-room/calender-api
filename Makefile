default: .go/install

pre-push: fmt lint test

fmt:
	golangci-lint fmt ./...

lint:
	golangci-lint run ./...

test:
	gotestsum --junitfile report.xml --format testname -- -cover -coverprofile=coverage.out -short ./internal/...

.go/install:
	go install ./cmd/...

go/mod/tidy:
	go mod tidy
