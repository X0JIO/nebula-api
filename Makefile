APP=nebula-api

.PHONY: run build tidy fmt test up down logs

run:
	go run ./cmd/api

build:
	go build -o bin/$(APP) ./cmd/api

tidy:
	go mod tidy

fmt:
	go fmt ./...

test:
	go test ./...

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f