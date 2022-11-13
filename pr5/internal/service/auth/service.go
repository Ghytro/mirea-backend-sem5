package auth

import (
	"context"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) StartSession(ctx context.Context, username string, password string) (string, time.Time, error) {
	return s.repo.StartSession(ctx, username, password)
}
