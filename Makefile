vup:
	@go build .

.PHONY: test
test:
	@go test -cover ./...

.PHONY: install
install:
	@go install .

.PHONY: lint
lint:
	@go tool golangci-lint run
