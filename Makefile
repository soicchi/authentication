.PHONY: fmt build test vendor

fmt:
	go fmt ./...

build:
	go build -o app ./...

test:
	go test -v ./...

vendor:
	go mod vendor