package s3

import (
	"context"
	"log/slog"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *s3Client) PresignGetObject(
	ctx context.Context,
	s3ObjectKey string,
	lifetimeSecs int64,
) (getObjectURL string, err error) {
	//nolint:exhaustruct // s3 struct, not our struct
	request, err := s.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(s3ObjectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(lifetimeSecs * int64(time.Second))
	})
	if err != nil {
		slog.ErrorContext(ctx, "PresignGetObject error #1",
			slog.String("Bucket", s.config.Bucket),
			slog.String("s3ObjectKey", s3ObjectKey),
			slog.Any("err", err),
		)
	}
	return request.URL, err
}
