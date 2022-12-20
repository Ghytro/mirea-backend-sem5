package files

import (
	"backendmirea/pr3/internal/entity"
	"context"
)

type UseCase interface {
	UploadFile(ctx context.Context, file *entity.File) (entity.FileID, error)
	DownloadFile(ctx context.Context, fileID entity.FileID) (*entity.File, error)
	DeleteFile(ctx context.Context, fileID entity.FileID) error
}
