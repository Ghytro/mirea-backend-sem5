package repository

import (
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/utils"
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type ReviewRepository struct {
	db database.DBI
}

func NewReviewRepository(db database.DBI) *ReviewRepository {
	return &ReviewRepository{
		db: db,
	}
}

func (r *ReviewRepository) WithTX(tx database.DBI) *ReviewRepository {
	return NewReviewRepository(tx)
}

func (r *ReviewRepository) RunInTransaction(ctx context.Context, fn func(tx *pg.Tx) error) error {
	return r.db.RunInTransaction(ctx, fn)
}

func (r *ReviewRepository) AddReview(ctx context.Context, review *entity.Review) error {
	return r.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := tx.ModelContext(ctx, review).Insert()
		return err
	})
}

func addedWhereRange[T any](q *orm.Query, columnName string, r utils.Range[T]) *orm.Query {
	var (
		whereCondition string
		args           []interface{}
	)
	if r.From != nil {
		whereCondition += "id >= ?"
		args = append(args, r.From)
		if r.To != nil {
			whereCondition += " AND id <= ?"
			args = append(args, r.To)
		}
	} else {
		if r.To != nil {
			whereCondition += "id <= ?"
			args = append(args, r.From)
		}
	}
	return q.Where(whereCondition, args...)
}

func (r *ReviewRepository) GetReviews(ctx context.Context, filter *ReviewFilter) ([]*entity.Review, error) {
	var result []*entity.Review
	err := r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		q := tx.ModelContext(ctx, &result)
		if filter == nil {
			return q.Select()
		}

		if filter.Id != nil {
			q = q.Where("id = ?", *filter.Id)
		} else if filter.Ids != nil {
			q = addedWhereRange(q, "id", *filter.Ids)
		}

		if filter.TimeRange != nil {
			q = addedWhereRange(q, "posted_at", *filter.TimeRange)
		}

		if filter.Rating != nil {
			q = q.Where("rating = ?", *filter.Rating)
		} else if filter.Ratings != nil {
			q = addedWhereRange(q, "rating", *filter.Ratings)
		}

		return q.Select()
	})
	return result, err
}
