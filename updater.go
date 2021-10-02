package dalgo2sql

import (
	"context"
	"github.com/strongo/dalgo"
)

func (dtb database) Update(ctx context.Context, key *dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	panic("implement me")
}

func (t transaction) Update(ctx context.Context, key *dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	panic("implement me")
}

func (dtb database) UpdateMulti(c context.Context, keys []*dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	panic("implement me")
}

func (t transaction) UpdateMulti(c context.Context, keys []*dalgo.Key, updates []dalgo.Update, preconditions ...dalgo.Precondition) error {
	panic("implement me")
}
