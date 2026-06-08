// survey which returns a complete project config

package prompt

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/config"
)

func Run() (*config.ProjectConfig, error) {

	cfg := config.DefaultConfig()
	cfg.IDE = config.PreferredIDE()
	cfg.Framework = config.DefaultFramework()
	cfg.OutputDir = config.DefaultOutputDir()

	customTheme := huh.ThemeFunc(func(isDark bool) *huh.Styles {

		theme := huh.ThemeCharm(isDark)

		theme.Focused.Title = theme.Focused.Title.Foreground(lipgloss.Color("#FF00FF"))

		return theme
	})

	// create-fast-cli banner

	printGoodBanner()

	form := huh.NewForm(
		GroupIdentity(&cfg),
		GroupAppType(&cfg),
		GroupFramework(&cfg),
		GroupFeatures(&cfg),
		GroupOutput(&cfg),
	).WithTheme(customTheme)

	if err := form.Run(); err != nil {
		if err == huh.ErrUserAborted {
			fmt.Println(MutedStyle.Render("\n aborted. No files were written."))
			os.Exit(0)
		}

		return nil, fmt.Errorf("prompt error: %w", err)
	}

	// normalize and deriving the fields

	cfg.ProjectName = strings.TrimSpace(cfg.ProjectName)

	cfg.ModulePath = strings.TrimSpace(cfg.ModulePath)

	cfg.OutputDir = userHomeDirName(strings.TrimSpace(cfg.OutputDir))

	cfg.OutputDir = filepath.Join(cfg.OutputDir, cfg.ProjectName)

	cfg.Language = "go"

	cfg.GoVersion = "1.26"

	// summary function when all write is done

	printSummary(&cfg)

	// final confirmation before writing anything

	var confirmed bool

	confirmForm := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title("Generate Project?").Description("Files will be written to: " + cfg.OutputDir).Value(&confirmed),
		),
	).WithTheme(customTheme)

	if err := confirmForm.Run(); err != nil {
		return nil, err
	}

	if !confirmed {
		fmt.Println(MutedStyle.Render(" cancelled. No files were written."))
	}

	return &cfg, nil

}

// print banner function

func printSummary(cfg *config.ProjectConfig) {

	var sb strings.Builder

	sb.WriteString(SectionStyle.Render(" Review your Choices") + "\n\n")

	row := func(key, val string) {
		sb.WriteString(" ")
		sb.WriteString(SummaryKeyStyle.Render(key))
		sb.WriteString(SummaryValueStyle.Render(val))
		sb.WriteString("\n")
	}

	row("Project name:", cfg.ProjectName)
	row("Module	path:", cfg.ModulePath)
	row("App type:", config.AppTypeLabel(cfg.AppType))
	row("Framework:", config.FrameworkLabel(cfg.Framework))
	row("TUI:", boolLabel(cfg.UseTUI))
	row("Logging:", boolLabel(cfg.UseLogging))
	row("Config:", boolLabel(cfg.UseConfig))
	row("Completions:", boolLabel(cfg.UseCompletions))
	row("Testing:", boolLabel(cfg.UseTesting))
	row("Output Dir:", cfg.OutputDir)
	row("IDE:", config.IDELabel(cfg.IDE))

	fmt.Println(SummaryBoxStyle.Render(sb.String()))

}

func boolLabel(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func userHomeDirName(path string) string {
	if !strings.HasPrefix(path, "-") {
		return path
	}

	home, err := os.UserHomeDir()

	if err != nil {
		return path
	}

	return filepath.Join(home, path[1:])

}
