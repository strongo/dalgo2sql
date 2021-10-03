package dalgo2sql

import (
	"context"
	"fmt"
	"github.com/strongo/dalgo"
)

func (dtb database) Set(ctx context.Context, record dalgo.Record) error {
	return setSingle(ctx, record, dtb.db.Query, dtb.db.Exec)
}

func (t transaction) Set(ctx context.Context, record dalgo.Record) error {
	return setSingle(ctx, record, t.tx.Query, t.tx.Exec)
}

func (dtb database) SetMulti(ctx context.Context, records []dalgo.Record) error {
	return setMulti(ctx, records, dtb.db.Query, dtb.db.Exec)
}

func (t transaction) SetMulti(ctx context.Context, records []dalgo.Record) error {
	return setMulti(ctx, records, t.tx.Query, t.tx.Exec)
}

func setSingle(_ context.Context, record dalgo.Record, execQuery queryExecutor, exec statementExecutor) error {
	exists, err := existsSingle(record.Key(), execQuery)
	if err != nil {
		return err
	}
	var qry query
	if exists {
		qry = buildSingleRecordQuery(update, record)
	} else {
		qry = buildSingleRecordQuery(insert, record)
	}
	if _, err := exec(qry.text, qry.args...); err != nil {
		return err
	}
	return nil
}

func setMulti(ctx context.Context, records []dalgo.Record, execQuery queryExecutor, execStatement statementExecutor) error {
	for _, record := range records {
		if err := setSingle(ctx, record, execQuery, execStatement); err != nil {
			return err
		}
	}
	return nil
}

func existsSingle(key *dalgo.Key, execQuery queryExecutor) (bool, error) {
	table := key.Kind()
	queryText := fmt.Sprintf("SELECT ID FROM %v WHERE ID = ?", table)
	rows, err := execQuery(queryText, key.ID)
	return err == nil && rows.Next(), err
}
