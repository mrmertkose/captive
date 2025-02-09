css:
	@npx tailwindcss -i assets/css/input.css -o assets/css/styles.min.css --watch --minify

run:
	air

templ:
	@templ generate --watch

migrate:
	go run pkg/database/migration/migration.go

build:
	@npx tailwindcss -i assets/css/input.css -o assets/css/styles.min.css --minify
	templ generate view
	go build -ldflags "-s -w" -o bin/main.exe .

up:
	@echo Starting Docker images...
	docker-compose up -d
	@echo Docker images started!

down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!

.PHONY: css run templ migrate build up down