// the final create-fast-cli binary is self contained

package templates

import "embed"

//go:embed all:files
var FS embed.FS
