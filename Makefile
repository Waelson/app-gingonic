build:
	go build -v -o bin/main cmd/main.go

image-build:
	docker build -t app-gingonic/latest .

run:
	go run cmd/main.go

infra-up:
	docker compose up -d # --build

infra-restart:
	docker compose restart

dep-up:
	docker run --name dev-mysql -e MYSQL_ROOT_PASSWORD=secret -d mysql

dep-down:
	docker stop dev-mysql
	docker rm dev-mysql

gen-sqlc:
	sqlc generate

infra-down:
	docker compose down