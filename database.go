package dalgo2sql

import (
	"context"
	"database/sql"
	"github.com/strongo/dalgo"
)

type database struct {
	db *sql.DB
}

func (dtb database) RunInTransaction(ctx context.Context, f func(ctx context.Context, tx dalgo.Transaction) error, options ...dalgo.TransactionOption) error {
	panic("implement me")
}


func (dtb database) Select(ctx context.Context, query dalgo.Query) (dalgo.Reader, error) {
	panic("implement me")
}

var _ dalgo.Database = (*database)(nil)

// NewDatabase creates a new instance of DALgo adapter for BungDB
func NewDatabase(db *sql.DB) dalgo.Database {
	if db == nil {
		panic("db is a required parameter, got nil")
	}
	return database{
		db: db,
	}
}
