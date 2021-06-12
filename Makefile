.PHONY: test/inquiry test/category key

deps:
	go install github.com/cespare/reflex

up:
	-docker-compose up -d
	reflex -r '\.go|config.yml\z' -s -- sh -c 'go run cmd/main.go cmd/config.go cmd/router.go'

down:
	-docker-compose down

ps:
	-docker-compose -f docker-compose.yml -f docker-compose.prod.yml ps

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

key/build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o app_key_generator app_key_generator.go

key:
	go run app_key_generator.go
