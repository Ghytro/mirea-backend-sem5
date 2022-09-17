package main

import (
	"mirea_backend/pr2/internal/api"
	drawer2 "mirea_backend/pr2/internal/api/drawer"
	"mirea_backend/pr2/internal/service/drawer"

	"github.com/gofiber/fiber/v2"
)

func Handlers(handlers ...api.Handlers) *fiber.App {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.Next()
	})

	for _, h := range handlers {
		h.Routers(app)
	}
	return app
}

func serve() {
	drawerService := drawer.NewService()
	drawerApi := drawer2.NewAPI(drawerService)

	Handlers(drawerApi).Listen(":3001")
}
