package service

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/ardwiinoo/micro-music/musics/internal/applications/service"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"io"
)

type firebaseStorage struct {
	client     *storage.Client
	bucketName string
}

// Upload implements service.FirebaseService.
func (f firebaseStorage) Upload(ctx context.Context, fileName string, file io.Reader) (string, error) {
	bucket := f.client.Bucket(f.bucketName)

	uniqueFileName := uuid.New().String() + "_" + fileName
	obj := bucket.Object("music/" + uniqueFileName)

	wc := obj.NewWriter(ctx)
	wc.ContentType = "audio/mpeg"

	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"https://firebasestorage.googleapis.com/v0/b/%s/o/music%%2F%s?alt=media",
		f.bucketName, uniqueFileName,
	), nil
}

// Delete implements service.FirebaseService.
func (f firebaseStorage) Delete(ctx context.Context, uniqueFileName string) error {
	obj := f.client.Bucket(f.bucketName).Object("music/" + uniqueFileName)
	return obj.Delete(ctx)
}

func NewFirebaseStorage(credentialFile string, bucketName string) (service.FirebaseService, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFile))
	if err != nil {
		return nil, err
	}

	return &firebaseStorage{
		client:     client,
		bucketName: bucketName,
	}, nil
}
