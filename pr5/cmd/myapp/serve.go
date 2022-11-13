package main

import (
	"backendmirea/pr3/internal/api"
	auth2 "backendmirea/pr3/internal/api/auth"
	files2 "backendmirea/pr3/internal/api/files"
	form2 "backendmirea/pr3/internal/api/form"
	review2 "backendmirea/pr3/internal/api/review"
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/logging"
	"backendmirea/pr3/internal/repository"
	"backendmirea/pr3/internal/service/auth"
	"backendmirea/pr3/internal/service/files"
	"backendmirea/pr3/internal/service/form"
	"backendmirea/pr3/internal/service/review"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v9"
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

	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewClient(opts)

	filesRepo := repository.NewFilesRepository(fileDB)
	filesService := files.NewService(filesRepo)
	filesApi := files2.NewAPI(filesService)

	formRepository := repository.NewFormRepository(myDB)
	reviewRepository := repository.NewReviewRepository(myDB)

	formService := form.NewService(formRepository)
	reviewService := review.NewService(reviewRepository)

	formApi := form2.NewAPI(formService)
	reviewApi := review2.NewAPI(reviewService)

	authRepo := repository.NewAuthRepository(myDB, redisClient)
	authService := auth.NewService(authRepo)
	authApi := auth2.NewAPI(authService)
	NewApiV1(db, redisClient, formApi, reviewApi, filesApi, authApi).Listen(":3001")
}

func NewApiV1(db *pg.DB, cacheDB *redis.Client, handlers ...api.Handlers) *fiber.App {
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
		if len(splAuth) != 2 {
			return errors.New("неверный формат авторизационной строки")
		}
		if splAuth[0] != "Token" {
			return errors.New("неверный метод авторизации, поддерживается только Token")
		}
		token := splAuth[1]
		var authedUser entity.AuthedUser
		err := cacheDB.HGetAll(c.Context(), token).Scan(&authedUser)
		if err != nil {
			log.Println(err)
			return &entity.ErrResponse{
				StatusCode: fiber.StatusUnauthorized,
				Err: &entity.ServerError{
					Message:   "unauthorized",
					Location:  "auth",
					ErrorCode: -1,
				},
			}
		}
		c.Locals("authed_user", &authedUser)
		return c.Next()
	}
	for _, h := range handlers {
		h.Routers(g, authHandler)
	}
	return r
}
