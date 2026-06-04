// the final create-fast-cli binary is self contained

package templates

import "embed"

//go:embed files/*
//go:embed files/**/*
var FS embed.FS
