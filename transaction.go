package dalgo2sql

import (
	"context"
	"database/sql"
	"github.com/strongo/dalgo/dal"
)

type transaction struct {
	tx             *sql.Tx
	options        Options
	isolationLevel dal.TxIsolationLevel
}

func (t transaction) IsolationLevel() dal.TxIsolationLevel {
	return t.isolationLevel
}

func (t transaction) Select(ctx context.Context, query dal.Select) (dal.Reader, error) {
	panic("implement me")
}

var _ dal.Transaction = (*transaction)(nil)
