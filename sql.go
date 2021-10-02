package dalgo2sql

import (
	"fmt"
	"github.com/strongo/dalgo"
	"reflect"
	"strings"
)

type operation = int

const (
	insert operation = iota
	update
)

type query struct {
	text string
	args []interface{}
}

func buildSingleRecordQuery(o operation, record dalgo.Record) (query query) {
	key := record.Key()
	switch o {
	case insert:
		query.text = fmt.Sprintf("INSERT INTO %v ", key.Kind())
	case update:
		query.text = fmt.Sprintf("UPDATE %v ", key.Kind())
	}
	var cols []string
	var q []string
	data := record.Data()
	val := reflect.ValueOf(data)
	if kind := val.Kind(); kind == reflect.Interface || kind == reflect.Ptr {
		val = val.Elem()
	}
	valType := val.Type()

	if key.ID != nil && o == insert {
		query.args = append(query.args, key.ID)
		cols = append(cols, "id")
		q = append(q, "?")
	}

	for i := 0; i < val.NumField(); i++ {
		cols = append(cols, valType.Field(i).Name)
		query.args = append(query.args, val.Field(i).Interface())
		switch o {
		case insert:
			q = append(q, "?")
		case update:
			q = append(q, valType.Field(i).Name+" = ?")
		}
	}

	switch o {
	case insert:
		query.text = fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)",
			key.Kind(),
			strings.Join(cols, ", "),
			strings.Join(q, ", "),
		)
	case update:
		query.text = fmt.Sprintf("UPDATE %v SET\n", key.Kind()) +
			strings.Join(q, ",\n") +
			"WHERE id = ?"
		query.args = append(query.args, key.ID)
	}
	return query
}
