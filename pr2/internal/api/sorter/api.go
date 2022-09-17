package sorter

import "mirea_backend/pr2/internal/service/sorter"

type API struct {
	service *sorter.Service
}

func NewAPI(service *sorter.Service) *API {
	return &API{
		service: service,
	}
}
