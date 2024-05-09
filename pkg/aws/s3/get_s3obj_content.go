package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *s3Client) GetS3ObjContentT1(
	ctx context.Context,
	s3Path string,
	fileName string,
) (contentBuf io.Reader, err error) {

	if s3Path == "" || fileName == "" {
		return nil, fmt.Errorf("invalid s3 path or filename")
	}

	s3Key := fmt.Sprintf("%s/%s", s3Path, fileName)
	return s.GetS3ObjContentT2(ctx, s3Key)

}

func (s *s3Client) GetS3ObjContentT2(
	ctx context.Context,
	s3Key string,
) (contentBuf io.Reader, err error) {

	if s3Key == "" {
		return nil, fmt.Errorf("invalid s3 key")
	}

	downloader := manager.NewDownloader(s.client)
	//nolint:exhaustruct // s3 struct, not our struct
	input := s3.GetObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(s3Key),
	}
	buffer := manager.NewWriteAtBuffer([]byte{})
	_, err = downloader.Download(ctx, buffer, &input)
	if err != nil {
		// Print the error and exit.
		slog.ErrorContext(ctx, "Get object error #1",
			slog.String("Bucket", s.config.Bucket),
			slog.String("s3Key", s3Key),
			slog.Any("err", err),
		)
		return nil, err
	}
	contentBuf = bytes.NewBuffer(buffer.Bytes())
	return contentBuf, nil
}
