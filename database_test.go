package dalgo2sql

import (
	"database/sql"
	_ "github.com/proullon/ramsql/driver"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	database := NewDatabase(db)
	if database == nil {
		t.Fatal("NewDatabase(db) returned nil")
	}
}

func openTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("ramsql", "TestNewDatabase")
	if err != nil {
		t.Fatalf("sql.Open : Error : %s\n", err)
	}
	batch := []string{
		"CREATE TABLE TestKind (id VARCHAR(10) PRIMARY KEY, StringProp TEXT, IntProp INT);",
	}
	for _, b := range batch {
		_, err := db.Exec(b)
		if err != nil {
			t.Fatalf("sql.Exec: Error: %s\n", err)
		}
	}
	return db
}
