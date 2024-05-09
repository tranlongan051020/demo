package logger

import (
	"io"
	"log/slog"
	"time"

	"github.com/lmittmann/tint"
	slogmulti "github.com/samber/slog-multi"
)

type Chat interface {
	SendText(text string) error
}

type LoggerConfig struct {
	Writer       io.Writer
	LogFromLevel slog.Level
	IsBeauty     bool
	ContextArgs  []string
	ChatBot      *LoggerConfigChatBox
}

type LoggerConfigChatBox struct {
	ChatbotByLevel map[slog.Level]Chat
	MsgTemplate    func(msg string, attrs []slog.Attr) string
}

func InitLogger(cf LoggerConfig) {
	var logHandler slog.Handler
	if cf.IsBeauty {
		logHandler = tint.NewHandler(cf.Writer, &tint.Options{
			Level:      cf.LogFromLevel,
			TimeFormat: time.RFC1123,
			AddSource:  true,
		})
	} else {
		logHandler = slog.NewJSONHandler(cf.Writer, &slog.HandlerOptions{
			Level: cf.LogFromLevel,
		})
	}

	if cf.ChatBot != nil {
		logHandler = slogmulti.Fanout(logHandler, newChatHandler(chatHandleConfig{
			LogContextKeys: cf.ContextArgs,
			Chatbot:        cf.ChatBot.ChatbotByLevel,
			MsgTemplate:    cf.ChatBot.MsgTemplate,
		}))
	}

	logHandler = slogmulti.
		Pipe(contextHandler(cf.LogFromLevel, cf.ContextArgs)).
		Handler(logHandler)

	slog.SetDefault(slog.New(logHandler))

}
