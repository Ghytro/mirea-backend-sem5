package sorter

import (
	"fmt"
	"mirea_backend/pr2/internal/service/sorter"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func errorHandler(c *fiber.Ctx, err error) error {
	_, _err := c.WriteString(fmt.Sprintf("<h1>Произошла ошибка: %v</h1>", err))
	return _err
}

type API struct {
	service *sorter.Service
}

func NewAPI(service *sorter.Service) *API {
	return &API{
		service: service,
	}
}

func (a *API) sortArray(c *fiber.Ctx) error {
	strArray := strings.Split(c.Query("array"), ",")
	intArr := make([]int, 0, len(strArray))
	for _, el := range strArray {
		intEl, err := strconv.Atoi(el)
		if err != nil {
			return err
		}
		intArr = append(intArr, intEl)
	}
	a.service.SortArray(c.Context(), intArr)
	for i, el := range intArr {
		strArray[i] = fmt.Sprint(el)
	}
	response := strings.Join(strArray, ",")
	if _, err := c.WriteString(response); err != nil {
		return err
	}
	return c.Next()
}

func (a *API) Routers(router fiber.Router) {
	r := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	r.Use(func(c *fiber.Ctx) error {
		if _, err := c.WriteString("<html><body>"); err != nil {
			return err
		}
		return c.Next()
	})
	r.Get("/", a.sortArray)
	r.Use(func(c *fiber.Ctx) error {
		_, err := c.WriteString("</body></html>")
		return err
	})
	router.Mount("/sorter", r)
}
