package prompt

import (
	"strings"

	"charm.land/huh/v2"
	"github.com/suthar345piyush/internal/config"
)

// build app type options for asking what user want to build, they select from this options

func buildAppTypeOptions() []huh.Option[config.AppType] {

	return []huh.Option[config.AppType]{

		huh.NewOption("Dev tool - build system, linter, code gen", config.AppTypeDevTool),

		huh.NewOption("Git client - custom git workflows & hooks", config.AppTypeGitClient),

		huh.NewOption("File explorer - navigate & manipulate the file system", config.AppTypeFileExplorer),

		huh.NewOption("Kubernetes tool -  kubectl plugin or cluster management tool", config.AppTypeK8sTool),

		huh.NewOption("AI assistant - LLM-powered terminal assistant", config.AppTypeAiAssistant),

		huh.NewOption("System moniter - CPU, memory, Disk, storage", config.AppTypeSystemMoniter),

		huh.NewOption("Package Manager - install / update / remove package", config.AppTypePackageManager),
	}

}

// build framework options - choose between two frameworks

func buildFrameworkOptions() []huh.Option[config.Framework] {
	return []huh.Option[config.Framework]{
		huh.NewOption("Cobra - (most popular Go CLI framework)", config.FrameworkCobra),

		huh.NewOption("urfave/cli - good cobra alternative CLI framework", config.FrameworkUrfaveCli),
	}
}

//build IDE options - choose between ide options

func buildIDEOptions() []huh.Option[config.IDE] {
	return []huh.Option[config.IDE]{
		huh.NewOption("VS Code", config.IDEVscode),
		huh.NewOption("Cursor", config.IDECursor),
		huh.NewOption("Don't open I'll do it myself", config.IDENone),
	}
}

// some validations like - project name validation

func projectNameValidation(s string) error {

	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return &ValidationError{"project name cannot be empty"}
	}

	if len(s) < 2 {
		return &ValidationError{"project name must be atleast of 2 characters"}
	}

	if len(s) > 64 {
		return &ValidationError{"project name must be 64 characters or fewer"}
	}

	for _, c := range s {
		if !isNameChar(c) {
			return &ValidationError{"only lowercase letters, numbers and hyphens are allowed"}
		}
	}

	if s[0] == '-' || s[len(s)-1] == '-' {

		return &ValidationError{"project name cannot start or end with hyphen"}

	}

	return nil

}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func isNameChar(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-' || c == '_'
}

// function for validating the module path

func validateModulePath(s string) error {

	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return &ValidationError{"module path name cannot be empty"}
	}

	if !strings.Contains(s, "/") {
		return &ValidationError{"module path should look like github.com/you/project"}
	}

	if strings.HasPrefix(s, "/") || strings.HasSuffix(s, "/") {
		return &ValidationError{"module path should not start or end with /	"}
	}

	return nil

}

//function to validate the output directory

func validateOutputDir(s string) error {
	if strings.TrimSpace(s) == "" {
		return &ValidationError{"output directory cannot be empty"}
	}
	return nil
}

// huh form in the form of *huh.Group

// basically asking user to select or to fill the form

func GroupIdentity(cfg *config.ProjectConfig) *huh.Group {

	return huh.NewGroup(

		huh.NewInput().Title("Project name").Description("Lowercase letters, numbers, hyphens only").Placeholder("my-cli-app").Value(&cfg.ProjectName).Validate(projectNameValidation),

		huh.NewInput().Title("Go module path").Description("Used in go.mod").Placeholder("github.com/you/my-cli-app").Value(&cfg.ModulePath).Validate(validateModulePath),
	)

}
