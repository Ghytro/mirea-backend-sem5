package main

import (
	"backendmirea/pr3/internal/api"
	form2 "backendmirea/pr3/internal/api/form"
	review2 "backendmirea/pr3/internal/api/review"
	"backendmirea/pr3/internal/repository"
	"backendmirea/pr3/internal/service/form"
	"backendmirea/pr3/internal/service/review"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

func serve() {
	opt, err := pg.ParseURL(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	db := pg.Connect(opt)

	formRepository := repository.NewFormRepository(db)
	reviewRepository := repository.NewReviewRepository(db)

	formService := form.NewService(formRepository)
	reviewService := review.NewService(reviewRepository)

	formApi := form2.NewAPI(formService)
	reviewApi := review2.NewAPI(reviewService)
	NewApiV1(formApi, reviewApi).Listen(":3001")
}

func NewApiV1(handlers ...api.Handlers) *fiber.App {
	r := fiber.New()

	g := r.Group("/api/v1")
	for _, h := range handlers {
		h.Routers(g)
	}
	return r
}
