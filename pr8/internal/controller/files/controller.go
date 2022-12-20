package files

import (
	"backendmirea/pr3/internal/entity"
	"context"
)

type Controller struct {
	repo Model
}

func NewController(r Model) *Controller {
	return &Controller{repo: r}
}

func (s *Controller) UploadFile(ctx context.Context, file *entity.File) (entity.FileID, error) {
	return s.repo.UploadFile(ctx, file)
}

func (s *Controller) DownloadFile(ctx context.Context, fileID entity.FileID) (*entity.File, error) {
	return s.repo.DownloadFile(ctx, fileID)
}

func (s *Controller) DeleteFile(ctx context.Context, fileID entity.FileID) error {
	return s.repo.DeleteFile(ctx, fileID)
}
