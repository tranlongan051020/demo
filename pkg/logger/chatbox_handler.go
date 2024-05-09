package logger

import (
	"context"
	"log/slog"

	slogcommon "github.com/samber/slog-common"
)

type chatHandleConfig struct {
	LogContextKeys []string
	Chatbot        map[slog.Level]Chat
	MsgTemplate    func(msg string, attrs []slog.Attr) string
}

type chatHandler struct {
	config chatHandleConfig
	attrs  []slog.Attr
	groups []string
}

func newChatHandler(cf chatHandleConfig) slog.Handler {
	return &chatHandler{
		config: cf,
	}
}

func (h *chatHandler) Enabled(ctx context.Context, lv slog.Level) bool {
	_, ok := h.config.Chatbot[lv]
	return ok
}

func (h *chatHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, key := range h.config.LogContextKeys {
		if value := ctx.Value(key); value != nil {
			h.attrs = append(h.attrs, slog.Any(key, value))
		}
	}

	text := h.config.MsgTemplate(record.Message, h.attrs)

	h.config.Chatbot[record.Level].SendText(text)
	return nil
}

func (h *chatHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.attrs = slogcommon.AppendAttrsToGroup(h.groups, h.attrs, attrs...)
	return h
}

func (h *chatHandler) WithGroup(name string) slog.Handler {
	h.groups = append(h.groups, name)
	return h
}
