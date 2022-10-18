package repository

import (
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"context"

	"github.com/go-pg/pg/v10"
)

type FormRepository struct {
	db database.DBI
}

func NewFormRepository(db database.DBI) *FormRepository {
	return &FormRepository{
		db: db,
	}
}

func (r *FormRepository) WithTX(tx database.DBI) *FormRepository {
	return NewFormRepository(tx)
}

func (r *FormRepository) RunInTransaction(ctx context.Context, fn func(tx *pg.Tx) error) error {
	return r.db.RunInTransaction(ctx, fn)
}

func (r *FormRepository) AddForm(ctx context.Context, form *entity.Form) error {
	return r.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := tx.ModelContext(ctx, form).Insert()
		return err
	})
}

func (r *FormRepository) GetForms(ctx context.Context) ([]*entity.Form, error) {
	result := make([]*entity.Form, 0)
	err := r.RunInTransaction(ctx, func(tx *pg.Tx) error {
		return tx.ModelContext(ctx, &result).Select()
	})
	return result, err
}
