init:
	cp .env.example .env
	go install github.com/cespare/reflex

db:
	docker-compose up -d
	make migrate/up

api:
	reflex -r '\.go|config.yml\z' -s -- sh -c 'go run cmd/main.go cmd/config.go cmd/router.go'

ps:
	docker compose ps

sh:
	docker compose exec mysql bash

migrate/up:
	go run database/migrations/migrate.go up

migrate/down:
	go run database/migrations/migrate.go down

migrate/fresh:
	make migrate/down
	make migrate/up

seed:
	make migrate/fresh
	go run database/seeds/*.go

test/inquiry:
	make seed
	go run tests/inquiry/main.go

test/category:
	go run tests/category/main.go

key:
	go run app_key_generator.go
