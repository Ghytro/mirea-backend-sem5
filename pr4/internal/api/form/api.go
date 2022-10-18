package form

import (
	"backendmirea/pr3/internal/entity"
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
	_, _err := c.WriteString("<h1>Произошла ошибка: " + err.Error() + "</h1>")
	return _err
}

func (a *API) Routers(router fiber.Router, authHandler fiber.Handler, middlewares ...fiber.Handler) {
	r := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	for _, m := range middlewares {
		r.Use(m)
	}

	r.Post("/", a.addForm)
	r.Use(authHandler)
	r.Get("/", a.getForms)

	router.Mount("/form", r)
}

func (a *API) getForms(c *fiber.Ctx) error {
	forms, err := a.service.GetForms(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(forms)
}

func (a *API) addForm(c *fiber.Ctx) error {
	form := new(entity.Form)
	if err := c.BodyParser(form); err != nil {
		return err
	}
	return a.service.AddForm(c.Context(), form)
}
