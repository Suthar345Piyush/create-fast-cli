package prompt

import (
	"strings"

	"charm.land/huh/v2"
	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/config"
)

// ValidationError is exported so test packages can reference it.
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string { return e.Message }

func validateProjectName(s string) error {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return &ValidationError{"project name cannot be empty"}
	}
	if len(s) < 2 {
		return &ValidationError{"project name must be at least 2 characters"}
	}
	if len(s) > 64 {
		return &ValidationError{"project name must be 64 characters or fewer"}
	}
	for _, c := range s {
		if !isNameChar(c) {
			return &ValidationError{"only lowercase letters, numbers, and hyphens are allowed"}
		}
	}
	if s[0] == '-' || s[len(s)-1] == '-' {
		return &ValidationError{"project name cannot start or end with a hyphen"}
	}
	return nil
}

func isNameChar(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-' || c == '_'
}

func validateModulePath(s string) error {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return &ValidationError{"module path cannot be empty"}
	}
	if !strings.Contains(s, "/") {
		return &ValidationError{"module path should look like github.com/you/project"}
	}
	if strings.HasPrefix(s, "/") || strings.HasSuffix(s, "/") {
		return &ValidationError{"module path should not start or end with /"}
	}
	return nil
}

func validateOutputDir(s string) error {
	if strings.TrimSpace(s) == "" {
		return &ValidationError{"output directory cannot be empty"}
	}
	return nil
}

func GroupIdentity(cfg *config.ProjectConfig) *huh.Group {
	return huh.NewGroup(
		huh.NewInput().
			Title("Project name").
			Description("Lowercase letters, numbers, hyphens only (e.g. my-dev-tool)").
			Placeholder("my-cli-app").
			Value(&cfg.ProjectName).
			Validate(validateProjectName),
		huh.NewInput().
			Title("Go module path").
			Description("Used in go.mod (e.g. github.com/you/my-cli-app)").
			Placeholder("github.com/you/my-cli-app").
			Value(&cfg.ModulePath).
			Validate(validateModulePath),
	)
}

func GroupAppType(cfg *config.ProjectConfig) *huh.Group {
	return huh.NewGroup(
		huh.NewSelect[config.AppType]().
			Title("What are you building?").
			Options(
				huh.NewOption("Dev tool         — build system, linter, code gen", config.AppTypeDevTool),
				huh.NewOption("Git client        — custom git workflows & hooks", config.AppTypeGitClient),
				huh.NewOption("File explorer    — navigate & manipulate the filesystem", config.AppTypeFileExplorer),
				huh.NewOption("Kubernetes tool  — kubectl plugin or cluster manager", config.AppTypeK8sTool),
				huh.NewOption("AI assistant     — LLM-powered terminal assistant", config.AppTypeAiAssistant),
				huh.NewOption("System monitor   — CPU, memory, disk, process viewer", config.AppTypeSystemMonitor),
				huh.NewOption("Package manager  — install / update / remove packages", config.AppTypePackageManager),
			).
			Value(&cfg.AppType),
	)
}

func GroupFramework(cfg *config.ProjectConfig) *huh.Group {
	return huh.NewGroup(
		huh.NewSelect[config.Framework]().
			Title("CLI framework").
			Options(
				huh.NewOption("Cobra     (recommended — most popular Go CLI framework)", config.FrameworkCobra),
				huh.NewOption("urfave/cli (lightweight alternative)", config.FrameworkUrfaveCli),
			).
			Value(&cfg.Framework),
	)
}

func GroupFeatures(cfg *config.ProjectConfig) *huh.Group {
	return huh.NewGroup(
		huh.NewConfirm().Title("Include TUI? (Bubbletea + Lipgloss)").Value(&cfg.UseTUI),
		huh.NewConfirm().Title("Include structured logging? (Uber Zap)").Value(&cfg.UseLogging),
		huh.NewConfirm().Title("Include config file support? (Viper)").Value(&cfg.UseConfig),
		huh.NewConfirm().Title("Include shell completions?").Value(&cfg.UseCompletions),
		huh.NewConfirm().Title("Include testing setup? (Testify)").Value(&cfg.UseTesting),
	)
}

func GroupOutput(cfg *config.ProjectConfig) *huh.Group {
	return huh.NewGroup(
		huh.NewInput().
			Title("Output directory").
			Placeholder("/home/you/projects").
			Value(&cfg.OutputDir).
			Validate(validateOutputDir),
		huh.NewSelect[config.IDE]().
			Title("Open project in which IDE after generation?").
			Options(
				huh.NewOption("VS Code", config.IDEVscode),
				huh.NewOption("Cursor", config.IDECursor),
				huh.NewOption("Don't open — I'll do it myself", config.IDENone),
			).
			Value(&cfg.IDE),
	)
}
