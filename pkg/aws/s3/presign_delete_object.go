package s3

import (
	"context"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *s3Client) PresignDeleteObject(
	ctx context.Context,
	s3ObjectKey string,
	lifetimeSecs int64,
) (deleteObjectURL string, err error) {
	//nolint:exhaustruct // s3 struct, not our struct
	request, err := s.presignClient.PresignDeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(s3ObjectKey),
	})
	if err != nil {
		slog.ErrorContext(ctx, "PresignDeleteObject error #1",
			slog.String("s3ObjectKey", s3ObjectKey),
			slog.Any("err", err),
		)
	}
	return request.URL, err
}
