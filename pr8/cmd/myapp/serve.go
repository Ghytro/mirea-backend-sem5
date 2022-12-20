package main

import (
	"backendmirea/pr3/internal/controller/auth"
	"backendmirea/pr3/internal/controller/files"
	"backendmirea/pr3/internal/controller/form"
	"backendmirea/pr3/internal/controller/review"
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/logging"
	"backendmirea/pr3/internal/model"
	api "backendmirea/pr3/internal/view"
	auth2 "backendmirea/pr3/internal/view/auth"
	files2 "backendmirea/pr3/internal/view/files"
	form2 "backendmirea/pr3/internal/view/form"
	review2 "backendmirea/pr3/internal/view/review"
	"backendmirea/pr3/internal/view/tempdata"
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

	filesRepo := model.NewFilesModel(fileDB)
	filesController := files.NewController(filesRepo)
	filesApi := files2.NewView(filesController)

	formModel := model.NewFormModel(myDB)
	reviewModel := model.NewReviewModel(myDB)

	formController := form.NewController(formModel)
	reviewController := review.NewController(reviewModel)

	formApi := form2.NewView(formController)
	reviewApi := review2.NewView(reviewController)

	authRepo := model.NewAuthModel(myDB, redisClient)
	authController := auth.NewController(authRepo)
	authApi := auth2.NewView(authController)

	tempDataApi := tempdata.NewView()
	NewApiV1(db, redisClient, formApi, reviewApi, filesApi, authApi, tempDataApi).Listen(":3001")
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
