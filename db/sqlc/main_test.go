package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	// _ : the the formatter to keep it
	_ "github.com/lib/pq" // side-effect import; registers the Postgres driver with database/sql so sql.Open("postgres", …) works
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries // global variable to hold the queries object
var testDB *sql.DB

func TestMain(m *testing.M) { // TestMain is a special function that is run before any tests are executed
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
