package dalgo2sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/strongo/dalgo/dal"
	"strings"
)

type statementExecutor = func(query string, args ...interface{}) (sql.Result, error)

func (dtb database) Delete(ctx context.Context, key *dal.Key) error {
	return deleteSingle(ctx, dtb.options, key, dtb.db.Exec)
}

func (t transaction) Delete(ctx context.Context, key *dal.Key) error {
	return deleteSingle(ctx, t.options, key, t.tx.Exec)
}

func (dtb database) DeleteMulti(ctx context.Context, keys []*dal.Key) error {
	return deleteMulti(ctx, dtb.options, keys, dtb.db.Exec)
}

func deleteSingle(_ context.Context, options Options, key *dal.Key, exec statementExecutor) error {
	collection := key.Kind()
	query := fmt.Sprintf("DELETE FROM %v WHERE ", key.Kind())
	if rs, hasOptions := options.Recordsets[collection]; hasOptions && len(rs.PrimaryKey) == 1 {
		query += rs.PrimaryKey[0].Name + " = ?"
	} else {
		query += "ID = ?"
	}
	_, err := exec(query, key.ID)
	if err != nil {
		return err
	}
	return nil
}

func deleteMulti(ctx context.Context, options Options, keys []*dal.Key, exec statementExecutor) error {
	var prevTable string
	var tableKeys []*dal.Key
	delete := func(table string, keys []*dal.Key) error {
		if len(keys) == 0 {
			return nil
		}
		if len(keys) == 1 {
			if err := deleteSingle(ctx, options, keys[0], exec); err != nil {
				return err
			}
			return nil
		}
		for _, key := range keys {
			if err := deleteSingle(ctx, options, key, exec); err != nil {
				return err
			}
		}
		//if err := deleteMultiInSingleTable(ctx, options, keys, exec); err != nil {
		//	return err
		//}
		return nil // TODO: code above commented out as tests are failing for RAMSQL driver.
	}
	for i, key := range keys {
		kind := key.Kind()
		if kind == prevTable {
			tableKeys = append(tableKeys, key)
			continue
		}
		if prevTable != "" {
			if err := delete(prevTable, tableKeys); err != nil {
				return err
			}
		}
		prevTable = kind
		tableKeys = make([]*dal.Key, 1, len(keys)-i)
		tableKeys[0] = key
	}
	if len(tableKeys) > 0 {
		if err := delete(prevTable, tableKeys); err != nil {
			return err
		}
	}
	return nil
}
func deleteMultiInSingleTable(_ context.Context, options Options, keys []*dal.Key, exec statementExecutor) error {
	pkCol := "ID"

	collection := keys[0].Kind()
	if rs, hasOptions := options.Recordsets[collection]; hasOptions && len(rs.PrimaryKey) == 1 {
		pkCol = rs.PrimaryKey[0].Name
	}

	query := fmt.Sprintf("DELETE FROM %v WHERE %v IN (", collection, pkCol)
	args := make([]interface{}, len(keys))
	q := make([]string, len(keys))
	for i, key := range keys {
		args[i] = key.ID
		q[i] = "?"
	}
	query += strings.Join(q, ", ") + ")"
	_, err := exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (t transaction) DeleteMulti(ctx context.Context, keys []*dal.Key) error {
	return deleteMulti(ctx, t.options, keys, t.tx.Exec)
}
