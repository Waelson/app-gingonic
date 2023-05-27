build:
	go build -v -o bin/main cmd/main.go

image-build:
	docker build -t app-gingonic/latest .

run:
	go run cmd/main.go

infra-up:
	docker compose up -d --build

infra-restart:
	docker compose restart

gen-sqlc:
	sqlc generate

infra-down:
	docker compose down