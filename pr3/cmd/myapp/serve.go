package main

import (
	"backendmirea/pr3/internal/api"
	form2 "backendmirea/pr3/internal/api/form"
	review2 "backendmirea/pr3/internal/api/review"
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/repository"
	"backendmirea/pr3/internal/service/form"
	"backendmirea/pr3/internal/service/review"
	"errors"
	"log"
	"os"
	"strings"

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
	NewApiV1(db, formApi, reviewApi).Listen(":3001")
}

func NewApiV1(db *pg.DB, handlers ...api.Handlers) *fiber.App {
	r := fiber.New()

	g := r.Group("/api/v1")

	authHandler := func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		splAuth := strings.Split(auth, " ")
		if len(splAuth) != 2 {
			return errors.New("неверный формат авторизационной строки")
		}
		if splAuth[0] != "Basic" {
			return errors.New("неверный метод авторизации, поддерживается только Basic")
		}
		userNamePass := strings.Split(splAuth[1], ": ")
		if len(userNamePass) != 2 {
			return errors.New("неверный формат авторизационной строки")
		}
		userName, userPass := userNamePass[0], userNamePass[1]
		err := db.ModelContext(c.Context(), (*entity.AuthedUser)(nil)).
			Where(
				"username = ? AND password = crypt(?, gen_salt('bf'))",
				userName,
				userPass,
			).
			Select()
		if err == pg.ErrNoRows {
			return errors.New("неверный логин или пароль")
		}
		return err
	}
	for _, h := range handlers {
		h.Routers(g, authHandler)
	}
	return r
}
