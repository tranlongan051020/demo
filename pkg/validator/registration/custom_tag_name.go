package registration

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	TagNameI18n = "i18n"
	TagNameJSON = "json"
	TagNameForm = "form"
	TagNameUri  = "uri"
)

func RegisterCustomGetTagName(v *validator.Validate) {
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		i18nTag, ok := fld.Tag.Lookup(TagNameI18n)
		if ok {
			return i18nTag
		}
		jsonTag, ok := fld.Tag.Lookup(TagNameJSON)
		if ok {
			name := strings.SplitN(jsonTag, ",", 2)[0]
			return name
		}
		formTag, ok := fld.Tag.Lookup(TagNameForm)
		if ok {
			name := strings.SplitN(formTag, ",", 2)[0]
			return name
		}
		uriTag, ok := fld.Tag.Lookup(TagNameUri)
		if ok {
			name := strings.SplitN(uriTag, ",", 2)[0]
			return name
		}
		return fld.Name
	})
}
