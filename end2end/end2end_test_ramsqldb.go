package end2end

import (
	"github.com/strongo/dalgo/end2end"
	"github.com/strongo/dalgo2sql"
	"github.com/strongo/dalgo2sql/end2end/ramsqldb"
	"testing"
)

func testEndToEndRAMSQLDB(t *testing.T, options dalgo2sql.Options) {
	db := ramsqldb.OpenTestDb(t)
	defer func() {
		_ = db.Close()
	}()
	database := dalgo2sql.NewDatabase(db, options)
	end2end.TestDalgoDB(t, database)
}
