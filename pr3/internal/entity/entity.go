package entity

import "time"

type Form struct {
	tableName struct{} `pg:"forms"`

	Name    string    `pg:"name" json:"name" form:"name"`
	Email   string    `pg:"email" json:"email" form:"email"`
	Message string    `pg:"message" json:"message" form:"message"`
	SentAt  time.Time `pg:"sent_at" json:"-"`
}

type Review struct {
	tableName struct{} `pg:"reviews"`

	Name     string    `pg:"name" json:"name" form:"name"`
	Rating   int       `pg:"rating" json:"rating" form:"rating"`
	Message  *string   `pg:"message" json:"message,omitempty" form:"message"`
	PostedAt time.Time `pg:"posted_at" json:"-"`
}
