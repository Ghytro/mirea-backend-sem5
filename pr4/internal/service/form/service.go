package form

import (
	"backendmirea/pr3/internal/entity"
	"context"
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
	return s.repo.AddForm(ctx, form)
}

func (s *Service) GetForms(ctx context.Context) ([]*entity.Form, error) {
	return s.repo.GetForms(ctx)
}
