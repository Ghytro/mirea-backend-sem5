package review

import (
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/repository"
	"context"
	"errors"
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

func (s *Service) DeleteReview(ctx context.Context, who entity.PK, id entity.PK) error {
	return s.repo.RunInTransaction(ctx, func(tx *database.TX) error {
		repo := s.repo.WithTX(tx)
		review, err := repo.GetReview(ctx, id)
		if err != nil {
			return err
		}
		if review.UserId != who {
			return errors.New("you can delete only your reviews")
		}
		return repo.DeleteReview(ctx, id)
	})
}

func (s *Service) UpdateReview(ctx context.Context, whoUpdates entity.PK, review *entity.Review) error {
	return s.repo.RunInTransaction(ctx, func(tx *database.TX) error {
		repo := s.repo.WithTX(tx)
		r, err := repo.GetReview(ctx, review.Id)
		if err != nil {
			return err
		}
		if r.UserId != whoUpdates {
			return errors.New("you can only edit your reviews")
		}
		return repo.UpdateReview(ctx, review)
	})
}
