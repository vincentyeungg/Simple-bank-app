package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	// database driver
	dbDriver = "postgres"
	// database source
	dbSource = "postgresql://root:password@localhost:5435/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	// sql open need to know dbDriver and dbSource, which are the parameters we use to 
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// New is defined in the db.go file generated by sqlc
	testQueries = New(conn)

	// run the unit tests
	os.Exit(m.Run())
}