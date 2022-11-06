package repository

import (
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"context"
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

func (r *FormRepository) RunInTransaction(ctx context.Context, fn func(tx *database.TX) error) error {
	return r.db.RunInTransaction(ctx, fn)
}

func (r *FormRepository) AddForm(ctx context.Context, form *entity.Form) error {
	return r.RunInTransaction(ctx, func(tx *database.TX) error {
		_, err := tx.ModelContext(ctx, form).
			Value("message", "?message").
			Value("user_id", "?user_id").
			Returning("*").
			Insert()
		return err
	})
}

func (r *FormRepository) GetForms(ctx context.Context) ([]*entity.Form, error) {
	result := make([]*entity.Form, 0)
	err := r.RunInTransaction(ctx, func(tx *database.TX) error {
		return tx.ModelContext(ctx, &result).Select()
	})
	return result, err
}

func (r *FormRepository) GetForm(ctx context.Context, id entity.PK) (*entity.Form, error) {
	model := entity.Form{
		Id: id,
	}

	if err := r.db.ModelContext(ctx, &model).WherePK().Relation("User").Select(); err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *FormRepository) DeleteForm(ctx context.Context, id entity.PK) error {
	model := entity.Form{
		Id: id,
	}
	_, err := r.db.ModelContext(ctx, &model).WherePK().Delete()
	return err
}

func (r *FormRepository) UpdateForm(ctx context.Context, form *entity.Form) error {
	return r.db.RunInTransaction(ctx, func(tx *database.TX) error {
		var f entity.Form
		if err := tx.ModelContext(ctx, &f).Where("id = ?", form.Id).Select(); err != nil {
			return err
		}
		_, err := tx.ModelContext(ctx, &f).Where("id = ?", f.Id).Set("message = ?", form.Message).Update()
		return err
	})
}
