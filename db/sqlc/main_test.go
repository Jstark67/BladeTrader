package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:75Cheetah@localhost:5432/backend_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	deleteAllData(context.Background())
	os.Exit(m.Run())

}

func deleteAllData(ctx context.Context) {
	_, err := testQueries.db.ExecContext(ctx, "DELETE FROM entries")
	if err != nil {
		log.Fatalf("Failed to delete entries: %v", err)
	}

	// delete all data from "accounts" because it has a foreign key to "users"
	_, err = testQueries.db.ExecContext(ctx, "DELETE FROM transfers")
	if err != nil {
		log.Fatalf("Failed to delete transfers: %v", err)
	}

	// delete all data from "users" because it has no more dependencies
	_, err = testQueries.db.ExecContext(ctx, "DELETE FROM accounts")
	if err != nil {
		log.Fatalf("Failed to delete accounts: %v", err)
	}
}
