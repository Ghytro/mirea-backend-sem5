package form

import (
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/repository"
	"context"
)

type Repository interface {
	Reader
	Writer

	WithTX(database.DBI) *repository.FormRepository
	RunInTransaction(context.Context, func(*database.TX) error) error
}

type Reader interface {
	GetForms(context.Context) ([]*entity.Form, error)
	GetForm(context.Context, entity.PK) (*entity.Form, error)
}

type Writer interface {
	AddForm(context.Context, *entity.Form) error
	UpdateForm(context.Context, *entity.Form) error
	DeleteForm(context.Context, entity.PK) error
}

type UseCaseForm interface {
	AddForm(context.Context, *entity.Form) error
	GetForms(context.Context) ([]*entity.Form, error)
	UpdateForm(context.Context, entity.PK, *entity.Form) error
	DeleteForm(context.Context, entity.PK, entity.PK) error
}
