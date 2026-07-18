APP=nebula-api

run:
	go run ./cmd/api

build:
	go build -o bin/$(APP) ./cmd/api

fmt:
	go fmt ./...

test:
	go test ./...

tidy:
	go mod tidy

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f