package dalgo2sql

import (
	"context"
	"github.com/strongo/dalgo"
)

func (dtb database) Insert(c context.Context, record dalgo.Record, opts ...dalgo.InsertOption) error {
	panic("implement me")
}

func (t transaction) Insert(c context.Context, record dalgo.Record, opts ...dalgo.InsertOption) error {
	panic("implement me")
}
