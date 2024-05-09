package enum

import "demo/pkg/aerror"

var ErrorTagToCode = map[string]aerror.ErrorCode{
	TagRequired:       aerror.ErrValidateRequired,
	TagRequiredIf:     aerror.ErrValidateRequired,
	TagRequiredWith:   aerror.ErrValidateRequired,
	TagRequiredUnless: aerror.ErrValidateRequired,
	TagLen:            aerror.ErrValidateLen,
	TagMin:            aerror.ErrValidateMin,
	TagMax:            aerror.ErrValidateMax,
	TagEq:             aerror.ErrValidateEq,
	TagNe:             aerror.ErrValidateNe,
	TagLt:             aerror.ErrValidateLt,
	TagLte:            aerror.ErrValidateLte,
	TagGt:             aerror.ErrValidateGt,
	TagGte:            aerror.ErrValidateGte,
	TagEqField:        aerror.ErrValidateEqField,
	TagEqCsField:      aerror.ErrValidateEqCsField,
	TagNecSField:      aerror.ErrValidateNecSField,
	TagGtCsField:      aerror.ErrValidateGtCsField,
	TagGteCsField:     aerror.ErrValidateGteCsField,
	TagLtCsField:      aerror.ErrValidateLtCsField,
	TagLteCsField:     aerror.ErrValidateLteCsField,
	TagNeField:        aerror.ErrValidateNeField,
	TagGteField:       aerror.ErrValidateGteField,
	TagGtField:        aerror.ErrValidateGtField,
	TagLteField:       aerror.ErrValidateLteField,
	TagLtField:        aerror.ErrValidateLtField,
	TagAlpha:          aerror.ErrValidateAlpha,
	TagAlphanum:       aerror.ErrValidateAlphanum,
	TagBoolean:        aerror.ErrValidateBoolean,
	TagNumeric:        aerror.ErrValidateNumeric,
	TagNumber:         aerror.ErrValidateNumber,
	TagEmail:          aerror.ErrValidateEmail,
	TagDateTimeLayout: aerror.ErrValidateDateTimeLayout,
	TagDateTimeBefore: aerror.ErrValidateDateTimeBefore,
	TagDateTimeAfter:  aerror.ErrValidateDateTimeAfter,
	TagOneOf:          aerror.ErrValidateOneOf,
}