// renderer - handles the final content for single output file

package scaffold

import (
	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/config"
	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/templates"
)

type RenderedFile struct {
	RelPath string // relative path to the project
	Content []byte // content of the file
}

// function render will render the common and typed file content in the respective file
// this will return RenderedFile as output and take project config as input parameter

func Render(cfg *config.ProjectConfig) (*RenderedFile, error) {

	// common and typed file system

	commonFS, typedFS, err := templates.SubFS(cfg.AppType, cfg.Framework)

	if err != nil {
		return nil, err
	}

	//slice of rendered files

	var files []RenderedFile

	//common templates will showed first, then typed template
	// using map - key:value pair
	// key - string , value - bool

	seen := map[string]bool{}

}
