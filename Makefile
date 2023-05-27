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
	docker run --name dev-mysql -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=dbauthor -e MYSQL_USER=admin -e MYSQL_PASSWORD=secret -p 3306:3306 -d mysql

dep-down:
	docker stop dev-mysql
	docker rm dev-mysql

gen-sqlc:
	sqlc generate

migrate-up:
	migrate -source file://./db/migrations -database "mysql://admin:secret@tcp(localhost:3306)/dbauthor" -verbose up 1

migrate-down:
	migrate -source file://./db/migrations -database "mysql://admin:secret@tcp(localhost:3306)/dbauthor" -verbose down 1

migrate-author:
	migrate create -ext sql -dir db/migrations -seq authors
infra-down:
	docker compose down