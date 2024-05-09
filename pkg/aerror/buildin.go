package aerror

// build-in error list for app
const (
	// validate error
	ErrValidateRequired       ErrorCode = "ErrValidateRequired"
	ErrValidateLen            ErrorCode = "ErrValidateLen"
	ErrValidateMin            ErrorCode = "ErrValidateMin"
	ErrValidateMax            ErrorCode = "ErrValidateMax"
	ErrValidateEq             ErrorCode = "ErrValidateEq"
	ErrValidateNe             ErrorCode = "ErrValidateNe"
	ErrValidateLt             ErrorCode = "ErrValidateLt"
	ErrValidateLte            ErrorCode = "ErrValidateLte"
	ErrValidateGt             ErrorCode = "ErrValidateGt"
	ErrValidateGte            ErrorCode = "ErrValidateGte"
	ErrValidateEqField        ErrorCode = "ErrValidateEqField"
	ErrValidateEqCsField      ErrorCode = "ErrValidateEqCsField"
	ErrValidateNecSField      ErrorCode = "ErrValidateNecSField"
	ErrValidateGtCsField      ErrorCode = "ErrValidateGtCsField"
	ErrValidateGteCsField     ErrorCode = "ErrValidateGteCsField"
	ErrValidateLtCsField      ErrorCode = "ErrValidateLtCsField"
	ErrValidateLteCsField     ErrorCode = "ErrValidateLteCsField"
	ErrValidateNeField        ErrorCode = "ErrValidateNeField"
	ErrValidateGteField       ErrorCode = "ErrValidateGteField"
	ErrValidateGtField        ErrorCode = "ErrValidateGtField"
	ErrValidateLteField       ErrorCode = "ErrValidateLteField"
	ErrValidateLtField        ErrorCode = "ErrValidateLtField"
	ErrValidateAlpha          ErrorCode = "ErrValidateAlpha"
	ErrValidateAlphanum       ErrorCode = "ErrValidateAlphanum"
	ErrValidateBoolean        ErrorCode = "ErrValidateBoolean"
	ErrValidateNumeric        ErrorCode = "ErrValidateNumeric"
	ErrValidateNumber         ErrorCode = "ErrValidateNumber"
	ErrValidateEmail          ErrorCode = "ErrValidateEmail"
	ErrValidateDateTimeLayout ErrorCode = "ErrValidateDateTimeLayout"
	ErrValidateDateTimeBefore ErrorCode = "ErrValidateDateTimeBefore"
	ErrValidateDateTimeAfter  ErrorCode = "ErrValidateDateTimeAfter"
	ErrValidateOneOf          ErrorCode = "ErrValidateOneOf"

	// Unexpected error
	ErrUnexpected ErrorCode = "ErrUnexpected"
)

type (
	ErrRequiredArgs = struct {
		FieldName string
	}
)
