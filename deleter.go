package dalgo2sql

import (
	"context"
	"github.com/strongo/dalgo"
)

func (dtb database) Delete(ctx context.Context, key *dalgo.Key) error {
	panic("implement me")
}

func (t transaction) Delete(ctx context.Context, key *dalgo.Key) error {
	panic("implement me")
}

func (dtb database) DeleteMulti(ctx context.Context, keys []*dalgo.Key) error {
	panic("implement me")
}

func (t transaction) DeleteMulti(ctx context.Context, keys []*dalgo.Key) error {
	panic("implement me")
}
