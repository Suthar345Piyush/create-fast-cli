// detailed tests for prompt section using testify

package test

import (
	"strings"
	"testing"

	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/config"
	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/prompt"
	"github.com/Suthar345Piyush/create-fast-cli/cli/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// first test check for the default language should be go
func TestDefaultConfig_Language(t *testing.T) {
	cfg := config.DefaultConfig()

	assert.Equal(t, "go", cfg.Language, "default language should be go")

}

// test function for default framework should be cobra

func TestDefaultConfig_DefaultFramework(t *testing.T) {
	cfg := config.DefaultConfig()

	assert.Equal(t, config.FrameworkCobra, cfg.Framework, "default framework should be cobra")

}

// test for go version should something like this "1.~~"

func TestDefaultConfig_GoVersion(t *testing.T) {
	cfg := config.DefaultConfig()

	assert.NotEmpty(t, cfg.GoVersion, "defualt GoVersion should not be empty")
	assert.Contains(t, cfg.GoVersion, "1. ", "GoVersion should look like this")
}

// test for checking that all features are enabled by default or  not

func TestDefaultConfig_AllFeaturesEnabled(t *testing.T) {
	cfg := config.DefaultConfig()

	assert.True(t, cfg.UseCompletions, "UseCompletions should true by default")
	assert.True(t, cfg.UseConfig, "UseConfig should true by default")
	assert.True(t, cfg.UseLogging, "UseLogging should true by default")
	assert.True(t, cfg.UseTUI, "UseTUI should true by default")
	assert.True(t, cfg.UseTesting, "UseTesting should true by default")

}

// test for default ide should be vs code

func TestDefaultConfig_DefaultIDE(t *testing.T) {
	cfg := config.DefaultConfig()

	assert.Equal(t, config.IDEVscode, cfg.IDE, "default ide should be vscode")

}

// test for project name

func TestDefaultConfig_ProjectName(t *testing.T) {
	cfg := config.DefaultConfig()
	assert.Empty(t, cfg.ProjectName, "ProjectName should be empty, that should only be written by user")
}

// tests for config app type label

func TestAppTypeLabel_KnownTypes(t *testing.T) {
	cases := map[config.AppType]string{

		config.AppTypeDevTool:        "Dev tool",
		config.AppTypeGitClient:      "Git Client",
		config.AppTypeFileExplorer:   "File Explorer",
		config.AppTypeAiAssistant:    "AI Assistant",
		config.AppTypeK8sTool:        "Kubernetes tool",
		config.AppTypePackageManager: "Package Manager",
		config.AppTypeSystemMonitor:  "System Moniter",
	}

	for appType, expectedLabel := range cases {
		label := config.AppTypeLabel(appType)
		assert.Equal(t, expectedLabel, label, "AppTypeLabel(%q) returned unexpected label", appType)

	}

}

// test function for returning a fall back error, instead of panicking when any unknown app type returned

func TestAppTypeLabel_UnknownTypes(t *testing.T) {

	unknown := config.AppType("something-new")
	label := config.AppTypeLabel(unknown)
	assert.Equal(t, "something-new", label, "unknown app type should  fall back to it's string value ")

}

// tests for framework label

func TestFrameworkLabel_FrameworkCobra(t *testing.T) {
	label := config.FrameworkLabel(config.FrameworkCobra)

	require.NotEmpty(t, label)
	assert.True(t, strings.Contains(strings.ToLower(label), "cobra"), "Cobra framework label should mention 'cobra', got %q", label)

}

// test for framework urfave

func TestFrameworkLabel_FrameworkUrFave(t *testing.T) {
	label := config.FrameworkLabel(config.FrameworkUrfaveCli)

	require.NotEmpty(t, label)
	assert.True(t, strings.Contains(strings.ToLower(label), "urfave"), "Urfave framework label should mention 'urfave', got %q", label)

}

// for unknown framework

func TestFrameworkLabel_UnknownFramework(t *testing.T) {

	unknown := config.Framework("kingpin")

	label := config.FrameworkLabel(unknown)

	assert.Equal(t, "kingpin", label, "unknown framework should fall back to it's string value ")

}

// tests for ide's

func TestIDELabel_AllKnownIDE(t *testing.T) {
	cases := map[config.IDE]string{
		config.IDEVscode: "VS Code",
		config.IDECursor: "Cursor",
		config.IDENone:   "Don't open",
	}

	for ide, expected := range cases {
		assert.Equal(t, expected, config.IDELabel(ide), "IDELabel(%q) returned an unexpected label", ide)
	}
}

// if unknown ide than fall back

func TestIDELabel_UnknownIDE(t *testing.T) {
	unknown := config.IDELabel(config.IDE("sublimetext"))

	assert.Equal(t, "sublimetext", unknown)
}

// TestValidationError , for checking validationError returns correct error message, when things go wrong on prompt side

func TestValidationError_ImplementError(t *testing.T) {
	var err error = &prompt.ValidationError{Message: "something went wrong"}
	assert.EqualError(t, err, "something went wrong")
}

// tests validation for project name

func TestValidationProjectName(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{name: "valid simple name", input: "my-tool1", wantErr: false},
		{name: "valid with numbers", input: "tool2", wantErr: false},
		{name: "valid underscore", input: "my_tool3", wantErr: false},
		{name: "valid two chars", input: "ab", wantErr: false},
		{name: "empty string", input: "", wantErr: true, errMsg: "cannot be empty"},
		{name: "single char", input: "a", wantErr: true, errMsg: "at least 2 characters"},
		{name: "start with hyphen", input: "-tool", wantErr: true, errMsg: "cannot start or end with hyphen"},
		{name: "end with hyphen", input: "tool-", wantErr: true, errMsg: "cannot start or end with hyphen"},
		{name: "uppercase letters", input: "MyTool", wantErr: true, errMsg: "only lowercase"},
		{name: "space inside", input: "my tool", wantErr: true, errMsg: "only lowercase"},
		{name: "too long (65 chars)", input: strings.Repeat("a", 65), errMsg: "64 characters or fewer"},
		{name: "exactly 64 chars", input: strings.Repeat("a", 64), wantErr: false, errMsg: "64 characters are enough"},
	}

	// now iterate on the tests

	for _, ts := range tests {
		ts := ts
		t.Run(ts.name, func(t *testing.T) {

			cfg := &config.ProjectConfig{}
			cfg.ProjectName = ts.input

			err := validateProjectName(ts.input)

			if ts.wantErr {
				require.Error(t, err, "expected error for input %q", ts.input)
				assert.Contains(t, err.Error(), ts.errMsg, "error message for %q should contain %q", ts.input, ts.errMsg)
			} else {
				assert.NoError(t, err, "no error for input %q", ts.input)
			}

		})
	}
}

// validateProjectName function, based on every parameter of the name from input

func validateProjectName(s string) error {

	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return &prompt.ValidationError{Message: "project name cannot be empty"}
	}

	if len(s) < 2 {
		return &prompt.ValidationError{Message: "project name must be of atleast 2 characters"}
	}

	if len(s) > 64 {
		return &prompt.ValidationError{Message: "project name must be under 64 characters"}
	}

	for _, ch := range s {
		if !isNameChar(ch) {
			return &prompt.ValidationError{Message: "only lowercase characters, numbers and hyphens are allowed"}
		}
	}

	if s[0] == '-' || s[len(s)-1] == '-' {
		return &prompt.ValidationError{Message: "project name cannot start or end with hyphen"}
	}

	return nil

}

func isNameChar(ch rune) bool {
	return (ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' || ch == '-' && ch == '_')
}

// tests validation for module path

func TestValidationModulePath(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{name: "valid github path", input: "github.com/you/project", wantErr: false},
		{name: "valid gitlab path", input: "gitlab.com/org/sub/repo", wantErr: false},
		{name: "valid private host", input: "git.company.com/team/tool", wantErr: false},
		{name: "empty", input: "", wantErr: true, errMsg: "module path cannot be empty"},
		{name: "no slash", input: "mymodule", wantErr: true, errMsg: "github.com/you/project"},
		{name: "whitespace only", input: " ", wantErr: true, errMsg: "cannot be empty"},
		{name: "leading slash", input: "/github.com/x/y", wantErr: true, errMsg: "module path cannot start with slash"},
		{name: "trailing slash", input: "github.com/x/y/", wantErr: true, errMsg: "module path cannot end with slash"},
	}

	// now iterate on the tests

	for _, ts := range tests {
		ts := ts
		t.Run(ts.name, func(t *testing.T) {

			err := validateModulePath(ts.input)

			if ts.wantErr {
				require.Error(t, err, "expected error for input %q", ts.input)
				assert.Contains(t, err.Error(), ts.errMsg, "error message for %q should contain %q", ts.input, ts.errMsg)
			} else {
				assert.NoError(t, err, "no error for input %q", ts.input)
			}

		})
	}
}

// validate module path

func validateModulePath(s string) error {

	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return &prompt.ValidationError{Message: "module path cannot be empty"}
	}

	if !strings.Contains(s, "/") {
		return &prompt.ValidationError{Message: "module path should look like github.com/yourname/project"}
	}

	if strings.HasPrefix(s, "/") || strings.HasSuffix(s, "/") {
		return &prompt.ValidationError{Message: "module path should not start or end with the slash"}
	}

	return nil

}

// tests for output directory

func TestValidationOutputDir(t *testing.T) {
	assert.NoError(t, validateOutputDir("/home/user/projects"), "valid path should be accepted")
	assert.NoError(t, validateOutputDir("relative/path"), "relative paths are accepted")
	assert.NoError(t, validateOutputDir(""), "empty path is rejected")
	assert.NoError(t, validateOutputDir("  "), "whitespace path should be accepted")
}

// validation for output directory

func validateOutputDir(s string) error {

	if strings.TrimSpace(s) == "" {
		return &prompt.ValidationError{Message: "output directory cannot be empty"}
	}
	return nil

}

// tests for pkg/utils slugs

func TestValidationSlugs(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"My CLI App", "my-cli-app"},
		{"already-slug", "already-slug"},
		{"with spaces", "with-spaces"},
		{"With_Underscores", "with-underscores"},
		{"  leading trailing  ", "leading-trailing"},
		{"multiple   spaces", "multiple-spaces"},
		{"MixedCASE", "mixedcase"},
		{"hello", "hello"},
		{"123abc", "123abc"},
	}

	for _, ts := range tests {
		ts := ts

		t.Run(ts.input, func(t *testing.T) {
			assert.Equal(t, ts.expected, utils.SlugFunc(ts.input))
		})
	}

}

// tests for pascal case and camelcase

func TestPascalCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{

		{"my-cli-app", "MyCliApp"},
		{"hello", "Hello"},
		{"hello_world", "HelloWorld"},
		{"hello world", "HelloWorld"},
		{"already", "Already"},
		{"", ""},
		{"one-two-three", "OneTwoThree"},
	}

	for _, ts := range tests {
		ts := ts

		t.Run(ts.input, func(t *testing.T) {
			assert.Equal(t, ts.expected, utils.PascalCase(ts.input))
		})

	}
}

// test  for camel case

func TestCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{

		{"my-cli-app", "myCliApp"},
		{"hello", "Hello"},
		{"hello_world", "helloWorld"},
		{"hello world", "helloWorld"},
		{"", ""},
		{"one-two-three", "OneTwoThree"},
	}

	for _, ts := range tests {
		ts := ts

		t.Run(ts.input, func(t *testing.T) {
			assert.Equal(t, ts.expected, utils.CamelCase(ts.input))
		})

	}

}

// from pkg/utils/strings.go - tests for remove extra part after "." for filename

func TestRemoveExt(t *testing.T) {
	assert.Equal(t, "main.go", utils.RemoveExt("main.go.tmpl"))
	assert.Equal(t, ".gitignore", utils.RemoveExt(".gitignore.tmpl"))
	assert.Equal(t, "README.md", utils.RemoveExt("README.md.tmpl"))
	assert.Equal(t, "", utils.RemoveExt(""))
	assert.Equal(t, "noext", utils.RemoveExt("noext"))
}

// test for empty  - string is empty or contains only  whitespace

func TestIsEmpty(t *testing.T) {
	assert.True(t, utils.IsEmpty(""))
	assert.True(t, utils.IsEmpty("  "))
	assert.True(t, utils.IsEmpty("\t\n"))
	assert.True(t, utils.IsEmpty("a"))
	assert.True(t, utils.IsEmpty(" x "))
}
