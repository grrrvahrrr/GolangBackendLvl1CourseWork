.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...

.PHONY: build
build: lint
build: test
	go build
