package api

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	Routers(router fiber.Router)
}
