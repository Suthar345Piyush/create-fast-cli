package scaffold

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/config"
	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/prompt"
)

var Steps = []string{
	"Validating output directory",
	"Rendering templates",
	"Writing project files",
	"Saving preferences",
	"Opening IDE",
}

func Generate(cfg *config.ProjectConfig) error {
	model := prompt.NewProgressModel(Steps)
	p := tea.NewProgram(model)
	errCh := make(chan error, 1)

	go func() {
		errCh <- runSteps(p, cfg)
	}()

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("TUI error: %w", err)
	}

	if err := <-errCh; err != nil {
		return err
	}

	_, err := OpenIDE(cfg.OutputDir, cfg.IDE)
	if err != nil {
		fmt.Println("warning:", err)
	}

	printSuccess(cfg)

	return nil

}

func runSteps(p *tea.Program, cfg *config.ProjectConfig) error {

	exists, err := OutputDirExists(cfg.OutputDir)
	if err != nil {
		p.Send(prompt.StepMsg{Err: err})
		return err
	}

	if exists {
		err = fmt.Errorf("output directory %q already exists and is not empty", cfg.OutputDir)
		p.Send(prompt.StepMsg{Err: err})
		return err
	}
	p.Send(prompt.StepMsg{Label: Steps[0]})

	//rendering

	rendered, err := Render(cfg)
	if err != nil {
		p.Send(prompt.StepMsg{Err: fmt.Errorf("render: %w", err)})
		return err
	}
	p.Send(prompt.StepMsg{Label: Steps[1]})

	//  write

	if err := EnsureDict(cfg.OutputDir); err != nil {
		p.Send(prompt.StepMsg{Err: err})
		return err
	}

	if err := Write(cfg.OutputDir, rendered); err != nil {
		_ = os.RemoveAll(cfg.OutputDir)
		p.Send(prompt.StepMsg{Err: fmt.Errorf("write: %w", err)})
		return err
	}
	p.Send(prompt.StepMsg{Label: Steps[2]})

	//  saving preferences

	_ = config.SavePreferences(cfg)
	p.Send(prompt.StepMsg{Label: Steps[3]})

	// opening selected ide

	// opened, err := OpenIDE(cfg.OutputDir, cfg.IDE)

	if err != nil {
		p.Send(prompt.StepMsg{
			Err: fmt.Errorf("open IDE: %w", err),
		})
		return err
	}

	// if opened {
	// 	p.Send(prompt.StepMsg{Label: Steps[4]})
	// }

	p.Send(prompt.DoneMsg{})

	return nil
}

func printSuccess(cfg *config.ProjectConfig) {
	fmt.Println()
	fmt.Printf("✓ Project %q created at %s\n\n", cfg.ProjectName, cfg.OutputDir)
	fmt.Println()
	fmt.Printf(" 📁 Location:")
	fmt.Printf(" %s\n", cfg.OutputDir)
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println()
	fmt.Printf(" cd  \"%s\"\n", cfg.OutputDir)
	fmt.Println(" go mod tidy")
	fmt.Println(" go run . --help")
	fmt.Println()
}
