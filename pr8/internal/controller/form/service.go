package form

import (
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"context"
	"errors"
)

type Controller struct {
	repo Model
}

func NewController(repo Model) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (s *Controller) AddForm(ctx context.Context, form *entity.Form) error {
	if form == nil {
		return errors.New("nil form passed to service")
	}
	if form.UserId == 0 {
		return errors.New("unknown sender of form")
	}
	return s.repo.AddForm(ctx, form)
}

func (s *Controller) GetForms(ctx context.Context) ([]*entity.Form, error) {
	return s.repo.GetForms(ctx)
}

func (s *Controller) DeleteForm(ctx context.Context, whoDeletes entity.PK, id entity.PK) error {
	return s.repo.RunInTransaction(ctx, func(tx *database.TX) error {
		repo := s.repo.WithTX(tx)
		f, err := repo.GetForm(ctx, id)
		if err != nil {
			return err
		}
		if f.UserId != whoDeletes {
			return errors.New("you are only able to delete your own forms")
		}
		return repo.DeleteForm(ctx, id)
	})
}

func (s *Controller) UpdateForm(ctx context.Context, whoUpdates entity.PK, form *entity.Form) error {
	return s.repo.RunInTransaction(ctx, func(tx *database.TX) error {
		repo := s.repo.WithTX(tx)
		f, err := repo.GetForm(ctx, form.Id)
		if err != nil {
			return err
		}
		if f.UserId != whoUpdates {
			return errors.New("you can update only your forms")
		}
		return repo.UpdateForm(ctx, form)
	})
}
