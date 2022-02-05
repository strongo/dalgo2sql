package dalgo2sql

import (
	"github.com/strongo/dalgo/dal"
	"reflect"
	"testing"
)

func Test_getSelectFields(t *testing.T) {
	type args struct {
		record    dal.Record
		includePK bool
		options   Options
	}

	tests := []struct {
		name       string
		args       args
		wantFields []string
	}{
		{
			name: "simple_fields",
			args: args{
				record: dal.NewRecordWithoutKey(struct {
					StrField string
					IntField int
				}{}),
			},
			wantFields: []string{"StrField", "IntField"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFields := getSelectFields(tt.args.record, tt.args.includePK, tt.args.options); !reflect.DeepEqual(gotFields, tt.wantFields) {
				t.Errorf("getSelectFields() = %v, want %v", gotFields, tt.wantFields)
			}
		})
	}
}
