package logger

import (
	"context"
	"demo/pkg/chat/googlechat"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"testing"

	"github.com/samber/lo"
)

func TestLogger(t *testing.T) {
	chat := googlechat.New(googlechat.ChatConfig{
		WebhookURL: "https://chat.googleapis.com/v1/spaces/AAAAL4z5TVw/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=Cs4JhKd-RIuwTG2DIR_RUCLQu5XrgahkOKHOkMaiZHI",
		AppName:    "Local test app",
		Env:        "0.0.1",
		Version:    "1.0.0",
	})
	cf := LoggerConfig{
		Writer:       os.Stdout,
		LogFromLevel: slog.LevelInfo,
		ChatBot: &LoggerConfigChatBox{
			ChatbotByLevel: map[slog.Level]Chat{
				slog.LevelInfo:  chat,
				slog.LevelWarn:  chat,
				slog.LevelError: chat,
			},
			MsgTemplate: func(msg string, attrs []slog.Attr) string {
				logAttrs := lo.Map(attrs, func(a slog.Attr, _ int) string {
					return fmt.Sprintf("- *%s*: %v", a.Key, a.Value)
				})
				text := msg
				if len(logAttrs) > 0 {
					text = strings.Join(logAttrs, "\n") + "\n" + msg
				}
				return text
			},
		},
		ContextArgs: []string{"lang", "request-trace-id"},
		IsBeauty:    false,
	}
	InitLogger(cf)

	ctx := context.WithValue(context.WithValue(context.Background(), "lang", "en"), "request-trace-id", "liqwbfpqwifqpwfbpwiq")
	slog.InfoContext(ctx, "Log info test", slog.String("arg1", "arg1 value"))
	fmt.Println(123)
}
