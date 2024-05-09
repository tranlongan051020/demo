package s3

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
)

func (s *s3Client) GetFileName(ctx context.Context, s3URL string) (fileName string, err error) {
	err = fmt.Errorf("uri_format")

	if s3URL == "" {
		slog.ErrorContext(ctx, "GetFileName error #1",
			slog.String("value", s3URL),
			slog.Any("err", err),
		)
		return "", err
	}

	parts := strings.Split(s3URL, "/")

	if len(parts) < 1 {
		slog.ErrorContext(ctx, "GetFileName error #2",
			slog.String("value", s3URL),
			slog.Any("err", err),
		)
		return "", err
	}

	return parts[len(parts)-1], nil
}
