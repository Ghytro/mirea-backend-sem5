package repository

import (
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"bytes"
	"context"
	"io"
)

type FileRepository struct {
	db *database.FileDB
}

func NewFilesRepository(db *database.FileDB) *FileRepository {
	return &FileRepository{db: db}
}

func (r *FileRepository) UploadFile(ctx context.Context, file *entity.File) (entity.FileID, error) {
	uploadStream, err := r.db.OpenUploadStream(file.OrigFileName)
	if err != nil {
		return entity.NilFileID, err
	}
	defer uploadStream.Close()

	fileID := uploadStream.FileID.(entity.FileID)
	data, err := io.ReadAll(file.File)
	if err != nil {
		return entity.NilFileID, err
	}
	_, err = uploadStream.Write(data)
	return fileID, err
}

func (r *FileRepository) DownloadFile(ctx context.Context, fileID entity.FileID) (*entity.File, error) {
	downloadStream, err := r.db.OpenDownloadStream(fileID)
	if err != nil {
		return nil, err
	}
	defer downloadStream.Close()
	deadline, ok := ctx.Deadline()
	if ok {
		downloadStream.SetReadDeadline(deadline)
	}
	var b []byte
	_, err = downloadStream.Read(b)
	if err != nil {
		return nil, err
	}
	var result bytes.Buffer
	_, err = result.Write(b)
	return &entity.File{
		OrigFileName: downloadStream.GetFile().Name,
		File:         &result,
	}, err
}

func (r *FileRepository) DeleteFile(ctx context.Context, fileID entity.FileID) error {
	return r.db.DeleteContext(ctx, fileID)
}
