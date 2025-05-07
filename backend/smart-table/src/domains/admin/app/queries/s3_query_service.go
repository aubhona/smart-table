package app

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/smart-table/src/dependencies"
	"github.com/smart-table/src/logging"
	"go.uber.org/zap"
)

type S3QueryService struct {
	client *minio.Client
	bucket string
}

func NewS3QueryService(dependencies *dependencies.Dependencies) *S3QueryService {
	return &S3QueryService{client: dependencies.S3Client, bucket: dependencies.Config.S3.Bucket}
}

func (s *S3QueryService) StoreImage(reader io.Reader, imageSize int64, imageKey string) error {
	ctx := context.Background()
	uploadInfo, err := s.client.PutObject(ctx, s.bucket, imageKey, reader, imageSize, minio.PutObjectOptions{
		ContentType: "image/png",
	})

	if err != nil {
		return err
	}

	logging.GetLogger().Debug(fmt.Sprintf("store image for key = %s", imageKey), zap.Any("info", uploadInfo))

	return nil
}

func (s *S3QueryService) GetImage(imageKey string) (io.ReadCloser, error) {
	ctx := context.Background()
	object, err := s.client.GetObject(ctx, s.bucket, imageKey, minio.GetObjectOptions{})

	if err != nil {
		return nil, err
	}

	return object, nil
}
