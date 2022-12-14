package model

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

type ReviewModel struct {
	db database.DBI
}

func NewReviewModel(db database.DBI) *ReviewModel {
	return &ReviewModel{
		db: db,
	}
}

func (r *ReviewModel) WithTX(tx database.DBI) *ReviewModel {
	return NewReviewModel(tx)
}

func (r *ReviewModel) RunInTransaction(ctx context.Context, fn func(tx *database.TX) error) error {
	return r.db.RunInTransaction(ctx, fn)
}

func (r *ReviewModel) AddReview(ctx context.Context, review *entity.Review) error {
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("некорректное значение рейтинга")
	}
	return r.RunInTransaction(ctx, func(tx *database.TX) error {
		_, err := tx.ModelContext(ctx, review).
			Value("rating", "?", review.Rating).
			Value("message", "?", review.Message).
			Value("user_id", "?", review.UserId).
			Returning("*").
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
		whereCondition += columnName + " >= ?"
		args = append(args, r.From)
		if r.To != nil {
			whereCondition += fmt.Sprintf(" AND %s <= ?", columnName)
			args = append(args, r.To)
		}
	} else {
		if r.To != nil {
			whereCondition += columnName + " <= ?"
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

func (r *ReviewModel) GetReview(ctx context.Context, id entity.PK) (*entity.Review, error) {
	var model entity.Review
	if err := r.db.ModelContext(ctx, &model).Where("id = ?", id).Select(); err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *ReviewModel) GetReviews(ctx context.Context, filter *ReviewFilter, order *ReviewOrder, pageNumber, pageSize *int) ([]*entity.Review, error) {
	var result []*entity.Review
	err := r.db.RunInTransaction(ctx, func(tx *database.TX) error {
		q := tx.ModelContext(ctx, &result)
		if filter != nil {
			q = addColumnFilters(q, `"review"."id"`, filter.Id, filter.Ids, filter.IdsRange)
			q = addColumnFilters(q, `"review"."posted_at"`, filter.Time, filter.Times, filter.TimeRange)
			q = addColumnFilters(q, `"review"."rating"`, filter.Rating, filter.Ratings, filter.RatingsRange)
			q = addColumnFilters(q, `"review"."name"`, filter.Name, nil, nil)
		}

		if order != nil {
			strOrder := "DESC"
			if order.IsAscending {
				strOrder = "ASC"
			}
			q = q.Order(fmt.Sprintf("review.%s %s", order.FieldName, strOrder))
		}
		if pageNumber != nil && pageSize != nil {
			q = q.Offset(*pageSize * *pageNumber).Limit(*pageNumber)
		}
		return q.Relation("User").Select()
	})
	if result == nil {
		result = make([]*entity.Review, 0)
	}
	return result, err
}

func (r *ReviewModel) DeleteReview(ctx context.Context, id entity.PK) error {
	model := entity.Review{
		Id: id,
	}
	_, err := r.db.ModelContext(ctx, &model).WherePK().Delete()
	return err
}

func (r *ReviewModel) UpdateReview(ctx context.Context, review *entity.Review) error {
	var (
		setStr string
		args   []interface{}
	)
	if review.Message != nil {
		setStr += "message = ?"
		args = append(args, review.Message)
	}
	if review.Rating > 0 && review.Rating < 6 {
		setStr += ", rating = ?"
		args = append(args, review.Rating)
	}
	if setStr == "" {
		return nil
	}

	return r.db.RunInTransaction(ctx, func(tx *database.TX) error {
		var rev entity.Review
		if err := tx.ModelContext(ctx, &rev).Where("id = ?", review.Id).Select(); err != nil {
			return err
		}
		_, err := tx.ModelContext(ctx, &rev).
			Set(setStr, args...).
			Where("id = ?", review.Id).
			Update()
		return err
	})
}
