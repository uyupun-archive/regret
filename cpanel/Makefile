.PHONY: up down

up:
	-docker compose --env-file ./.env.production up -d --build

down:
	-docker compose down

ps:
	docker compose ps
