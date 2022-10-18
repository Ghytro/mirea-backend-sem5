package form

import (
	"backendmirea/pr3/internal/entity"
	"context"
	"errors"
	"net/mail"

	"github.com/go-pg/pg/v10"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddForm(ctx context.Context, form *entity.Form) error {
	if form == nil {
		return errors.New("nil form passed to service")
	}
	if form.Name == "" {
		return errors.New("empty name in form")
	}
	if form.Email == "" {
		return errors.New("empty email in form")
	}
	if form.Message == "" {
		return errors.New("empty message in form")
	}
	if _, err := mail.ParseAddress(form.Email); err != nil {
		return errors.New("incorrect format of email")
	}
	return s.repo.AddForm(ctx, form)
}

func (s *Service) GetForms(ctx context.Context) ([]*entity.Form, error) {
	return s.repo.GetForms(ctx)
}

func (s *Service) DeleteForm(ctx context.Context, id entity.PK) error {
	return s.repo.RunInTransaction(ctx, func(tx *pg.Tx) error {
		repo := s.repo.WithTX(tx)
		if _, err := repo.GetForm(ctx, id); err != nil {
			return err
		}
		return repo.DeleteForm(ctx, id)
	})
}
