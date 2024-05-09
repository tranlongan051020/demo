package s3

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *s3Client) UploadFolderZip(
	ctx context.Context,
	folderName string,
	localPath string,
	s3Path string,
) (s3URL string,
	err error) {

	filePath := localPath + "/" + folderName + ".zip"
	file, err := os.Open(filePath)
	if err != nil {
		return s3URL, fmt.Errorf("open file [%w]", err)
	}

	defer file.Close()

	// upload S3 and then return URL
	uploader := manager.NewUploader(s.client)

	//nolint:exhaustruct // s3 struct, not our struct
	input := s3.PutObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(s3Path + "/" + folderName + ".zip"),
		Body:   file,
	}

	output, err := uploader.Upload(ctx, &input)
	if err != nil {
		// Print the error and exit.
		slog.ErrorContext(ctx, "UploadFolderZip error #1",
			slog.String("Bucket", s.config.Bucket),
			slog.String("filePath", filePath),
			slog.Any("err", err),
		)
		return s3URL, err
	}

	if output == nil || output.Location == "" {
		return s3URL, fmt.Errorf("put_object")
	}

	suffixIdx := strings.LastIndex(output.Location, "/")
	if suffixIdx < 0 {
		return s3URL, fmt.Errorf("put_object")
	}
	fileNameInURL := output.Location[suffixIdx+1:]
	s3URL = strings.Replace(output.Location, fileNameInURL, folderName+".zip", 1)

	return s3URL, nil
}
