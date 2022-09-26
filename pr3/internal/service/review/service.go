package review

import (
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/repository"
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

func (s *Service) GetReviews(ctx context.Context, filter *repository.ReviewFilter) ([]*entity.Review, error) {
	return s.repo.GetReviews(ctx, filter)
}

func (s *Service) AddReview(ctx context.Context, review *entity.Review) error {
	return s.repo.AddReview(ctx, review)
}
