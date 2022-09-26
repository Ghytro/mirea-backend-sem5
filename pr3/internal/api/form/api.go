package form

import (
	"backendmirea/pr3/internal/service/form"

	"github.com/gofiber/fiber/v2"
)

type API struct {
	service form.UseCaseForm
}

func NewAPI(s form.UseCaseForm) *API {
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

	r.Get("/", a.getForms)
	r.Post("/", a.addForm)

	router.Mount("/form", r)
}

func (a *API) getForms(c *fiber.Ctx) error {
	return nil
}

func (a *API) addForm(c *fiber.Ctx) error {
	return nil
}
