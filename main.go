package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mehdieidi/bank/api"
	db "github.com/mehdieidi/bank/db/sqlc"
	"github.com/mehdieidi/bank/pkg/config"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
