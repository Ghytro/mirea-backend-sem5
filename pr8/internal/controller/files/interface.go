package files

import (
	"backendmirea/pr3/internal/entity"
	"context"
)

type Reader interface {
	DownloadFile(ctx context.Context, fileID entity.FileID) (*entity.File, error)
}

type Writer interface {
	UploadFile(ctx context.Context, file *entity.File) (entity.FileID, error)
	DeleteFile(ctx context.Context, fileID entity.FileID) error
}

type Model interface {
	Reader
	Writer
}
