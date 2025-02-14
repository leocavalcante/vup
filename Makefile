vup:
	@CGO_ENABLED=0 go build -o bin .

.PHONY: test
test:
	@go test -cover ./...

.PHONY: install
install:
	@go install .

.PHONY: lint
lint:
	@go tool golangci-lint run

.PHONY: image
image:
	@docker build -t vup .
