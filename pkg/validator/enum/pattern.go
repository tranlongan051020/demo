package enum

const (
	RegexDDMMYYYYSlash          = `^(\d{2})/(\d{2})/(\d{4})$`                                  // 19/04/2023
	RegexYYYYMMDDHHMMSSSlash    = `^(\d{4})/(\d{2})/(\d{2})\s(\d{2}):(\d{2}):(\d{2})$`         // 2023/04/19 15:30:45
	RegexYYYYMMDDHHMMSSXXXSlash = `^(\d{4})/(\d{2})/(\d{2})\s(\d{2}):(\d{2}):(\d{2}).(\d{3})$` // 2023/04/19 15:30:45.999
)
