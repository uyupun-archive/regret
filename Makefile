deps:
	cp .env.example .env
	go install github.com/cespare/reflex

dev:
	-docker-compose up -d
	reflex -r '\.go|config.yml\z' -s -- sh -c 'go run cmd/main.go cmd/config.go cmd/router.go'

migrate/up:
	go run database/migrations/migrate.go up

migrate/down:
	go run database/migrations/migrate.go down

migrate/fresh:
	make migrate/down
	make migrate/up

seed:
	go run database/seeds/*.go

test/inquiry:
	make seed
	go run tests/inquiry/main.go

test/category:
	go run tests/category/main.go

key:
	go run app_key_generator.go
