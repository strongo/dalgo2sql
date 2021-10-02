package end2end_test

import (
	end2end "github.com/strongo/dalgo-end2end-tests"
	"github.com/strongo/dalgo2sql"
	"github.com/strongo/dalgo2sql/ramsqldb"
	"testing"
)

func TestEndToEnd(t *testing.T) {
	db := ramsqldb.OpenTestDb(t)
	defer db.Close()
	database := dalgo2sql.NewDatabase(db)
	end2end.TestDalgoDB(t, database)
}
