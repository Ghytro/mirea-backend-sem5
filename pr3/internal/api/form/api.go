package form

import (
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/service/form"
	"errors"

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
	adminPassword := c.Query("pass")
	if adminPassword != "123123" {
		return errors.New("неверный пароль для доступа к админской странице")
	}
	forms, err := a.service.GetForms(c.Context())
	if err != nil {
		return err
	}
	resultHTML := "<html><body>"
	for _, f := range forms {
		resultHTML += "<div><h1>Заявка от: " + f.Name + "</h1>"
		resultHTML += "<h2>Прислана: " + f.SentAt.String() + "<br>"
		resultHTML += "Отвечать на email: " + f.Email + "</h2>"
		resultHTML += f.Message + "</div><br>"
	}
	resultHTML += "</body></html>"
	c.Set("Content-Type", "text/html;charset=utf-8")
	_, err = c.WriteString(resultHTML)
	return err
}

func (a *API) addForm(c *fiber.Ctx) error {
	form := new(entity.Form)
	if err := c.BodyParser(form); err != nil {
		return err
	}
	return a.service.AddForm(c.Context(), form)
}
