package review

import (
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/service/review"
	"fmt"
	"strconv"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

func errParsingJson(location string, err error) *entity.ServerError {
	return &entity.ServerError{
		Message:   "unable to parse json",
		Location:  location,
		ErrorCode: -1,
		BaseError: err,
	}
}

type API struct {
	service review.UseCaseReview
}

func NewAPI(s review.UseCaseReview) *API {
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

	r.Get("/", a.getReviews)
	r.Post("/", a.addReview)
	r.Use(authHandler)
	r.Delete("/:id", a.deleteReview)

	router.Mount("/review", r)
}

func (a *API) getReviews(c *fiber.Ctx) error {
	var model GetReviewsRequest
	if err := c.BodyParser(&model); err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusBadRequest,
			Err:        errParsingJson("reviews getter", err),
		}
	}
	reviews, err := a.service.GetReviews(c.Context(), model.Filter, model.Order, model.Page, model.PageSize)
	if err != nil {
		if err == pg.ErrNoRows {
			return c.JSON([]*entity.Review{})
		}
		return &entity.ErrResponse{
			StatusCode: fiber.StatusInternalServerError,
			Err: &entity.ServerError{
				Message:   err.Error(),
				Location:  "review getter",
				ErrorCode: -1,
				BaseError: err,
			},
		}
	}
	return c.JSON(reviews)
}

func (a *API) addReview(c *fiber.Ctx) error {
	form := new(entity.Review)
	if err := c.BodyParser(form); err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusBadRequest,
			Err:        errParsingJson("review add", err),
		}
	}
	if err := a.service.AddReview(c.Context(), form); err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusInternalServerError,
			Err: &entity.ServerError{
				Message:   fmt.Sprintf("unable to add review: %s", err.Error()),
				Location:  "add review",
				ErrorCode: -1,
				BaseError: err,
			},
		}
	}
	return c.Status(fiber.StatusOK).Send(nil)
}

func (a *API) deleteReview(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusBadRequest,
			Err: &entity.ServerError{
				Message:   "incorrect format of id",
				Location:  "review deletion",
				ErrorCode: -1,
			},
		}
	}
	if err := a.service.DeleteReview(c.Context(), entity.PK(id)); err != nil {
		if err == pg.ErrNoRows {
			return &entity.ErrResponse{
				StatusCode: fiber.StatusBadRequest,
				Err: &entity.ServerError{
					Message:   "review with id not found",
					Location:  "review deletion",
					ErrorCode: -1,
				},
			}
		}
		return &entity.ErrResponse{
			StatusCode: fiber.StatusInternalServerError,
			Err: &entity.ServerError{
				Message:   err.Error(),
				Location:  "review deletion",
				ErrorCode: -1,
			},
		}
	}
	return c.Status(fiber.StatusOK).Send(nil)
}
