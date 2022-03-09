.PHONY: test
test:
	go test ./...

.PHONY: build
build: test
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./app/bitme ./cmd
