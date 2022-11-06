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

func errParsingJson(location string, err error) *entity.ServerError {
	return &entity.ServerError{
		Message:   "unable to parse json",
		Location:  location,
		ErrorCode: -1,
		BaseError: err,
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

	r.Use(authHandler)
	r.Post("/", a.addForm)
	r.Get("/", a.getForms)
	r.Delete("/:id", a.deleteForm)
	r.Patch("/:id", a.updateForm)

	router.Mount("/form", r)
}

func (a *API) getForms(c *fiber.Ctx) error {
	user, ok := c.Locals("authed_user").(*entity.AuthedUser)
	if !ok {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusUnauthorized,
			Err: &entity.ServerError{
				Message:   "debug: unable to get auth entity",
				Location:  "form get",
				ErrorCode: -1,
			},
		}
	}
	if !user.IsAdmin {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusForbidden,
			Err: &entity.ServerError{
				Message:   "accessible only to admins",
				Location:  "form get",
				ErrorCode: -1,
			},
		}
	}
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
	user, ok := c.Locals("authed_user").(*entity.AuthedUser)
	if !ok {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusUnauthorized,
			Err: &entity.ServerError{
				Message:   "debug: unable to get auth entity",
				Location:  "form get",
				ErrorCode: -1,
			},
		}
	}
	form.UserId = user.Id
	if err := a.service.AddForm(c.Context(), form); err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusBadRequest,
			Err: &entity.ServerError{
				Message:   err.Error(),
				Location:  "add form",
				ErrorCode: -1,
			},
		}
	}
	return c.JSON(AddFormResponse{FormId: form.Id})
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
	user, ok := c.Locals("authed_user").(*entity.AuthedUser)
	if !ok {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusUnauthorized,
			Err: &entity.ServerError{
				Message:   "debug: unable to get auth entity",
				Location:  "form get",
				ErrorCode: -1,
			},
		}
	}
	if err := a.service.DeleteForm(c.Context(), user.Id, entity.PK(id)); err != nil {
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

func (a *API) updateForm(c *fiber.Ctx) error {
	formId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusBadRequest,
			Err: &entity.ServerError{
				Message:   "incorrect id passed",
				Location:  "form update",
				ErrorCode: -1,
			},
		}
	}
	var f entity.Form
	if err := c.BodyParser(&f); err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusBadRequest,
			Err:        errParsingJson("form update", err),
		}
	}
	f.Id = entity.PK(formId)
	user, ok := c.Locals("authed_user").(*entity.AuthedUser)
	if !ok {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusUnauthorized,
			Err: &entity.ServerError{
				Message:   "debug: unable to get auth entity",
				Location:  "form update",
				ErrorCode: -1,
			},
		}
	}
	if err := a.service.UpdateForm(c.Context(), user.Id, &f); err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusBadRequest,
			Err: &entity.ServerError{
				Message:   err.Error(),
				Location:  "update form",
				ErrorCode: -1,
			},
		}
	}
	return c.Send(nil)
}
