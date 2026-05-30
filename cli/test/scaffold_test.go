// testing for scaffold section of project, using testify checks for every thing to be available, if not returns error

package test

import (
	"path/filepath"
	"testing"

	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/config"
	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/scaffold"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// default project config gives a full project config

func defaultCfg() *config.ProjectConfig {
	return &config.ProjectConfig{
		ProjectName:    "test-tool",
		ModulePath:     "github.com/test/test-tool",
		AppType:        config.AppTypeDevTool,
		Framework:      config.FrameworkCobra,
		GoVersion:      "1.26",
		UseTUI:         true,
		UseLogging:     true,
		UseConfig:      true,
		UseCompletions: true,
		UseTesting:     true,
		IDE:            config.IDENone,
		OutputDir:      "/tmp/test-tool",
	}
}

// build index function maps, []RenderedFile to map (relativePath(relPath) -> content)

func buildIndex(files []scaffold.RenderedFile) map[string]string {
	idx := make(map[string]string, len(files))

	for _, f := range files {
		idx[filepath.ToSlash(f.RelPath)] = string(f.Content)
	}

	return idx

}

// all tests functions related to rendering

// return files test function, which verify render returns an list of default devtool cobra configuration - more alike default project template with default settings

func TestRender_ReturnFiles(t *testing.T) {
	files, err := scaffold.Render(defaultCfg())

	require.NoError(t, err, "Render should not return an error for a valid config")
	assert.NotEmpty(t, files, "Render should return at least one file")

}

// test function to check their must be file are present while rendering happens

func TestRender_RequiredFilesPresent(t *testing.T) {

	files, err := scaffold.Render(defaultCfg())
	require.NoError(t, err)

	// index

	idx := buildIndex(files)

	// required slice with file names

	required := []string{
		"main.go",
		"go.mod",
		"cmd/root.go",
		"cmd/run.go",
		"README.md",
		"Makefile",
		".gitignore",
	}

	// iterate on the required, to check file path is present or not

	for _, path := range required {
		assert.Contains(t, idx, path, "required file %q is missing from render output", path)
	}
}

// test to check the module path

func TestRender_ModulePathPresent(t *testing.T) {
	cfg := defaultCfg()

	cfg.ModulePath = "github.com/myorg/super-tool"

	files, err := scaffold.Render(cfg)
	require.NoError(t, err)

	idx := buildIndex(files)

	gomod, ok := idx["go.mod"]
	require.True(t, ok, "go.mod must be present in render output")

	assert.Contains(t, gomod, "github.com/myorg/super-tool", "go.mod should contain the module path")

}
