package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	mongoConn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url).SetMaxPoolSize(5))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	if err := mongoConn.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	bucket, err := gridfs.NewBucket(mongoConn.Database(DefaultFileDBName))
	if err != nil {
		panic(err)
	}
	return &FileDB{bucket}
}
