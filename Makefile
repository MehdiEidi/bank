DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate -f ./db/sqlc_conf/sqlc.yaml

test:
	go test -v -cover ./...

tidy:
	go mod tidy

server: 
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/mehdieidi/bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test tidy server mock