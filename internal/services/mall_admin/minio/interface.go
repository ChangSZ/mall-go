package minio

import (
	"context"
	"mime/multipart"
)

var _ Service = (*service)(nil)

// MinIO对象存储管理Service
type Service interface {
	i()

	Upload(ctx context.Context, file multipart.File, filename string) (string, string, error)

	Delete(ctx context.Context, objectName string) error

	PresignedURL(ctx context.Context, bucketName, objectName string) (string, error)
}
