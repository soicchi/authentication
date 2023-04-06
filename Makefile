.PHONY: fmt build test vendor tidy

fmt:
	docker compose run --rm web go fmt ./...

build:
	docker compose run --rm web go build -o app ./cmd/...

test:
	docker compose run --rm web go test -v ./...

vendor:
	docker compose run --rm web go mod vendor

tidy:
	docker compose run --rm web go mod tidy