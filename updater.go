package dalgo2sql

import (
	"context"
	"fmt"
	"github.com/strongo/dalgo"
)

func (dtb database) Update(ctx context.Context, key *dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	return updateSingle(ctx, dtb.db.Exec, key, updates, preconditions...)
}

func (t transaction) Update(ctx context.Context, key *dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	return updateSingle(ctx, t.tx.Exec, key, updates, preconditions...)
}

func (dtb database) UpdateMulti(ctx context.Context, keys []*dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	return updateMulti(ctx, dtb.db.Exec, keys, updates, preconditions...)
}

func (t transaction) UpdateMulti(ctx context.Context, keys []*dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	return updateMulti(ctx, t.tx.Exec, keys, updates, preconditions...)
}

func updateSingle(_ context.Context, execStatement statementExecutor, key *dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	qry := query{
		text: fmt.Sprintf("UPDATE %v SET", key.Kind()),
	}
	for _, update := range updates {
		qry.text += fmt.Sprintf("\n\t%v = ?", update.Field)
		qry.args = append(qry.args, update.Value)
	}
	if _, err := execStatement(qry.text, qry.args...); err != nil {
		return err
	}
	return nil
}

func updateMulti(ctx context.Context, execStatement statementExecutor, keys []*dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	for _, key := range keys {
		if err := updateSingle(ctx, execStatement, key, updates, preconditions...); err != nil {
			return err
		}
	}
	return nil
}
