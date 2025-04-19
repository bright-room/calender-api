default: .go/install

fmt: go/fmt

lint: go/lint

test: go/ut

.go/install:
	go install ./cmd/...

go/fmt:
	gofumpt -l -w -extra ./

go/lint:
	golangci-lint run ./...

go/ut:
	gotestsum --junitfile report.xml --format testname -- -cover -coverprofile=coverage.out -short ./internal/...

go/mod/tidy:
	go mod tidy
