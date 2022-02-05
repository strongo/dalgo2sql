package ramsqldb

import (
	"database/sql"
	// Calling OpenTestDb will use RamSQL database
	_ "github.com/proullon/ramsql/driver"
	"testing"
)

// OpenTestDb opens a test database using RamSQL
func OpenTestDb(t *testing.T) *sql.DB {
	db, err := sql.Open("ramsql", "TestNewDatabase")
	if err != nil {
		t.Fatalf("sql.Open : Error : %s\n", err)
	}
	batch := []string{
		"CREATE TABLE DalgoE2E_E2ETest1 (ID1 VARCHAR(10) PRIMARY KEY, StringProp TEXT, IntegerProp INT);",
		"CREATE TABLE DalgoE2E_E2ETest2 (ID VARCHAR(10) PRIMARY KEY, StringProp TEXT, IntegerProp INT);",
		"CREATE TABLE NonExistingKind (ID VARCHAR(10) PRIMARY KEY, StringProp TEXT, IntegerProp INT);",
	}
	for _, b := range batch {
		_, err := db.Exec(b)
		if err != nil {
			t.Fatalf("sql.Exec: Error: %s\n", err)
		}
	}
	return db
}
