package repository

import (
	"backendmirea/pr3/internal/utils"
	"time"
)

type idAbleFilter struct {
	Id       *int              `json:"id"`
	Ids      []int             `json:"ids"`
	IdsRange *utils.Range[int] `json:"ids_range"`
}

type timestampFilter struct {
	Time      *time.Time              `json:"time"`
	Times     []time.Time             `json:"times"`
	TimeRange *utils.Range[time.Time] `json:"time_range"`
}

type ReviewFilter struct {
	idAbleFilter
	timestampFilter

	Rating       *int              `json:"rating"`
	Ratings      []int             `json:"ratings"`
	RatingsRange *utils.Range[int] `json:"ratings_range"`

	Name *string `json:"name"`
}

type ReviewFieldName string

const (
	ReviewFieldId       = ReviewFieldName("id")
	ReviewFieldPostedAt = ReviewFieldName("posted_at")
)

type ReviewOrder struct {
	FieldName   ReviewFieldName `json:"field_name"`
	IsAscending bool            `json:"is_ascending"`
}

type FormFilter struct {
	idAbleFilter
	timestampFilter

	Name  *string  `json:"name"`
	Names []string `json:"names"`

	Email  *string  `json:"email"`
	Emails []string `json:"emails"`

	Message  *string  `json:"message"`
	Messages []string `json:"messages"`
}
