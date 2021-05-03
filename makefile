# run a docker container for postgres
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13.2-alpine

# create the database before we can migrate it
createdb:
	docker exec -it postgres createdb --username=root --owner=root postgresdb

dropdb:
	docker exec -it postgres dropdb postgresdb

# migration using https://github.com/golang-migrate/migrate
migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/postgresdb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/postgresdb?sslmode=disable" -verbose down

# Turn SQL queries into Go code by using sqlc
# if you have sqlc installed, just run 'sqlc generate' 
# this one uses Docker images version of sqlc
sqlc:
	docker run --rm -v $(pwd):/src -w //src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown generate test server

# Initialize project by:
# make postgres
# make createdb
# make migrateup
# make server