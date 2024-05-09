package registration

import (
	"demo/pkg/validator/custom"
	"demo/pkg/validator/enum"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidations(v *validator.Validate) {
	v.RegisterValidation(enum.TagDateTimeAfter, custom.ValidateDateTimeAfter)
	v.RegisterValidation(enum.TagDateTimeLayout, custom.ValidateDateTimeLayout)
	v.RegisterValidation(enum.TagDateTimeBefore, custom.ValidateDateTimeBefore)
}
