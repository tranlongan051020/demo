package logger

import (
	"context"
	"log/slog"

	slogmulti "github.com/samber/slog-multi"
)

func contextHandler(minLv slog.Level, contextArgs []string) slogmulti.Middleware {
	enabledFunc := func(ctx context.Context, level slog.Level, next func(context.Context, slog.Level) bool) bool {
		if minLv > level {
			return false
		}
		return next(ctx, level)
	}

	handleFunc := func(ctx context.Context, record slog.Record, next func(context.Context, slog.Record) error) error {
		if ctx == nil {
			return next(ctx, record)
		}

		for _, key := range contextArgs {
			if value := ctx.Value(key); value != nil {
				record.AddAttrs(slog.Any(key, value))
			}
		}

		return next(ctx, record)
	}

	withAttrsFunc := func(attrs []slog.Attr, next func([]slog.Attr) slog.Handler) slog.Handler {
		return next(attrs)
	}

	withGroupFunc := func(name string, next func(string) slog.Handler) slog.Handler {
		return next(name)
	}

	mw := slogmulti.NewInlineMiddleware(enabledFunc, handleFunc, withAttrsFunc, withGroupFunc)

	return mw
}
