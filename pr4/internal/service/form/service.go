package form

import (
	"backendmirea/pr3/internal/entity"
	"context"
	"errors"
	"net/mail"
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
