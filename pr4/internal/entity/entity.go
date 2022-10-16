package entity

import (
	"fmt"
	"time"
)

type PK uint

type Form struct {
	tableName struct{} `pg:"forms"`

	Name    string    `pg:"name" json:"name" form:"name"`
	Email   string    `pg:"email" json:"email" form:"email"`
	Message string    `pg:"message" json:"message" form:"message"`
	SentAt  time.Time `pg:"sent_at" json:"-"`
}

type Review struct {
	tableName struct{} `pg:"reviews"`

	Id PK `pg:"id,pk" json:"id" form:"id"`

	Name     string    `pg:"name" json:"name" form:"name"`
	Rating   int       `pg:"rating" json:"rating" form:"rating"`
	Message  *string   `pg:"message" json:"message,omitempty" form:"message"`
	PostedAt time.Time `pg:"posted_at" json:"posted_at"`
}

type AuthedUser struct {
	tableName struct{} `pg:"users"`

	Id PK `pg:"id,pk" json:"id" form:"id"`

	UserName string `pg:"username"`
	Password string `pg:"password"`
}

type ErrResponse struct {
	StatusCode int
	Err        error
}

func (e *ErrResponse) Error() string {
	return e.Err.Error()
}

func (e *ErrResponse) Unwrap() error {
	return e.Err
}

type ServerError struct {
	Message   string `json:"message"`
	Location  string `json:"location"`
	ErrorCode int    `json:"code"`
	BaseError error
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("%s; %s: %s (%d)", e.Message, e.Location, e.BaseError.Error(), e.ErrorCode)
}

func (e *ServerError) Unwrap() error {
	return e.BaseError
}
