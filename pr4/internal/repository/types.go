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

	Name string
}

type ReviewFieldName string

const (
	ReviewFieldId       = ReviewFieldName("id")
	ReviewFieldPostedAt = ReviewFieldName("posted_at")
)

type ReviewGetterOrder bool

func (o ReviewGetterOrder) String() (strValue string) {
	if o == ReviewGetterOrderAsc {
		return "ASC"
	}
	return "DESC"
}

const (
	ReviewGetterOrderAsc  = true
	ReviewGetterOrderDesc = false
)

type ReviewOrder struct {
	FieldName ReviewFieldName
	Order     ReviewGetterOrder
}

type FormFilter struct {
	idAbleFilter
	timestampFilter

	Name  *string
	Names []string

	Email  *string
	Emails []string

	Message  *string
	Messages []string
}
