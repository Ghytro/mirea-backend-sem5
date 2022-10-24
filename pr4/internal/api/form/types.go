package form

import "backendmirea/pr3/internal/entity"

type AddFormResponse struct {
	FormId entity.PK `json:"form_id"`
}
