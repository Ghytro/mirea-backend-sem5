package database

import (
	"context"

	"github.com/go-pg/pg/v10"
)

type DB struct {
	*pg.DB
}

func (d *DB) RunInTransaction(ctx context.Context, f func(tx *TX) error) error {
	return d.DB.RunInTransaction(ctx, func(tx *pg.Tx) error {
		return f(&TX{Tx: tx})
	})
}

type TX struct {
	*pg.Tx
}

func (t *TX) RunInTransaction(ctx context.Context, f func(tx *TX) error) error {
	return f(t)
}
