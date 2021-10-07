package dalgo2sql

import (
	"context"
	"github.com/strongo/dalgo/dal"
)

func (dtb database) Upsert(ctx context.Context, record dal.Record) error {
	panic("implement me")
}

func (t transaction) Upsert(ctx context.Context, record dal.Record) error {
	panic("implement me")
}
