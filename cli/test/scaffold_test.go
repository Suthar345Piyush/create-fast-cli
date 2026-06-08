// testing for scaffold section of project, using testify checks for every thing to be available, if not returns error

package test

import (
	"os"
	"path/filepath"
	"strings"
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

// test function for checking project name

func TestRender_ProjectNamePresent(t *testing.T) {

	cfg := defaultCfg()

	cfg.ProjectName = "my-cool-cli"

	cfg.ModulePath = "github.com/myorg/cool-tool"

	files, err := scaffold.Render(cfg)

	require.NoError(t, err)

	idx := buildIndex(files)

	readme, ok := idx["README.md"]
	require.True(t, ok, "README.md should be present in file")
	assert.Contains(t, readme, "my-cool-cli", "README.md should reference the project name")

	rootGo, ok := idx["cmd/root.go"]
	require.True(t, ok, "cmd/root.go must be present")
	assert.Contains(t, rootGo, "my-cool-cli", "cmd/root.go should reference the project name in the Use field")

}

// test function for checking go version

func TestRender_GoVersionPresent(t *testing.T) {
	cfg := defaultCfg()

	cfg.GoVersion = "1.26"

	files, err := scaffold.Render(cfg)

	require.NoError(t, err)

	idx := buildIndex(files)
	goMod, ok := idx["go.mod"]
	require.True(t, ok)
	assert.Contains(t, goMod, "1.26", "go.mod should declare the configured Go version")

}

// test function for checking logging is enabled or not

func TestRender_LoggingEnbled(t *testing.T) {

	cfg := defaultCfg()

	cfg.UseLogging = true

	files, err := scaffold.Render(cfg)
	require.NoError(t, err)

	idx := buildIndex(files)

	goMod := idx["go.mod"]
	assert.Contains(t, goMod, "go.uber.org/zap", "go.mod should declare zap when UseLogging=true")

}

// test function for logging disabled

func TestRender_LoggingDisabled(t *testing.T) {

	cfg := defaultCfg()
	cfg.UseLogging = false

	files, err := scaffold.Render(cfg)

	require.NoError(t, err)

	idx := buildIndex(files)

	goMod := idx["go.mod"]
	assert.NotContains(t, goMod, "go.uber.org/zap", "go.mod should NOT import zap when UseLogging=false")

}

// test function for config enabled - viper

func TestRender_ConfigEnabled(t *testing.T) {

	cfg := defaultCfg()
	cfg.UseConfig = true

	files, err := scaffold.Render(cfg)
	require.NoError(t, err)

	idx := buildIndex(files)

	goMod := idx["go.mod"]
	assert.Contains(t, goMod, "github.com/spf13/viper", "go.mod declare vuper when UseConfig=true")

}

// config disabled

func TestRender_ConfigDisabled(t *testing.T) {
	cfg := defaultCfg()
	cfg.UseConfig = false

	files, err := scaffold.Render(cfg)

	require.NoError(t, err)

	idx := buildIndex(files)

	goMod := idx["go.mod"]
	assert.NotContains(t, goMod, "github.com/spf13/viper", "go.mod should NOT import viper when UseConfig=false")

}

// test function for TUI enabled

func TestRender_TUIEnabled(t *testing.T) {

	cfg := defaultCfg()
	cfg.UseTUI = true

	files, err := scaffold.Render(cfg)

	require.NoError(t, err)

	idx := buildIndex(files)

	goMod := idx["go.mod"]
	assert.Contains(t, goMod, "charm.land/bubbletea/v2", "go.mod should declare bubbletea with UseTUI=true")

}

// disabled function for TUI

func TestRender_TUIDisabled(t *testing.T) {
	cfg := defaultCfg()
	cfg.UseTUI = false

	files, err := scaffold.Render(cfg)

	require.NoError(t, err)

	idx := buildIndex(files)

	goMod := idx["go.mod"]
	assert.NotContains(t, goMod, "charm.land/bubbletea/v2", "go.mod should not declare bubbletea with UseTUI=false")
}

// test function for testing (testify package) is enabled

func TestRender_TestingEnabled(t *testing.T) {

	cfg := defaultCfg()

	cfg.UseTesting = true

	files, err := scaffold.Render(cfg)
	require.NoError(t, err)

	idx := buildIndex(files)

	goMod := idx["go.mod"]
	assert.Contains(t, goMod, "github.com/strechr/testify v1.11.1", "go.mod should declare testify testing with UseTesting=true")

}

// testing disabled

func TestRender_TestingDisabled(t *testing.T) {

	cfg := defaultCfg()

	cfg.UseTesting = false

	files, err := scaffold.Render(cfg)
	require.NoError(t, err)

	idx := buildIndex(files)

	goMod := idx["go.mod"]
	assert.NotContains(t, goMod, "github.com/strechr/testify v1.11.1", "go.mod should declare testify testing with UseTesting=false")

}

// test function to check all features are disabled , while still it produces a file with required configs

func TestRender_AllFeaturesDisabled(t *testing.T) {

	cfg := defaultCfg()
	cfg.UseTUI = false
	cfg.UseConfig = false
	cfg.UseLogging = false
	cfg.UseCompletions = false
	cfg.UseTesting = false

	files, err := scaffold.Render(cfg)
	require.NoError(t, err)

	assert.NotEmpty(t, files, "Render should produce files even when all features are disabled")
	idx := buildIndex(files)

	for _, path := range []string{"go.mod", "main.go", "cmd/root.go"} {

		assert.Contains(t, idx, path, "core file %q must exist even with all features are false/off", path)

	}

}

// test function for checking rendering removed the template file suffix (tmpl) or not

func TestRender_RemoveSuffixTMPL(t *testing.T) {

	files, err := scaffold.Render(defaultCfg())
	require.NoError(t, err)

	for _, f := range files {
		assert.False(t, strings.HasSuffix(f.RelPath, ".tmpl"), "rendered file %q still contains .tmpl suffix", f.RelPath)
	}

}

// test function to check the files contain some content in them, they should not be completely empty
func TestRender_ContentCheck(t *testing.T) {
	files, err := scaffold.Render(defaultCfg())

	require.NoError(t, err)

	for _, i := range files {
		assert.NotEmpty(t, i.Content, "rendered file %q does not contain any content in them", i.Content)
	}

}

// test function to check all app types must to be present

func TestRender_AllAppTypePresent(t *testing.T) {
	appTypes := []config.AppType{
		config.AppTypeDevTool,
		config.AppTypeGitClient,
		config.AppTypeAiAssistant,
		config.AppTypeFileExplorer,
		config.AppTypeK8sTool,
		config.AppTypePackageManager,
		config.AppTypeSystemMonitor,
	}

	// iterate on the app types

	for _, aptyp := range appTypes {
		aptyp := aptyp

		t.Run(string(aptyp), func(t *testing.T) {

			cfg := defaultCfg()
			cfg.AppType = aptyp

			files, err := scaffold.Render(cfg)
			require.NoError(t, err, "Render should not error for app type %q", aptyp)
			assert.NotEmpty(t, files, "Render should return files for app type %q", aptyp)

			idx := buildIndex(files)

			assert.Contains(t, idx, "main.go", "main.go must exist for app type %q", aptyp)

			assert.Contains(t, idx, "go.mod", "go.mod must exist for app type %q", aptyp)

		})
	}

}

// test function for invalid app type

func TestRender_InvalidAppType(t *testing.T) {
	cfg := defaultCfg()

	cfg.AppType = config.AppType("does-not-exist")

	_, err := scaffold.Render(cfg)
	assert.Error(t, err, "Render should return an error for an unknown app type")

}

// test function to check cobra framework is present or not in default creation

func TestRender_CobraFramework(t *testing.T) {
	cfg := defaultCfg()

	cfg.Framework = config.FrameworkCobra

	files, err := scaffold.Render(cfg)

	require.NoError(t, err)

	idx := buildIndex(files)

	goMod := idx["go.mod"]
	assert.Contains(t, goMod, "github.com/spf13/cobra", "go.mod should declare cobra when framework=cobra")

}

// test function to check urfavecli

func TestRender_UrFaveCli(t *testing.T) {
	cfg := defaultCfg()
	cfg.Framework = config.FrameworkUrfaveCli

	files, err := scaffold.Render(cfg)
	require.NoError(t, err)

	idx := buildIndex(files)

	goMod := idx["go.mod"]
	assert.Contains(t, goMod, "github.com/urfave/cli", "go.mod should declare urfave when framework=urfave")
	assert.NotContains(t, goMod, "github.com/spf13/cobra", "go.mod should NOT declare cobra when framework=urfave")

}

// test function to check readme file uses pascalcase format to write the project name

func TestRender_PascalCaseInReadme(t *testing.T) {
	cfg := defaultCfg()
	cfg.ProjectName = "my-cool-cli"
	cfg.ModulePath = "github.com/x/my-cool-tool"

	files, err := scaffold.Render(cfg)
	require.NoError(t, err)

	idx := buildIndex(files)
	readme, ok := idx["README.md"]
	require.True(t, ok, "README.md must exist")
	require.Contains(t, readme, "MyCoolTool", "README.md title should be in pascalcase")

}

// test function to check the  git-client status command, must be present

func TestRender_GitClientStatusCommand(t *testing.T) {
	cfg := defaultCfg()
	cfg.AppType = config.AppTypeGitClient

	files, err := scaffold.Render(cfg)
	require.NoError(t, err)

	idx := buildIndex(files)
	assert.Contains(t, idx, "cmd/status.go",
		"git-client template should include cmd/status.go")
}

// ai-assistant should have, the chat command

func TestRender_AIAssistantHasChatCommand(t *testing.T) {
	cfg := defaultCfg()
	cfg.AppType = config.AppTypeAiAssistant

	files, err := scaffold.Render(cfg)
	require.NoError(t, err)

	idx := buildIndex(files)
	assert.Contains(t, idx, "cmd/chat.go",
		"ai-assistant template should include cmd/chat.go")
}

// now - tests for writes, like file's content written on disk or not

func TestWrite_CreateFilesOnDisk(t *testing.T) {
	cfg := defaultCfg()
	dir := t.TempDir()

	files, err := scaffold.Render(cfg)
	require.NoError(t, err)

	err = scaffold.Write(dir, files)
	require.NoError(t, err, "Write should not return an error")

	for _, f := range files {
		dest := filepath.Join(dir, f.RelPath)
		assert.FileExists(t, dest, "expected file %q to exist on disk after write", f.RelPath)
	}

}

// test function for content matches from in-memory  rendered content or not

func TestWrite_ContentMatchesRendered(t *testing.T) {
	cfg := defaultCfg()
	dir := t.TempDir()

	files, err := scaffold.Render(cfg)
	require.NoError(t, err)

	require.NoError(t, scaffold.Write(dir, files))

	for _, f := range files {
		dest := filepath.Join(dir, f.RelPath)
		written, err := os.ReadFile(dest)
		require.NoError(t, err, "should be able to read back %q", dest)
		assert.Equal(t, f.Content, written,
			"disk content of %q should match in-memory render", f.RelPath)
	}
}

// test function for nested directories are being created or not

func TestWrite_CreatesNestedDirectories(t *testing.T) {
	dir := t.TempDir()

	nested := []scaffold.RenderedFile{
		{RelPath: filepath.Join("a", "b", "c", "file.go"), Content: []byte("package c\n")},
		{RelPath: filepath.Join("x", "y", "main.go"), Content: []byte("package main\n")},
	}

	err := scaffold.Write(dir, nested)
	require.NoError(t, err)

	assert.FileExists(t, filepath.Join(dir, "a", "b", "c", "file.go"))
	assert.FileExists(t, filepath.Join(dir, "x", "y", "main.go"))
}

// test function for getting zero error, with no writes at all

func TestWrite_EmptyFileList(t *testing.T) {

	dir := t.TempDir()
	err := scaffold.Write(dir, []scaffold.RenderedFile{})
	require.NoError(t, err, "Write with empty file list should succeed")

}

// test function for correct file permissions are given

func TestWrite_FilePermissions(t *testing.T) {
	dir := t.TempDir()

	files := []scaffold.RenderedFile{
		{RelPath: "hello.go", Content: []byte("package main\n")},
	}

	require.NoError(t, scaffold.Write(dir, files))

	info, err := os.Stat(filepath.Join(dir, "hello.go"))
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0o644), info.Mode().Perm(), "written files should have 0644 permissions")

}

// some last tests - related to the directory present or not (ensureDirectory) and output directory exists or not

//test function to check new directory created everytime

func TestEnsureDir_CreatesDirectory(t *testing.T) {
	base := t.TempDir()
	newDir := filepath.Join(base, "new", "nested", "dir")

	err := scaffold.EnsureDict(newDir)
	require.NoError(t, err)

	info, err := os.Stat(newDir)

	require.NoError(t, err, "directory should exist after EnsureDir")
	require.True(t, info.IsDir(), "%q should be a directory", newDir)

}

// test function to check for already existing directory, should not return error in that case

func TestEnsureDir_IdempotentOnExisting(t *testing.T) {
	existing := t.TempDir()
	err := scaffold.EnsureDict(existing)
	require.NoError(t, err, "EnsureDir on an existing directory should be a no-op")
}

// test function for checking output directory exists or not

func TestOutputDirExists_ReturnsFalseForMissing(t *testing.T) {

	exists, err := scaffold.OutputDirExists("/tmp/this-path-should-never-exist-fastcli-test")
	require.NoError(t, err)

	assert.False(t, exists, "non-existent directory should report as not existing")

}

// test function for any empty directory will not be allowed

func TestOutputDirExists_ReturnsFalseForEmpty(t *testing.T) {
	dir := t.TempDir()

	exists, err := scaffold.OutputDirExists(dir)

	require.NoError(t, err)
	assert.False(t, exists, "empty directory should not count as existsing.")

}

// directory with atleast one entry will be considered as existing

func TestOutputDirExists_ReturnsFalseForNonEmpty(t *testing.T) {
	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "something.txt"), []byte("x"), 0o644))

	exists, err := scaffold.OutputDirExists(dir)

	require.NoError(t, err)

	assert.True(t, exists, "non-empty directory should report as existing")

}

// some test related to IDE

// test function , when no ide is opened by user

func TestOpenIDE_NoneIsNoOp(t *testing.T) {
	opened, err := scaffold.OpenIDE("/some/dir", config.IDENone)
	require.NoError(t, err, "IDENone should never produce an error")
	assert.False(t, opened, "IDENone should report opened=false")
}

// testing function to return unknown return ide error

func TestOpenIDE_ReturnsUnknownIDEError(t *testing.T) {
	_, err := scaffold.OpenIDE("/some/dir", config.IDE("sublimetext"))
	assert.Error(t, err, "unknown IDE constant should produce an error")
}

// test function , if ide binary is missing, and return false instead of error, cause user can open ide manually

func TestOpenIDE_MissingIDEBinaryFallbackMessage(t *testing.T) {
	opened, err := scaffold.OpenIDE(t.TempDir(), config.IDEVscode)

	require.NoError(t, err, "missing ide binary should be non-fatal")

	_ = opened
}

// test function for slice containing, all the ide's installed on user's machine

func TestOpenIDE_AllInstalledIDESlice(t *testing.T) {

	// all ides

	ides := scaffold.DetectIDE()

	assert.NotNil(t, ides, "DetectIDE function will always returns not-nil slice")

	// map contains ide with bool value present or not

	ideMap := map[config.IDE]bool{
		config.IDEVscode: true,
		config.IDECursor: true,
	}

	for _, ide := range ides {
		assert.True(t, ideMap[ide], "detected ide %q is not a known constant", ide)
	}

}
