package repository

import (
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/utils"
	"context"
	"errors"
	"fmt"

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
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("некорректное значение рейтинга")
	}
	if review.Name == "" {
		review.Name = "<Аноним>"
	}
	return r.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := tx.ModelContext(ctx, review).
			Value("name", "?", review.Name).
			Value("rating", "?", review.Rating).
			Value("message", "?", review.Message).
			Insert()
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

func addedWhereInValues(q *orm.Query, columnName string, values interface{}) *orm.Query {
	return q.Where(fmt.Sprintf("%s IN (?)", columnName), pg.In(values))
}

func addColumnFilters[T any](q *orm.Query, columnName string, exact *T, multiple []T, r *utils.Range[T]) *orm.Query {
	if exact != nil {
		return q.Where(fmt.Sprintf("%s = ?", columnName), *exact)
	}
	if multiple != nil {
		return addedWhereInValues(q, columnName, multiple)
	}
	if r != nil {
		return addedWhereRange(q, columnName, *r)
	}
	return q
}

func (r *ReviewRepository) GetReviews(ctx context.Context, filter *ReviewFilter, order *ReviewOrder, pageNumber, pageSize *int) ([]*entity.Review, error) {
	var result []*entity.Review
	err := r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		q := tx.ModelContext(ctx, &result)
		if filter != nil {
			q = addColumnFilters(q, "id", filter.Id, filter.Ids, filter.IdsRange)
			q = addColumnFilters(q, "posted_at", filter.Time, filter.Times, filter.TimeRange)
			q = addColumnFilters(q, "rating", filter.Rating, filter.Ratings, filter.RatingsRange)
			q = addColumnFilters(q, "name", filter.Name, nil, nil)
		}

		if order != nil {
			strOrder := "DESC"
			if order.IsAscending {
				strOrder = "ASC"
			}
			q = q.Order(fmt.Sprintf("%s %s", order.FieldName, strOrder))
		}
		if pageNumber != nil && pageSize != nil {
			q = q.Offset(*pageSize * *pageNumber).Limit(*pageNumber)
		}
		return q.Select()
	})
	if result == nil {
		result = make([]*entity.Review, 0)
	}
	return result, err
}
