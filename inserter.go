package dalgo2sql

import (
	"context"
	"github.com/strongo/dalgo"
)

func (dtb database) Insert(ctx context.Context, record dalgo.Record, opts ...dalgo.InsertOption) error {
	return insertSingle(ctx, dtb.options, record, dtb.db.Exec)
}

func (t transaction) Insert(ctx context.Context, record dalgo.Record, opts ...dalgo.InsertOption) error {
	return insertSingle(ctx, t.options, record, t.tx.Exec)
}

func insertSingle(_ context.Context, options Options, record dalgo.Record, exec statementExecutor, opts ...dalgo.InsertOption) error {
	query := buildSingleRecordQuery(insert, options, record)
	if _, err := exec(query.text, query.args...); err != nil {
		return err
	}
	return nil
}
