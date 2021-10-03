package dalgo2sql

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/strongo/dalgo"
)

type Field struct {
	Name string
}

type database struct {
	db *sql.DB
	pk map[string][]Field // Primary keys by table name
}

func (dtb database) RunInTransaction(ctx context.Context, f func(ctx context.Context, tx dalgo.Transaction) error, options ...dalgo.TransactionOption) error {
	dalgoTxOptions := dalgo.NewTransactionOptions(options...)
	sqlTxOptions := sql.TxOptions{}
	if dalgoTxOptions.IsReadonly() {
		sqlTxOptions.ReadOnly = true
	}
	dbTx, err := dtb.db.BeginTx(ctx, &sqlTxOptions)
	if err != nil {
		return err
	}
	if err = f(ctx, transaction{tx: dbTx}); err != nil {
		if rollbackErr := dbTx.Rollback(); rollbackErr != nil {
			return dalgo.NewRollbackError(rollbackErr, err)
		}
		return err
	}
	if err := dbTx.Commit(); err != nil {
		return errors.WithMessage(err, "failed to commit transaction")
	}
	return nil
}

func (dtb database) Select(ctx context.Context, query dalgo.Query) (dalgo.Reader, error) {
	panic("implement me")
}

var _ dalgo.Database = (*database)(nil)

// NewDatabase creates a new instance of DALgo adapter for BungDB
func NewDatabase(db *sql.DB, pk map[string][]Field) dalgo.Database {
	if db == nil {
		panic("db is a required parameter, got nil")
	}
	return database{
		db: db,
		pk: pk,
	}
}
