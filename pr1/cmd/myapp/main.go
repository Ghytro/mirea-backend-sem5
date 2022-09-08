package main

import (
	"backendmirea/pr1/internal/entity"
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var dbConn *pg.DB

func init() {
	opt, err := pg.ParseURL(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	dbConn = pg.Connect(opt)
}

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.Next()
	})
	api := app.Group("/api")
	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello from server! (get handler)")
	})
	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello from server! (post handler)")
	})
	api.Get("/db", func(c *fiber.Ctx) error {
		result := &entity.GenericDBResponse{
			Id: 1,
		}
		if err := dbConn.Model(result).WherePK().Select(); err != nil {
			log.Println(err)
			c.Status(fiber.StatusInternalServerError).SendString("something wrong with db")
		}
		return c.SendString(fmt.Sprintf("column1: %s, column2: %d", result.Column1, result.Column2))
	})
	app.Listen(":3001")
}
