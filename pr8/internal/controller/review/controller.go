package review

import (
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/model"
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

func (s *Controller) GetReviews(ctx context.Context, filter *model.ReviewFilter, order *model.ReviewOrder, pageNumber, pageSize *int) ([]*entity.Review, error) {
	return s.repo.GetReviews(ctx, filter, order, pageNumber, pageSize)
}

func (s *Controller) AddReview(ctx context.Context, review *entity.Review) error {
	return s.repo.AddReview(ctx, review)
}

func (s *Controller) DeleteReview(ctx context.Context, who entity.PK, id entity.PK) error {
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

func (s *Controller) UpdateReview(ctx context.Context, whoUpdates entity.PK, review *entity.Review) error {
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
