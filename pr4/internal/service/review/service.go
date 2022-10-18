package review

import (
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/repository"
	"context"

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

func (s *Service) GetReviews(ctx context.Context, filter *repository.ReviewFilter, order *repository.ReviewOrder, pageNumber, pageSize *int) ([]*entity.Review, error) {
	return s.repo.GetReviews(ctx, filter, order, pageNumber, pageSize)
}

func (s *Service) AddReview(ctx context.Context, review *entity.Review) error {
	return s.repo.AddReview(ctx, review)
}

func (s *Service) DeleteReview(ctx context.Context, id entity.PK) error {
	return s.repo.RunInTransaction(ctx, func(tx *pg.Tx) error {
		repo := s.repo.WithTX(tx)
		if _, err := repo.GetReview(ctx, id); err != nil {
			return err
		}
		return repo.DeleteReview(ctx, id)
	})
}
