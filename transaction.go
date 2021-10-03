package dalgo2sql

import (
	"context"
	"database/sql"
	"github.com/strongo/dalgo"
)

type transaction struct {
	tx      *sql.Tx
	options Options
}

func (t transaction) Select(ctx context.Context, query dalgo.Query) (dalgo.Reader, error) {
	panic("implement me")
}

var _ dalgo.Transaction = (*transaction)(nil)
