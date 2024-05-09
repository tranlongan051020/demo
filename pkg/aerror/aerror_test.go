package aerror

import (
	"context"
	"demo/constant"
	"fmt"
	"testing"
)

const (
	customSpecialErr ErrorCode = "CustomSpecialErr" // Custom special error with param {{.Param}}
)

type (
	customSpecialErrArgs struct {
		Param string
	}
)

func TestError(t *testing.T) {
	ctx := context.WithValue(context.Background(), constant.ContextKey.Language, "vi")

	// case error build in
	err1 := New(ctx, ErrValidateRequired, ErrRequiredArgs{
		FieldName: "name",
	})
	fmt.Println("Result error 1: ", err1.Message)

	// case error custom on special domain logic
	err2 := New(ctx, customSpecialErr, customSpecialErrArgs{
		Param: "email",
	})
	fmt.Println("Result error 2: ", err2.Message)
}
