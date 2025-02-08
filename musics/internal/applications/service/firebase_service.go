package service

import (
	"context"
	"io"
)

type FirebaseService interface {
	Upload(ctx context.Context, fileName string, file io.Reader) (string, error)
	Delete(ctx context.Context, fileName string) error
}
