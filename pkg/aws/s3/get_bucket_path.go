package s3

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
)

func (s *s3Client) GetBucketPath(ctx context.Context, s3URL string) (bucketPath string, err error) {
	err = fmt.Errorf("uri_format")

	if s3URL == "" {
		slog.ErrorContext(ctx, "GetBucketPath error #1",
			slog.String("value", s3URL),
			slog.Any("err", err),
		)
		return "", err
	}

	httpIdx := strings.Index(s3URL, "//")
	if httpIdx < 0 {
		slog.ErrorContext(ctx, "GetBucketPath error #2 ",
			slog.String("value", s3URL),
			slog.Any("err", err),
		)
		return "", err
	}

	domainURL := s3URL[httpIdx+2:]
	domainIdx := strings.Index(domainURL, "/")

	if domainIdx < 0 {
		slog.ErrorContext(ctx, "GetBucketPath error #3 ",
			slog.String("value", s3URL),
			slog.Any("err", err),
		)
		return "", err
	}

	return domainURL[domainIdx+1:], nil
}
