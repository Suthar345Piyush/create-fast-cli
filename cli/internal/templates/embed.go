// the final create-fast-cli binary is self contained

package templates

import "embed"

//go:embed files
//go:embed *.tmpl
var FS embed.FS
