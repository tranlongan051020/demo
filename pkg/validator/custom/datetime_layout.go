package custom

import (
	"regexp"

	"demo/pkg/validator/enum"

	"github.com/go-playground/validator/v10"
)

// Map from layout to regex
var LayoutRegexs = []string{
	enum.RegexDDMMYYYYSlash,
	enum.RegexYYYYMMDDHHMMSSSlash,
	enum.RegexYYYYMMDDHHMMSSXXXSlash,
	enum.RegexYYYYMMDDHHMMSSXXXSlash,
}

func ValidateDateTimeLayout(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	layoutParam := fl.Param()

	if len(value) == 0 {
		return true
	}

	var regexValidate string
	for _, regex := range LayoutRegexs {
		match, _ := regexp.MatchString(regex, layoutParam)
		if match {
			regexValidate = regex
			break
		}
	}
	if len(regexValidate) == 0 {
		return false
	}

	match, _ := regexp.MatchString(regexValidate, value)
	return match
}
