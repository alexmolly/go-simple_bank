package main

import (
	"database/sql"
	"log"

	"github.com/alexmolly/simple_bank/api"
	db "github.com/alexmolly/simple_bank/db/sqlc"
	"github.com/alexmolly/simple_bank/util"
	_ "github.com/lib/pq"
)

func main() {

	var conn *sql.DB
	var err error

	var config util.Config

	config, err = util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config file:", err)
	}

	conn, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	var store db.Store = db.NewStore(conn)
	var server *api.Server = api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
