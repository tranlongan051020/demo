package localization

import "embed"

//go:embed *.yaml
var LocalizationFiles embed.FS
