package dalgo2sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/georgysavva/scany/sqlscan"
	"github.com/strongo/dalgo"
	"reflect"
	"strings"
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
	if err = rowIntoRecord(rows, record); err != nil {
		return err
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
	columns := []string{"StringProp", "IntegerProp"}
	queryText := fmt.Sprintf("SELECT %v FROM %v WHERE ID IN (",
		strings.Join(columns, ", "),
		records[0].Key().Kind(),
	)
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

	for rows.Next() {
		var id string
		cells := []interface{}{&id}
		if err = rows.Scan(cells...); err != nil {
			return err
		}
		for i, record := range records {
			if record.Key().ID == id {
				records = append(records[:i], records[i+1:]...)
				if err = rowIntoRecord(rows, record); err != nil {
					return err
				}
				break
			}
		}
	}
	if err = rows.Err(); err == sql.ErrNoRows {
		err = nil
	} else if err != nil {
		return err
	}
	for _, record := range records {
		record.SetError(dalgo.NewErrNotFoundByKey(record.Key(), nil))
	}
	return err
}

func rowIntoRecord(rows *sql.Rows, record dalgo.Record) error {
	data := record.Data()
	if data != nil {
		if err := scanIntoData(rows, data, record); err != nil {
			return err
		}
	}
	return delayedScanWithDataTo(rows, record)
}

func delayedScanWithDataTo(rows *sql.Rows, record dalgo.Record) error {
	row, err := scanIntoMap(rows)
	if err != nil {
		record.SetError(err)
		return err
	}
	record.SetDataTo(func(target interface{}) error {
		t := reflect.ValueOf(target)
		val := t.Elem()
		valType := val.Type()
		for i := 0; i < val.NumField(); i++ {
			if val.Field(i).CanSet() {
				fieldName := valType.Field(i).Name
				if v, hasValue := row[fieldName]; hasValue {
					val.Set(reflect.ValueOf(v))
				}
			}
		}
		return nil
	})
	return nil
}

func scanIntoData(rows *sql.Rows, data interface{}, record dalgo.Record) error {
	return sqlscan.ScanRow(data, rows)
}

func scanIntoMap(rows *sql.Rows) (row map[string]interface{}, err error) {

	cols, err := rows.Columns()

	// Create a slice of interface{}'s to represent each cell,
	// and a second slice to contain pointers to each item in the cells slice.
	cells := make([]interface{}, len(cols))
	cellPointers := make([]interface{}, len(cols))
	for i := range cells {
		cellPointers[i] = &cells[i]
	}

	// Scan the row into the cell pointers...
	if err := rows.Scan(cellPointers...); err != nil {
		return nil, err
	}

	// Create our map, and retrieve the value for each column from the pointers slice,
	// storing it in the map with the name of the column as the key.
	m := make(map[string]interface{}, len(cols))
	for i, colName := range cols {
		val := cellPointers[i].(*interface{})
		m[colName] = *val
	}
	return m, nil
}
