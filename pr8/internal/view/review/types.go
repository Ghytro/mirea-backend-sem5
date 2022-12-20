package review

import (
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/model"
)

type GetReviewsRequest struct {
	Filter *model.ReviewFilter `json:"filter"`
	Order  *model.ReviewOrder  `json:"order"`

	Page     *int `json:"page"`
	PageSize *int `json:"page_size"`
}

type AddReviewResponse struct {
	ReviewId entity.PK `json:"review_id"`
}
