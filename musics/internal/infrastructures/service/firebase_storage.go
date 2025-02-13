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
	obj := bucket.Object("music/" + uuid.New().String() + "_" + fileName)

	wc := obj.NewWriter(ctx)
	wc.ContentType = "audio/mpeg"

	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	// Generate publicUrl
	return fmt.Sprintf("https://storage.googleapis.com/%s/music/%s", f.bucketName, fileName), nil
}

// Delete implements service.FirebaseService.
func (f firebaseStorage) Delete(ctx context.Context, fileName string) error {
	obj := f.client.Bucket(f.bucketName).Object("music/" + fileName)
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
