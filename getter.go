package dalgo2sql

import (
	"context"
	"github.com/strongo/dalgo"
)

func (dtb database) Get(ctx context.Context, record dalgo.Record) error {
	panic("implement me")
}

func (t transaction) Get(ctx context.Context, record dalgo.Record) error {
	panic("implement me")
}

func (dtb database) GetMulti(ctx context.Context, records []dalgo.Record) error {
	panic("implement me")
}

func (t transaction) GetMulti(ctx context.Context, records []dalgo.Record) error {
	panic("implement me")
}
