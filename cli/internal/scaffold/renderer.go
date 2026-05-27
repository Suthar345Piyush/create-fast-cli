// renderer - handles the final content for single output file

package scaffold

import (
	"bytes"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/config"
	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/templates"
	"github.com/Suthar345Piyush/create-fast-cli/cli/pkg/utils"
)

type RenderedFile struct {
	RelPath string // relative path to the project
	Content []byte // content of the file
}

// function render will render the common and typed file content in the respective file
// this will return RenderedFile as output and take project config as input parameter

func Render(cfg *config.ProjectConfig) ([]RenderedFile, error) {

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

	// iterate on the common and typed file system

	for _, sub := range []fs.FS{commonFS, typedFS} {

		rendered, err := renderFS(sub, cfg)
		if err != nil {
			return nil, err
		}

		for _, rf := range rendered {
			seen[rf.RelPath] = true
			files = append(files, rf)
		}

	}

	_ = seen

	return files, nil

}

// render func will renders all the template files

func renderFS(fsys fs.FS, cfg *config.ProjectConfig) ([]RenderedFile, error) {

	var results []RenderedFile

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".tmpl") {
			return nil
		}

		raw, err := fs.ReadFile(fsys, path)

		if err != nil {
			return fmt.Errorf("read template %q: %w", path, err)
		}

		tmpl, err := template.New(path).Funcs(templateFuncs()).Parse(string(raw))

		if err != nil {
			return fmt.Errorf("execute template %q: %w", path, err)
		}

		// executing those, who are in buffer state

		var buf bytes.Buffer

		if err := tmpl.Execute(&buf, cfg); err != nil {
			return fmt.Errorf("execute template %q: %w", path, err)
		}

		// removing .tmpl suffix

		outPath := strings.TrimSuffix(path, ".tmpl")

		// removing all "/" from the output path

		outPath = filepath.FromSlash(outPath)

		//going to return the result

		results = append(results, RenderedFile{
			RelPath: outPath,
			Content: buf.Bytes(),
		})
		return nil
	})

	if err != nil {
		return nil, err
	}

	// at last return the results and nil

	return results, nil

}

// templateFuncs() function

func templateFuncs() template.FuncMap {

	// it will return a func map
	// using this we can map "names" -> "functions"
	// example - title -> PascalCase

	return template.FuncMap{

		"title": utils.PascalCase,
		"slug":  utils.SlugFunc,
		"hasFeature": func(cfg *config.ProjectConfig, feature string) bool {
			switch feature {
			case "UseTUI":
				return cfg.UseTUI
			case "UseLogging":
				return cfg.UseLogging
			case "UseConfig":
				return cfg.UseConfig
			case "UseCompletions":
				return cfg.UseCompletions
			case "UseTesting":
				return cfg.UseTesting
			}
			return false
		},

		// if chosen framework is cobra them return true

		"isCobra": func(cfg *config.ProjectConfig) bool {
			return cfg.Framework == config.FrameworkCobra
		},
	}

}
