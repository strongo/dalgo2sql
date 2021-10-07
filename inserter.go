package dalgo2sql

import (
	"context"
	"github.com/strongo/dalgo/dal"
)

func (dtb database) Insert(ctx context.Context, record dal.Record, opts ...dal.InsertOption) error {
	return insertSingle(ctx, dtb.options, record, dtb.db.Exec)
}

func (t transaction) Insert(ctx context.Context, record dal.Record, opts ...dal.InsertOption) error {
	return insertSingle(ctx, t.options, record, t.tx.Exec)
}

func insertSingle(_ context.Context, options Options, record dal.Record, exec statementExecutor, opts ...dal.InsertOption) error {
	query := buildSingleRecordQuery(insert, options, record)
	if _, err := exec(query.text, query.args...); err != nil {
		return err
	}
	return nil
}
