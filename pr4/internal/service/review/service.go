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

func (s *Service) GetReviews(ctx context.Context, filter *repository.ReviewFilter, order *repository.ReviewOrder, pageNumber, pageSize *int) ([]*entity.Review, error) {
	return s.repo.GetReviews(ctx, filter, order, pageNumber, pageSize)
}

func (s *Service) AddReview(ctx context.Context, review *entity.Review) error {
	return s.repo.AddReview(ctx, review)
}
