package review

import (
	"backendmirea/pr3/internal/service/review"

	"github.com/gofiber/fiber/v2"
)

type API struct {
	service review.UseCaseReview
}

func NewAPI(s review.UseCaseReview) *API {
	return &API{
		service: s,
	}
}

func errorHandler(c *fiber.Ctx, err error) error {
	return nil
}

func (a *API) Routers(router fiber.Router, middlewares ...fiber.Handler) {
	r := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	r.Get("/", a.getReviews)
	r.Post("/", a.addReview)

	router.Mount("/review", r)
}

func (a *API) getReviews(c *fiber.Ctx) error {
	return nil
}

func (a *API) addReview(c *fiber.Ctx) error {
	return nil
}
