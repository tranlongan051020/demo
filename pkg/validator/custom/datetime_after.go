package custom

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateDateTimeAfter(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	param := fl.Param()

	var regexValidate string
	for _, regex := range LayoutRegexs {
		match, _ := regexp.MatchString(regex, param)
		if match {
			regexValidate = regex
			break
		}
	}
	if len(regexValidate) == 0 {
		return false
	}

	match, _ := regexp.MatchString(regexValidate, value)
	if !match {
		return false
	}

	result := strings.Compare(value, param)
	if result > 0 {
		return true
	}

	return false
}
