package files

import (
	"backendmirea/pr3/internal/entity"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
)

func errorHandler(c *fiber.Ctx, err error) error {
	if errResponse, ok := err.(*entity.ErrResponse); ok {
		if _, ok := errResponse.Unwrap().(*entity.ServerError); !ok {
			errResponse.Err = &entity.ServerError{
				Message:   errResponse.Err.Error(),
				Location:  "unknown",
				ErrorCode: -1,
			}
		}
		return c.Status(errResponse.StatusCode).JSON(errResponse.Err.(*entity.ServerError))
	}
	resp := &entity.ErrResponse{
		StatusCode: fiber.StatusInternalServerError,
		Err: &entity.ServerError{
			Message:   err.Error(),
			Location:  "unknown",
			ErrorCode: -1,
		},
	}
	return c.Status(resp.StatusCode).JSON(resp.Err.(*entity.ServerError))
}

type API struct {
	service UseCase
}

func NewAPI(s UseCase) *API {
	return &API{service: s}
}

func (a *API) Routers(router fiber.Router, authHandler fiber.Handler, middlewares ...fiber.Handler) {
	r := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	for _, m := range middlewares {
		r.Use(m)
	}
	r.Get("/:id", a.getFile)
	r.Post("/", a.postFile)
	r.Delete("/:id", a.deleteFile)
	router.Mount("/files", r)
}

func (a *API) getFile(c *fiber.Ctx) error {
	fileID, err := entity.ParseFileID(c.Params("id"))
	if err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusBadRequest,
			Err: &entity.ServerError{
				Message:   err.Error(),
				Location:  "ошибка при получении файла",
				ErrorCode: -1,
			},
		}
	}
	file, err := a.service.DownloadFile(c.Context(), fileID)
	if err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusBadRequest,
			Err: &entity.ServerError{
				Message:   err.Error(),
				Location:  "ошибка при получении файла",
				ErrorCode: -1,
			},
		}
	}
	tempFile, err := os.CreateTemp("/tmp", "temp-*.pdf")
	if err != nil {
		return err
	}
	if _, err := io.Copy(tempFile, file.File); err != nil {
		return err
	}
	return c.Download(tempFile.Name(), file.OrigFileName)
}

func (a *API) postFile(c *fiber.Ctx) error {
	fileID, err := a.service.UploadFile(c.Context(), &entity.File{
		OrigFileName: c.Get("X-Filename"),
		File:         c.Context().RequestBodyStream(),
	})
	if err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusBadRequest,
			Err: &entity.ServerError{
				Message:   err.Error(),
				Location:  "не удалось загрузить файл",
				ErrorCode: -1,
			},
		}
	}
	return c.JSON(struct {
		FileID string `json:"file_id"`
	}{FileID: fileID.String()})
}

func (a *API) deleteFile(c *fiber.Ctx) error {
	fileID, err := entity.ParseFileID(c.Params("id"))
	if err != nil {
		return &entity.ErrResponse{
			StatusCode: fiber.StatusBadRequest,
			Err: &entity.ServerError{
				Message:   err.Error(),
				Location:  "ошибка при получении файла",
				ErrorCode: -1,
			},
		}
	}
	return a.service.DeleteFile(c.Context(), fileID)
}
