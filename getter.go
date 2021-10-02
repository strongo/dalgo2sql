package dalgo2sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/strongo/dalgo"
	"reflect"
)

type queryExecutor = func(query string, args ...interface{}) (*sql.Rows, error)

func (dtb database) Get(ctx context.Context, record dalgo.Record) error {
	return getSingle(ctx, record, dtb.db.Query)
}

func (t transaction) Get(ctx context.Context, record dalgo.Record) error {
	return getSingle(ctx, record, t.tx.Query)
}

func (dtb database) GetMulti(ctx context.Context, records []dalgo.Record) error {
	return getMulti(ctx, records, dtb.db.Query)
}

func (t transaction) GetMulti(ctx context.Context, records []dalgo.Record) error {
	return getMulti(ctx, records, t.tx.Query)
}

func getSingle(_ context.Context, record dalgo.Record, exec queryExecutor) error {
	queryText := fmt.Sprintf("SELECT * FROM %v", record.Key().Kind())
	rows, err := exec(queryText)
	if err != nil {
		record.SetError(err)
		return err
	}
	if !rows.Next() {
		record.SetError(dalgo.ErrRecordNotFound)
		return dalgo.NewErrNotFoundByKey(record.Key(), dalgo.ErrNoMoreRecords)
	}
	if data := record.Data(); data == nil {
		record.SetDataTo(func(target interface{}) error {
			return rows.Scan(target)
		})
	} else {
		if err := rows.Scan(data); err != nil {
			return err
		}
	}
	if rows.Next() {
		return errors.New("expected to get single row but got multiple")
	}
	return nil
}

func getMulti(ctx context.Context, records []dalgo.Record, exec queryExecutor) error {
	return getMultiFromSingleTable(ctx, records, exec)
}

func getMultiFromSingleTable(_ context.Context, records []dalgo.Record, exec queryExecutor) error {
	queryText := fmt.Sprintf("SELECT * FROM %v WHERE id IN (", records[0].Key().Kind())
	args := make([]interface{}, len(records))
	for i, record := range records {
		args[i] = record.Key().ID
		queryText += "?,"
	}
	queryText = queryText[:len(queryText)-1] + ")"
	rows, err := exec(queryText, args...)
	if err != nil {
		return err
	}
	var row struct {
		ID interface{}
	}
	for rows.Next() {
		if err = rows.Scan(&row); err != nil {
			return err
		}
		for i, record := range records {
			if record.Key().ID == row.ID {
				records = append(records[:i], records[i+1:]...)
				data := record.Data()
				if data != nil {
					if err := rows.Scan(data); err != nil {
						record.SetError(err)
						return err
					}
					break
				}
				row := make(map[string]interface{})
				if err = rows.Scan(row); err != nil {
					record.SetError(err)
					return err
				}
				r := record
				r.SetDataTo(func(target interface{}) error {
					val := reflect.ValueOf(target).Elem()
					for i := 0; i < val.NumField(); i++ {
						if val.Field(i).CanSet() {
							//panic("implement me")
						}
					}
					return nil
				})
			}
		}
	}
	for _, record := range records {
		record.SetError(dalgo.NewErrNotFoundByKey(record.Key(), dalgo.ErrRecordNotFound))
	}
	return nil
}
