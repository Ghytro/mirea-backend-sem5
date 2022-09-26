package repository

import (
	"backendmirea/pr3/internal/utils"
	"time"
)

type ReviewFilter struct {
	Id  *int
	Ids *utils.Range[int]

	TimeRange *utils.Range[time.Time]

	Rating  *int
	Ratings *utils.Range[int]
}
