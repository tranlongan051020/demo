package s3

import (
	"context"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *s3Client) Download(
	ctx context.Context,
	inputURLs []string,
) (
	mapURLFilepath map[string]string, err error) {

	tmpDir := os.TempDir()

	downloader := manager.NewDownloader(s.client)
	downloader.Concurrency = len(inputURLs)

	mapURLFilepath = make(map[string]string, 0)
	for _, URL := range inputURLs {
		name, _ := s.GetFileName(ctx, URL)
		awsPath, _ := s.GetBucketPath(ctx, URL)
		downloadFile := tmpDir + name

		fileDownload, err1 := os.Create(downloadFile)
		if err1 != nil {
			slog.ErrorContext(ctx, "Download error #1",
				slog.String("URL", URL),
				slog.String("downloadFile", downloadFile),
				slog.Any("err", err),
			)
			return nil, err1
		}

		defer fileDownload.Close()
		mapURLFilepath[URL] = downloadFile

		//nolint:exhaustruct // s3 struct, not our struct
		_, err = downloader.Download(ctx, fileDownload, &s3.GetObjectInput{
			Bucket: aws.String(s.config.Bucket),
			Key:    aws.String(awsPath),
		})
		if err != nil {
			slog.ErrorContext(ctx, "Download error #2",
				slog.String("URL", URL),
				slog.String("downloadFile", downloadFile),
				slog.Any("err", err),
			)
			return nil, err
		}
	}
	return mapURLFilepath, nil
}
