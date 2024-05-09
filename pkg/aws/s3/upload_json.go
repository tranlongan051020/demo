package s3

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/goccy/go-json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (s *s3Client) UploadJSON(
	ctx context.Context,
	jsonableStruct any,
) (s3Location string, err error) {

	// create uuid
	nanoid, _ := gonanoid.New()
	// now
	now := time.Now()
	unixFormat := now.Unix()
	uniqid := fmt.Sprintf("%d-%s", unixFormat, nanoid)

	jsonBytes, err := json.Marshal(jsonableStruct)
	if err != nil {
		slog.ErrorContext(ctx, "UploadJSON error #1",
			slog.Any("err", err),
		)
		return "", err
	}

	// upload S3 and then return URL
	jsonBytesReader := bytes.NewReader(jsonBytes)
	uploader := manager.NewUploader(s.client)
	//nolint:exhaustruct // s3 struct, not our struct
	input := s3.PutObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(uniqid + ".json"),
		Body:   jsonBytesReader,
	}

	output, err := uploader.Upload(ctx, &input)
	if err != nil {
		// Print the error and exit.
		slog.ErrorContext(ctx, "UploadJSON error #2",
			slog.String("Bucket", s.config.Bucket),
			slog.Any("err", err),
		)
		return "", err
	}

	return output.Location, nil
}
