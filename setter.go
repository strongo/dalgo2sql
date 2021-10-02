package dalgo2sql

import (
	"context"
	"github.com/strongo/dalgo"
)

func (dtb database) Set(ctx context.Context, record dalgo.Record) error {
	panic("implement me")
}

func (t transaction) Set(ctx context.Context, record dalgo.Record) error {
	panic("implement me")
}

func (dtb database) SetMulti(ctx context.Context, records []dalgo.Record) error {
	panic("implement me")
}

func (t transaction) SetMulti(ctx context.Context, records []dalgo.Record) error {
	panic("implement me")
}
