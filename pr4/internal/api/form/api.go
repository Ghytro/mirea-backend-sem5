package form

import (
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/service/form"
	"strconv"

	"github.com/go-pg/pg/v10"
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
	if errResponse, ok := err.(*entity.ErrResponse); ok {
		if _, ok := errResponse.Unwrap().(*entity.ServerError); !ok {
			errResponse.Err = &entity.ServerError{
				Message:   errResponse.Err.Error(),
				Location:  "unknown",
				ErrorCode: -1,
			}
		}
		return c.Status(errResponse.StatusCode).JSON(errResponse.Err.(*entity.ServerError))
	}
	resp := &entity.ErrResponse{
		StatusCode: fiber.StatusInternalServerError,
		Err: &entity.ServerError{
			Message:   err.Error(),
			Location:  "unknown",
			ErrorCode: -1,
		},
	}
	return c.Status(resp.StatusCode).JSON(resp.Err.(*entity.ServerError))
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
	r.Delete("/:id", a.deleteForm)

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

func (a *API) deleteForm(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusBadRequest,
			Err: &entity.ServerError{
				Message:   "incorrect id passed",
				Location:  "form deletion",
				ErrorCode: -1,
			},
		}
	}
	if err := a.service.DeleteForm(c.Context(), entity.PK(id)); err != nil {
		if err == pg.ErrNoRows {
			return &entity.ErrResponse{
				StatusCode: fiber.StatusBadRequest,
				Err: &entity.ServerError{
					Message:   "no form found with that id",
					Location:  "form deletion",
					ErrorCode: -1,
				},
			}
		}
		return &entity.ErrResponse{
			StatusCode: fiber.StatusInternalServerError,
			Err: &entity.ServerError{
				Message:   "unable to delete form",
				Location:  "form deletion",
				ErrorCode: -1,
			},
		}
	}
	return c.Status(fiber.StatusOK).Send(nil)
}
