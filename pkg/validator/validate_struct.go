package validator

import (
	"context"

	"demo/pkg/aerror"
	"demo/pkg/validator/enum"

	"github.com/go-playground/validator/v10"
)

func (v *validation) ValidateStruct(ctx context.Context, s any) (errs []error) {
	err := v.validate.Struct(s)
	if err != nil {
		validateErrs := err.(validator.ValidationErrors)
		for _, e := range validateErrs {
			errCode := enum.ErrorTagToCode[e.ActualTag()]
			aErr := aerror.New(ctx, errCode, e.Param())
			errs = append(errs, (error)(aErr))
		}
	}
	return errs
}
