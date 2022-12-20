package model

import (
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"bytes"
	"context"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileModel struct {
	db *database.FileDB
}

func NewFilesModel(db *database.FileDB) *FileModel {
	return &FileModel{db: db}
}

func (r *FileModel) UploadFile(ctx context.Context, file *entity.File) (entity.FileID, error) {
	uploadStream, err := r.db.OpenUploadStream(file.OrigFileName)
	if err != nil {
		return entity.NilFileID, err
	}
	defer uploadStream.Close()

	fileID := entity.FileID{uploadStream.FileID.(primitive.ObjectID)}
	data, err := io.ReadAll(file.File)
	if err != nil {
		return entity.NilFileID, err
	}
	_, err = uploadStream.Write(data)
	return fileID, err
}

func (r *FileModel) DownloadFile(ctx context.Context, fileID entity.FileID) (*entity.File, error) {
	downloadStream, err := r.db.OpenDownloadStream(fileID.ObjectID)
	if err != nil {
		return nil, err
	}
	defer downloadStream.Close()
	deadline, ok := ctx.Deadline()
	if ok {
		downloadStream.SetReadDeadline(deadline)
	}
	b := make([]byte, downloadStream.GetFile().Length)
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

func (r *FileModel) DeleteFile(ctx context.Context, fileID entity.FileID) error {
	return r.db.DeleteContext(ctx, fileID)
}
