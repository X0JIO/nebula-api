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

DB_URL=postgres://nebula:nebula@localhost:5432/nebula?sslmode=disable


migrate-up:
	goose -dir migrations postgres "$(DB_URL)" up


migrate-down:
	goose -dir migrations postgres "$(DB_URL)" down


migrate-status:
	goose -dir migrations postgres "$(DB_URL)" status

sqlc:
	sqlc generate