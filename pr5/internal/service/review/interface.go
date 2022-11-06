package review

import (
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/repository"
	"context"
)

type Repository interface {
	Reader
	Writer

	WithTX(database.DBI) *repository.ReviewRepository
	RunInTransaction(context.Context, func(*database.TX) error) error
}

type Writer interface {
	AddReview(context.Context, *entity.Review) error
	UpdateReview(context.Context, *entity.Review) error
	DeleteReview(context.Context, entity.PK) error
}

type Reader interface {
	GetReview(context.Context, entity.PK) (*entity.Review, error)
	GetReviews(ctx context.Context, filter *repository.ReviewFilter, order *repository.ReviewOrder, pageNumber *int, pageSize *int) ([]*entity.Review, error)
}

type UseCaseReview interface {
	GetReviews(ctx context.Context, filter *repository.ReviewFilter, order *repository.ReviewOrder, pageNumber *int, pageSize *int) ([]*entity.Review, error)
	AddReview(context.Context, *entity.Review) error
	UpdateReview(ctx context.Context, whoUpdates entity.PK, review *entity.Review) error
	DeleteReview(context.Context, entity.PK, entity.PK) error
}
