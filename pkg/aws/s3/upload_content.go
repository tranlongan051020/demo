package s3

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (s *s3Client) UploadContent(
	ctx context.Context,
	buffer io.Reader,
	s3Path *string,
	fileName *string,
) (s3Region string, s3Bucket string, s3Key string, err error) {

	if s3Path == nil {
		todayStr := time.Now().Format("2006/01/02")
		s3Path = &todayStr
	}
	if fileName == nil {
		randomStr, _ := gonanoid.New()
		randomStr += `.json`
		fileName = &randomStr
	}
	s3Key = fmt.Sprintf("%s/%s", *s3Path, *fileName)

	uploader := manager.NewUploader(s.client)
	//nolint:exhaustruct // s3 struct, not our struct
	input := s3.PutObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(s3Key),
		Body:   buffer,
	}
	_, err = uploader.Upload(ctx, &input)
	if err != nil {
		// Print the error and exit.
		slog.ErrorContext(ctx, "UploadContent error #1",
			slog.String("Bucket", s.config.Bucket),
			slog.String("fileName", *fileName),
			slog.Any("err", err),
		)
		return "", "", "", err
	}

	return s.config.Region, s.config.Bucket, s3Key, nil
}
