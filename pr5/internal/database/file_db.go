package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DefaultFileDBName = "mydb"
const DefaultFileCollectionName = "fs.files"

type UploadedFile struct {
	File         *os.File
	OrigFileName string
}

type FileDB struct {
	*gridfs.Bucket
}

func NewFileDB(ctx context.Context, url string) *FileDB {
	opts := options.Client()
	opts.ApplyURI(os.Getenv("FILE_DB_URL"))
	opts.SetMaxPoolSize(5) // должно хватить
	mongoConn, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	bucket, err := gridfs.NewBucket(mongoConn.Database(DefaultFileDBName))
	if err != nil {
		panic(err)
	}
	return &FileDB{bucket}
}
