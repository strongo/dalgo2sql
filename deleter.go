package dalgo2sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/strongo/dalgo"
)

type executor = func(query string, args ...interface{}) (sql.Result, error)

func (dtb database) Delete(ctx context.Context, key *dalgo.Key) error {
	return deleteSingle(ctx, key, dtb.db.Exec)
}

func (t transaction) Delete(ctx context.Context, key *dalgo.Key) error {
	return deleteSingle(ctx, key, t.tx.Exec)
}

func (dtb database) DeleteMulti(ctx context.Context, keys []*dalgo.Key) error {
	return deleteMulti(ctx, keys, dtb.db.Exec)
}

func deleteSingle(_ context.Context, key *dalgo.Key, exec executor) error {
	query := fmt.Sprintf(`DELETE FROM %v WHERE id = ?`, key.Kind())
	_, err := exec(query, key.ID)
	if err != nil {
		return err
	}
	return nil
}

func deleteMulti(ctx context.Context, keys []*dalgo.Key, exec executor) error {
	var prevTable string
	var tableKeys []*dalgo.Key
	delete := func(table string, keys []*dalgo.Key) error {
		if len(keys) == 0 {
			return nil
		}
		if len(keys) == 1 {
			if err := deleteSingle(ctx, keys[0], exec); err != nil {
				return err
			}
			return nil
		}
		if err := deleteMultiInSingleTable(ctx, keys, exec); err != nil {
			return err
		}
		return nil
	}
	for _, key := range keys {
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
		tableKeys = make([]*dalgo.Key, 0)
	}
	if len(tableKeys) > 0 {
		if err := delete(prevTable, tableKeys); err != nil {
			return err
		}
	}
	return nil
}
func deleteMultiInSingleTable(_ context.Context, keys []*dalgo.Key, exec executor) error {
	query := fmt.Sprintf(`DELETE FROM %v WHERE id IN (`, keys[0].Kind())
	args := make([]interface{}, len(keys))
	query += ""
	for i, key := range keys {
		args[i] = key.ID
		query += "?"
	}
	query = query[:len(query)-1] + ")"
	_, err := exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (t transaction) DeleteMulti(ctx context.Context, keys []*dalgo.Key) error {
	return deleteMulti(ctx, keys, t.tx.Exec)
}
