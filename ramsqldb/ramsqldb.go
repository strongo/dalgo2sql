package ramsqldb

import (
	"database/sql"
	_ "github.com/proullon/ramsql/driver"
	"testing"
)

func OpenTestDb(t *testing.T) *sql.DB {
	db, err := sql.Open("ramsql", "TestNewDatabase")
	if err != nil {
		t.Fatalf("sql.Open : Error : %s\n", err)
	}
	batch := []string{
		"CREATE TABLE E2ETest1 (id VARCHAR(10) PRIMARY KEY, StringProp TEXT, IntProp INT);",
		"CREATE TABLE NonExistingKind (id VARCHAR(10) PRIMARY KEY);",
	}
	for _, b := range batch {
		_, err := db.Exec(b)
		if err != nil {
			t.Fatalf("sql.Exec: Error: %s\n", err)
		}
	}
	return db
}
