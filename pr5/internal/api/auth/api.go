package auth

import (
	"backendmirea/pr3/internal/entity"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type API struct {
	service UseCase
}

func NewAPI(s UseCase) *API {
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
	r.Post("/", a.startSession)
	router.Mount("/auth", r)
}

func (a *API) startSession(c *fiber.Ctx) error {
	fmt.Println("here")
	var tokenReq NewTokenRequest
	if err := c.BodyParser(&tokenReq); err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusBadRequest,
			Err: &entity.ServerError{
				Message:   err.Error(),
				Location:  "auth",
				ErrorCode: -1,
			},
		}
	}
	fmt.Println("here")
	token, expires, err := a.service.StartSession(c.Context(), tokenReq.Username, tokenReq.Password)
	if err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusUnauthorized,
			Err: &entity.ServerError{
				Message:   err.Error(),
				Location:  "auth",
				ErrorCode: -1,
			},
		}
	}
	fmt.Println("here")
	c.Cookie(&fiber.Cookie{
		Name:    "username",
		Value:   tokenReq.Username,
		Expires: time.Now().Add(time.Hour * 24 * 7),
	})
	return c.JSON(&entity.AuthToken{
		Token:   token,
		Expires: expires,
	})
}
