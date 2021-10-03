package end2end_test

import (
	end2end "github.com/strongo/dalgo-end2end-tests"
	"github.com/strongo/dalgo2sql"
	"github.com/strongo/dalgo2sql/ramsqldb"
	"testing"
)

func TestEndToEnd(t *testing.T) {
	db := ramsqldb.OpenTestDb(t)
	defer func() {
		_ = db.Close()
	}()
	options := dalgo2sql.Options{
		Recordsets: map[string]dalgo2sql.Recordset{
			"E2ETest1": {
				Name:       "E2ETest1",
				Type:       dalgo2sql.Table,
				PrimaryKey: []dalgo2sql.Field{{Name: "ID1"}},
			},
		},
	}
	database := dalgo2sql.NewDatabase(db, options)
	end2end.TestDalgoDB(t, database)
}
