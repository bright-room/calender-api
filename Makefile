default: .go/install

fmt:
	gofumpt -l -w -extra ./

lint:
	golangci-lint run ./...

test:
	gotestsum --junitfile report.xml --format testname -- -cover -coverprofile=coverage.out -short ./internal/...

.go/install:
	go install ./cmd/...

go/mod/tidy:
	go mod tidy
