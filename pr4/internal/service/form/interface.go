package form

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

	WithTX(database.DBI) *repository.FormRepository
	RunInTransaction(context.Context, func(*pg.Tx) error) error
}

type Reader interface {
	GetForms(context.Context) ([]*entity.Form, error)
}

type Writer interface {
	AddForm(context.Context, *entity.Form) error
}

type UseCaseForm interface {
	AddForm(context.Context, *entity.Form) error
	GetForms(context.Context) ([]*entity.Form, error)
}
