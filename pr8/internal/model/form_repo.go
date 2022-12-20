package model

import (
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"context"
)

type FormModel struct {
	db database.DBI
}

func NewFormModel(db database.DBI) *FormModel {
	return &FormModel{
		db: db,
	}
}

func (r *FormModel) WithTX(tx database.DBI) *FormModel {
	return NewFormModel(tx)
}

func (r *FormModel) RunInTransaction(ctx context.Context, fn func(tx *database.TX) error) error {
	return r.db.RunInTransaction(ctx, fn)
}

func (r *FormModel) AddForm(ctx context.Context, form *entity.Form) error {
	return r.RunInTransaction(ctx, func(tx *database.TX) error {
		_, err := tx.ModelContext(ctx, form).
			Value("message", "?message").
			Value("user_id", "?user_id").
			Returning("*").
			Insert()
		return err
	})
}

func (r *FormModel) GetForms(ctx context.Context) ([]*entity.Form, error) {
	result := make([]*entity.Form, 0)
	err := r.RunInTransaction(ctx, func(tx *database.TX) error {
		return tx.ModelContext(ctx, &result).Select()
	})
	return result, err
}

func (r *FormModel) GetForm(ctx context.Context, id entity.PK) (*entity.Form, error) {
	model := entity.Form{
		Id: id,
	}

	if err := r.db.ModelContext(ctx, &model).WherePK().Relation("User").Select(); err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *FormModel) DeleteForm(ctx context.Context, id entity.PK) error {
	model := entity.Form{
		Id: id,
	}
	_, err := r.db.ModelContext(ctx, &model).WherePK().Delete()
	return err
}

func (r *FormModel) UpdateForm(ctx context.Context, form *entity.Form) error {
	return r.db.RunInTransaction(ctx, func(tx *database.TX) error {
		var f entity.Form
		if err := tx.ModelContext(ctx, &f).Where("id = ?", form.Id).Select(); err != nil {
			return err
		}
		_, err := tx.ModelContext(ctx, &f).Where("id = ?", f.Id).Set("message = ?", form.Message).Update()
		return err
	})
}
