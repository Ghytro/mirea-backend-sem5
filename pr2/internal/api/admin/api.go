package admin

import (
	"fmt"
	"mirea_backend/pr2/internal/service/admin"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func errorHandler(c *fiber.Ctx, err error) error {
	_, _err := c.WriteString(fmt.Sprintf("<h1>Произошла ошибка: %v</h1>", err))
	return _err
}

type API struct {
	service *admin.Service
}

func NewAPI(service *admin.Service) *API {
	return &API{
		service: service,
	}
}

func (a *API) runCmd(c *fiber.Ctx) error {
	splittedCommand := strings.Split(c.Query("cmd"), " ")
	command := splittedCommand[0]
	args := splittedCommand[1:]
	result, err := a.service.ExecCommand(c.Context(), command, args...)
	if err != nil {
		return err
	}
	_, err = c.WriteString(fmt.Sprintf(
		`<div style="font-family:courier, courier new, serif;">
			$ %s<br>
			%s
		</div>`,
		c.Query("cmd"),
		result,
	))
	if err != nil {
		return err
	}
	return c.Next()
}

func (a *API) Routers(router fiber.Router) {
	r := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	r.Use(func(c *fiber.Ctx) error {
		if _, err := c.WriteString("<html><body>"); err != nil {
			return err
		}
		return c.Next()
	})
	r.Get("/", a.runCmd)
	r.Use(func(c *fiber.Ctx) error {
		_, err := c.WriteString("</body></html>")
		return err
	})
	router.Mount("/admin", r)
}
