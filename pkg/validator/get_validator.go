package validator

import (
	"context"
	"sync"

	"demo/pkg/validator/registration"

	"github.com/go-playground/validator/v10"
)

var (
	once     sync.Once
	instance *validation
)

type Validation interface {
	ValidateStruct(ctx context.Context, s any) (errs []error)
}

type validation struct {
	validate *validator.Validate
}

func GetValidator() Validation {
	once.Do(func() {
		v := validator.New()
		registration.RegisterCustomValidations(v)
		registration.RegisterCustomGetTagName(v)
		instance = new(validation)
		instance.validate = v
	})

	return instance
}
