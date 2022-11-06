package files

import (
	"backendmirea/pr3/internal/entity"
	"context"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) UploadFile(ctx context.Context, file *entity.File) (entity.FileID, error) {
	return s.repo.UploadFile(ctx, file)
}

func (s *Service) DownloadFile(ctx context.Context, fileID entity.FileID) (*entity.File, error) {
	return s.repo.DownloadFile(ctx, fileID)
}

func (s *Service) DeleteFile(ctx context.Context, fileID entity.FileID) error {
	return s.repo.DeleteFile(ctx, fileID)
}
