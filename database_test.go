package dalgo2sql

import (
	"github.com/strongo/dalgo2sql/ramsqldb"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	db := ramsqldb.OpenTestDb(t)
	defer db.Close()
	database := NewDatabase(db, Options{})
	if database == nil {
		t.Fatal("NewDatabase(db) returned nil")
	}
}
