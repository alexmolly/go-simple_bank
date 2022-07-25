package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/alexmolly/simple_bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	log.Println("--------------------Running TestMain")

	var conn *sql.DB
	var err error
	var config util.Config

	config, err = util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config file:", err)
	}

	conn, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
