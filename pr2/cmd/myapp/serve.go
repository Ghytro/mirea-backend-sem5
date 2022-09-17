package main

import (
	"mirea_backend/pr2/internal/api"
	drawer2 "mirea_backend/pr2/internal/api/drawer"
	sorter2 "mirea_backend/pr2/internal/api/sorter"
	"mirea_backend/pr2/internal/service/drawer"
	"mirea_backend/pr2/internal/service/sorter"

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

	sorterService := sorter.NewService()
	sorterApi := sorter2.NewAPI(sorterService)

	Handlers(drawerApi, sorterApi).Listen(":3001")
}
