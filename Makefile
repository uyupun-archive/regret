init:
	cp .env.example .env
	go install github.com/cespare/reflex

db:
	docker-compose up -d
	make migrate

api:
	reflex -r '\.go|config.yml\z' -s -- sh -c 'go run cmd/main.go cmd/router.go'

ps:
	docker compose ps

sh:
	docker compose exec mysql bash

migrate:
	go run db/migrate.go
