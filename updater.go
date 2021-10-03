package dalgo2sql

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/strongo/dalgo"
)

func (dtb database) Update(ctx context.Context, key *dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	return updateSingle(ctx, dtb.options, dtb.db.Exec, key, updates, preconditions...)
}

func (t transaction) Update(ctx context.Context, key *dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	return updateSingle(ctx, t.options, t.tx.Exec, key, updates, preconditions...)
}

func (dtb database) UpdateMulti(ctx context.Context, keys []*dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	return updateMulti(ctx, dtb.options, dtb.db.Exec, keys, updates, preconditions...)
}

func (t transaction) UpdateMulti(ctx context.Context, keys []*dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	return updateMulti(ctx, t.options, t.tx.Exec, keys, updates, preconditions...)
}

func updateSingle(_ context.Context, options Options, execStatement statementExecutor, key *dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	qry := query{
		text: fmt.Sprintf("UPDATE %v SET", key.Kind()),
	}
	for _, update := range updates {
		qry.text += fmt.Sprintf("\n\t%v = ?", update.Field)
		qry.args = append(qry.args, update.Value)
	}
	result, err := execStatement(qry.text, qry.args...)
	if err != nil {
		return errors.WithMessage(err, "failed to update a single record")
	}
	collection := key.Kind()
	if rs, hasOptions := options.Recordsets[collection]; hasOptions && len(rs.PrimaryKey) == 1 {
		qry.text += fmt.Sprintf("\n\tWHERE %v = ?", rs.PrimaryKey[0].Name)
	} else {
		qry.text += "\n\tWHERE ID = ?"
	}
	qry.args = append(qry.args, key.ID)
	if count, err := result.RowsAffected(); err == nil && count > 1 {
		return fmt.Errorf("expected to update a single row, number of affected rows: %v", count)
	}
	return nil
}

func updateMulti(ctx context.Context, options Options, execStatement statementExecutor, keys []*dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	for i, key := range keys {
		if err := updateSingle(ctx, options, execStatement, key, updates, preconditions...); err != nil {
			return errors.WithMessagef(err, "failed to update record #%v of %v", i+1, len(keys))
		}
	}
	return nil
}
