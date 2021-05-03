package main

import (
	"backend/api"
	db "backend/db/sqlc"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// in order to connect to a server, we have to connect to a DB and create a store first
const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/postgresdb?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	// establish a connection to DB
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("Cannot start server!", err)
	}
}
