package main

import (
	"database/sql"
	"log"

	"github.com/KaungthuKhant/simplebank/api"
	db "github.com/KaungthuKhant/simplebank/db/sqlc"
	"github.com/KaungthuKhant/simplebank/util"
	_ "github.com/lib/pq"
)

// to create server, we need to connect to the database and create a store
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// establish connection to the database
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
