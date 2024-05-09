package aerror

import (
	"bytes"
	"context"
	"html/template"
	"net/http"

	"demo/constant"
	"demo/resource/localization"

	"gopkg.in/yaml.v2"
)

const (
	defauldemonguage = "en"
	defaultFileExt   = ".yaml"
)

type ErrorCode string

type AError struct {
	Code    string
	Message string

	HTTPStatusCode int
	Detail         any
}

func (e AError) Error() string {
	return e.Message
}

func New(ctx context.Context, code ErrorCode, args ...any) *AError {
	var lang = defauldemonguage
	langCtx := ctx.Value(constant.ContextKey.Language)
	if langCtx != nil {
		lang = ctx.Value(constant.ContextKey.Language).(string)
	}

	message := getMessageFromYAML(string(code), lang)

	if len(args) > 0 {
		message = formatMessage(message, args[0])
	}

	return &AError{
		Code:           string(code),
		Message:        message,
		HTTPStatusCode: http.StatusInternalServerError, // default error get status 500
	}
}

func formatMessage(message string, params any) string {
	tpl, _ := template.New("msg").Parse(message)
	var tplBuffer bytes.Buffer
	tpl.Execute(&tplBuffer, params)
	return tplBuffer.String()
}

func getMessageFromYAML(code, lang string) string {
	fileName := lang + defaultFileExt // only use yaml file format
	fileInfos, err := localization.LocalizationFiles.ReadDir(".")
	if err != nil {
		return ""
	}

	fileExist := false
	for _, fileInfo := range fileInfos {
		if fileInfo.Name() == fileName {
			fileExist = true
			break
		}
	}

	if !fileExist {
		if len(fileInfos) > 0 {
			fileName = fileInfos[0].Name()
		} else {
			return ""
		}
	}

	data, err := localization.LocalizationFiles.ReadFile(fileName)
	if err != nil {
		return ""
	}

	var messages map[string]string
	err = yaml.Unmarshal(data, &messages)
	if err != nil {
		return ""
	}

	message, found := messages[code]
	if !found {
		return ""
	}

	return message
}
