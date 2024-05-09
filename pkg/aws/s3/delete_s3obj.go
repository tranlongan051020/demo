package s3

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *s3Client) DeleteS3ObjT1(
	ctx context.Context,
	s3Path string,
	fileName string,
) (err error) {

	if s3Path == "" || fileName == "" {
		return fmt.Errorf("invalid s3 path or filename")
	}

	s3Key := fmt.Sprintf("%s/%s", s3Path, fileName)
	return s.DeleteS3ObjT2(ctx, s3Key)
}

func (s *s3Client) DeleteS3ObjT2(
	ctx context.Context,
	s3Key string,
) (err error) {
	if s3Key == "" {
		return fmt.Errorf("invalid s3 key")
	}

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(s3Key),
	}

	_, err = s.client.DeleteObject(ctx, input)
	if err != nil {
		// Print the error and exit.
		slog.ErrorContext(ctx, "Delete object error #1",
			slog.String("Bucket", s.config.Bucket),
			slog.String("s3Key", s3Key),
			slog.Any("err", err),
		)
		return err
	}
	return nil
}
