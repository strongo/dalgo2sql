package dalgo2sql

import (
	"context"
	"github.com/strongo/dalgo"
)

func (dtb database) Upsert(ctx context.Context, record dalgo.Record) error {
	panic("implement me")
}

func (t transaction) Upsert(ctx context.Context, record dalgo.Record) error {
	panic("implement me")
}
