vup:
	@go build -o bin .

.PHONY: test
test:
	@go test -cover ./...

.PHONY: install
install:
	@go install .

.PHONY: lint
lint:
	@go tool golangci-lint run
