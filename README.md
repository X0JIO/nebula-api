# Nebula API

Backend коммерческого VPN-сервиса Nebula.

## Требования

- Go 1.25+
- Docker
- Docker Compose

## Запуск

```bash
cp .env.example .env

docker compose up -d

go mod tidy

make run
```

На текущем этапе сервер еще не реализован — он будет добавлен в следующем коммите.