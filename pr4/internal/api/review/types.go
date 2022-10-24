package review

import (
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/repository"
)

type GetReviewsRequest struct {
	Filter *repository.ReviewFilter `json:"filter"`
	Order  *repository.ReviewOrder  `json:"order"`

	Page     *int `json:"page"`
	PageSize *int `json:"page_size"`
}

type AddReviewResponse struct {
	ReviewId entity.PK `json:"review_id"`
}
