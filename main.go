package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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

	// dbsource := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable", config.DBUser, config.DBPassword, config.DBHost, config.DBName)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode)

	conn, err = sql.Open(config.DBDriver, psqlInfo)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// err = server.Start(config.ServerAddress)
	err = server.Start(config.ServerAddress + ":" + config.Port)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
