package review

import (
	"backendmirea/pr3/internal/entity"
	"backendmirea/pr3/internal/service/review"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type API struct {
	service review.UseCaseReview
}

func NewAPI(s review.UseCaseReview) *API {
	return &API{
		service: s,
	}
}

func errorHandler(c *fiber.Ctx, err error) error {
	return nil
}

func (a *API) Routers(router fiber.Router, middlewares ...fiber.Handler) {
	r := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	r.Get("/", a.getReviews)
	r.Post("/", a.addReview)

	router.Mount("/review", r)
}

func (a *API) getReviews(c *fiber.Ctx) error {
	reviews, err := a.service.GetReviews(c.Context(), nil)
	if err != nil {
		return err
	}
	resultHTML := "<html><body>"
	for _, r := range reviews {
		resultHTML += "<div><h1>" + r.Name + "; " + fmt.Sprint(r.Rating) + "*</h1>"
		resultHTML += "<h2>Опубликовано: " + r.PostedAt.String() + "</h2>"
		if r.Message != nil {
			resultHTML += *r.Message
		}
		resultHTML += "</div><br>"
	}
	resultHTML += "</body></html>"
	c.Set("Content-Type", "text/html;charset=utf-8")
	_, err = c.WriteString(resultHTML)
	return err
}

func (a *API) addReview(c *fiber.Ctx) error {
	form := new(entity.Review)
	if err := c.BodyParser(form); err != nil {
		return err
	}
	return a.service.AddReview(c.Context(), form)
}
