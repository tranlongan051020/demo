package enum

var (
	// Build-in tags
	TagRequired        = "required"        // Validates that a value is present and not empty
	TagRequiredIf      = "required_if"     //
	TagRequiredWith    = "required_with"   //
	TagRequiredUnless  = "required_unless" //
	TagLen             = "len"             // Validates the length of a value
	TagMin             = "min"             // Validates that the len of value is greater than or equal to a specified minimum
	TagMax             = "max"             // Validates that the len of value is less than or equal to a specified maximum
	TagEq              = "eq"              // Validates that a value is equal to a specified value
	TagNe              = "ne"              // Validates that a value is not equal to a specified value
	TagLt              = "lt"              // Validates that a value is less than a specified value
	TagLte             = "lte"             // Validates that a value is less than or equal to a specified value
	TagGt              = "gt"              // Validates that a value is greater than a specified value
	TagGte             = "gte"             // Validates that a value is greater than or equal to a specified value
	TagEqField         = "eqfield"         // Validates that a value is equal to the value of another field
	TagEqCsField       = "eqcsfield"       // Validates that a value is equal to the value of another case-sensitive field
	TagNecSField       = "necsfield"       // Validates that a value is not equal to the value of another case-sensitive field
	TagGtCsField       = "gtcsfield"       // Validates that a value is greater than the value of another case-sensitive field
	TagGteCsField      = "gtecsfield"      // Validates that a value is greater than or equal to the value of another case-sensitive field
	TagLtCsField       = "ltcsfield"       // Validates that a value is less than the value of another case-sensitive field
	TagLteCsField      = "ltecsfield"      // Validates that a value is less than or equal to the value of another case-sensitive field
	TagNeField         = "nefield"         // Validates that a value is not equal to the value of another field
	TagGteField        = "gtefield"        // Validates that a value is greater than or equal to the value of another field
	TagGtField         = "gtfield"         // Validates that a value is greater than the value of another field
	TagLteField        = "ltefield"        // Validates that a value is less than or equal to the value of another field
	TagLtField         = "ltfield"         // Validates that a value is less than the value of another field
	TagAlpha           = "alpha"           // Validates that a value contains only alphabetic characters
	TagAlphanum        = "alphanum"        // Validates that a value contains only alphanumeric characters
	TagAlphaUnicode    = "alphaunicode"    // Validates that a value contains only Unicode alphabetic characters
	TagAlphanumUnicode = "alphanumunicode" // Validates that a value contains only Unicode alphanumeric characters
	TagBoolean         = "boolean"         // Validates that a value is a boolean
	TagNumeric         = "numeric"         // Validates that a value is a numeric value
	TagNumber          = "number"          // Validates that a value is a numeric value
	TagEmail           = "email"           // Validates that a value is a valid email address
	TagLowercase       = "lowercase"       // Validates that a value contains only lowercase characters
	TagUppercase       = "uppercase"       // Validates that a value contains only uppercase characters
	TagOneOf           = "oneof"           //

	// Custom tags
	TagFullSize       = "fullsize"        // Full size format japanese
	TagHalfSize       = "halfsize"        // Half size format japanese
	TagDateTimeLayout = "datetime_layout" // Date time layout format
	TagDateTimeBefore = "datetime_before" // Date time before format
	TagDateTimeAfter  = "datetime_after"  // Date time after format
)
