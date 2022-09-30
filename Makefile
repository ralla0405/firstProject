mod:
	go mod tidy

run: 	mod
	go run main.go

build:
	go build -o build/server main.go

run-test:
	go run ./test/main.go

createdb:
	createdb --username=postgres --owner=postgres simple_bank

dropdb:
	dropdb simple_bank

migrationup:
	migrate -path db/migration -database "postgres://postgres:@Demian2020@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrationdown:
	migrate -path db/migration -database "postgres://postgres:@Demian2020@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrationup migrationdown