package dalgo2sql

import (
	"context"
	"database/sql"
	"github.com/strongo/dalgo/dal"
)

type transaction struct {
	tx         *sql.Tx
	sqlOptions Options
	txOptions  dal.TransactionOptions
}

func (t transaction) Options() dal.TransactionOptions {
	return t.txOptions
}

func (t transaction) Select(ctx context.Context, query dal.Select) (dal.Reader, error) {
	panic("implement me")
}

var _ dal.Transaction = (*transaction)(nil)
