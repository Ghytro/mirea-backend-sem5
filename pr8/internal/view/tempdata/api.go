package tempdata

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type View struct {
}

func NewView() *View {
	return &View{}
}

func (a *View) Routers(router fiber.Router, authHandler fiber.Handler, middlewares ...fiber.Handler) {
	r := fiber.New()
	r.Use(authHandler)
	for _, m := range middlewares {
		r.Use(m)
	}
	r.Post("/theme", a.setTheme)
	r.Post("/lang", a.setLang)
	router.Mount("/tempdata", r)
}

func (a *View) setTheme(c *fiber.Ctx) error {
	data := string(c.Body())
	c.Cookie(&fiber.Cookie{
		Name:    "theme",
		Value:   data,
		Expires: time.Now().Add(time.Hour * 24 * 7),
	})
	return c.Send(nil)
}

func (a *View) setLang(c *fiber.Ctx) error {
	data := string(c.Body())
	c.Cookie(&fiber.Cookie{
		Name:    "lang",
		Value:   data,
		Expires: time.Now().Add(time.Hour * 24 * 7),
	})
	return c.Send(nil)
}
