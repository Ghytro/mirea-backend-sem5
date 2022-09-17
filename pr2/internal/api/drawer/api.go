package drawer

import (
	"errors"
	"fmt"
	"mirea_backend/pr2/internal/service/drawer"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func errorHandler(c *fiber.Ctx, err error) error {
	_, _err := c.WriteString(fmt.Sprintf("<h1>Произошла ошибка: %v</h1>", err))
	return _err
}

type API struct {
	service *drawer.Service
}

func NewAPI(service *drawer.Service) *API {
	return &API{
		service: service,
	}
}

func (a *API) drawShape(c *fiber.Ctx) error {
	shapeID, err := strconv.Atoi(c.Query("num"))
	if err != nil {
		return errors.New("ошибка при получении id фигуры")
	}
	boundingBox, err := a.service.DrawShape(c.Context(), uint32(shapeID))
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = c.WriteString(boundingBox.HTMLString())
	if err != nil {
		return err
	}
	return c.Next()
}

func (a *API) Routers(router fiber.Router) {
	r := fiber.New(
		fiber.Config{
			ErrorHandler: errorHandler,
		})
	r.Use(func(c *fiber.Ctx) error {
		if _, err := c.WriteString("<html><body>"); err != nil {
			return err
		}
		return c.Next()
	})
	r.Get("/", a.drawShape)
	r.Use(func(c *fiber.Ctx) error {
		_, err := c.WriteString("</body></html>")
		return err
	})
	router.Mount("/drawer", r)
}
