package auth

import (
	"context"
	"time"
)

type Controller struct {
	repo Model
}

func NewController(repo Model) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (s *Controller) StartSession(ctx context.Context, username string, password string) (string, time.Time, error) {
	return s.repo.StartSession(ctx, username, password)
}
