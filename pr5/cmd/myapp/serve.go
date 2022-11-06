package main

import (
	"backendmirea/pr3/internal/api"
	files2 "backendmirea/pr3/internal/api/files"
	form2 "backendmirea/pr3/internal/api/form"
	review2 "backendmirea/pr3/internal/api/review"
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/logging"
	"backendmirea/pr3/internal/repository"
	"backendmirea/pr3/internal/service/files"
	"backendmirea/pr3/internal/service/form"
	"backendmirea/pr3/internal/service/review"
	"errors"
	"fmt"
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
	db.AddQueryHook(logging.DBLogger{})
	myDB := &database.DB{DB: db}

	fileDB := database.NewFileDB(db.Context(), os.Getenv("FILE_DB_URL"))

	filesRepo := repository.NewFilesRepository(fileDB)
	filesService := files.NewService(filesRepo)
	filesApi := files2.NewAPI(filesService)

	formRepository := repository.NewFormRepository(myDB)
	reviewRepository := repository.NewReviewRepository(myDB)

	formService := form.NewService(formRepository)
	reviewService := review.NewService(reviewRepository)

	formApi := form2.NewAPI(formService)
	reviewApi := review2.NewAPI(reviewService)
	NewApiV1(db, formApi, reviewApi, filesApi).Listen(":3001")
}

func NewApiV1(db *pg.DB, handlers ...api.Handlers) *fiber.App {
	r := fiber.New()

	g := r.Group("/api/v1")

	g.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

	authHandler := func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		fmt.Println(auth)
		splAuth := strings.Split(auth, " ")
		if len(splAuth) != 3 {
			return errors.New("неверный формат авторизационной строки")
		}
		if splAuth[0] != "Basic" {
			return errors.New("неверный метод авторизации, поддерживается только Basic")
		}
		userName, userPass := splAuth[1], splAuth[2]
		var authedUser entity.AuthedUser
		if err := db.ModelContext(c.Context(), &authedUser).
			Where(
				"username = ? AND password = crypt(?, password)",
				userName,
				userPass,
			).
			Select(); err != nil {
			if err == pg.ErrNoRows {
				return errors.New("неверный логин или пароль")
			}
			return err
		}
		c.Locals("authed_user", &authedUser)
		return c.Next()
	}
	for _, h := range handlers {
		h.Routers(g, authHandler)
	}
	return r
}
