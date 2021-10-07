package dalgo2sql

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/strongo/dalgo/dal"
)

// Field defines field
type Field struct {
	Name string
}

// RecordsetType defines type of a database recordset
type RecordsetType = int

const (
	// Table identifies a table in a database
	Table RecordsetType = iota
	// View identifies a view in a database
	View
	// StoredProcedure identifies a stored procedure in a database
	StoredProcedure
)

// Recordset hold recordset settings
type Recordset struct {
	Type       RecordsetType
	Name       string
	PrimaryKey []Field // Primary keys by table name
}

type database struct {
	db      *sql.DB
	options Options
}

// Options provides database sqlOptions for DALgo
type Options struct {
	Recordsets map[string]Recordset
}

func (dtb database) RunReadonlyTransaction(ctx context.Context, f dal.ROTxWorker, options ...dal.TransactionOption) error {
	return nil
}

func (dtb database) RunReadwriteTransaction(ctx context.Context, f dal.RWTxWorker, options ...dal.TransactionOption) error {
	dalgoTxOptions := dal.NewTransactionOptions(options...)
	sqlTxOptions := sql.TxOptions{}
	if dalgoTxOptions.IsReadonly() {
		sqlTxOptions.ReadOnly = true
	}
	dbTx, err := dtb.db.BeginTx(ctx, &sqlTxOptions)
	if err != nil {
		return err
	}
	if err = f(ctx, transaction{tx: dbTx, sqlOptions: dtb.options}); err != nil {
		if rollbackErr := dbTx.Rollback(); rollbackErr != nil {
			return dal.NewRollbackError(rollbackErr, err)
		}
		return err
	}
	if err := dbTx.Commit(); err != nil {
		return errors.WithMessage(err, "failed to commit transaction")
	}
	return nil
}

func (dtb database) Select(ctx context.Context, query dal.Select) (dal.Reader, error) {
	panic("implement me")
}

var _ dal.Database = (*database)(nil)

// NewDatabase creates a new instance of DALgo adapter for BungDB
func NewDatabase(db *sql.DB, options Options) dal.Database {
	if db == nil {
		panic("db is a required parameter, got nil")
	}
	return database{
		db:      db,
		options: options,
	}
}
