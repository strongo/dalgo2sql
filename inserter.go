package dalgo2sql

import (
	"context"
	"github.com/strongo/dalgo"
)

func (dtb database) Insert(c context.Context, record dalgo.Record, opts ...dalgo.InsertOption) error {
	return nil
}

func (t transaction) Insert(c context.Context, record dalgo.Record, opts ...dalgo.InsertOption) error {
	return nil
}
