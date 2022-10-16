package review

import (
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/repository"
	"context"

	"github.com/go-pg/pg/v10"
)

type Repository interface {
	Reader
	Writer

	WithTX(database.DBI) *repository.ReviewRepository
	RunInTransaction(context.Context, func(*pg.Tx) error) error
}

type Reader interface {
	AddReview(context.Context, *entity.Review) error
}

type Writer interface {
	GetReviews(ctx context.Context, filter *repository.ReviewFilter, order *repository.ReviewOrder, pageNumber *int, pageSize *int) ([]*entity.Review, error)
}

type UseCaseReview interface {
	GetReviews(ctx context.Context, filter *repository.ReviewFilter, order *repository.ReviewOrder, pageNumber *int, pageSize *int) ([]*entity.Review, error)
	AddReview(context.Context, *entity.Review) error
}
